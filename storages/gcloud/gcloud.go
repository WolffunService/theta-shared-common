package gcloud

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/appengine/file"
)

var (
	ErrObjectAlreadyExist = errors.New("the object is already exist")
	ErrStorageClientIsNil = errors.New("storageClient is nil")
)

type ClientUploader struct {
	Client            *storage.Client
	Bucket            *storage.BucketHandle
	CurrentBucketName string
	UploadPath        string
}

var storageClient *storage.Client

func InitGCloudBucket(opts ...option.ClientOption) error {
	client, err := storage.NewClient(context.Background(), opts...)
	if err != nil {
		return fmt.Errorf("failed to create storage client: %v", err)
	}
	storageClient = client
	return nil
}

func NewClientUploader(bucketName string, uploadPath string) (*ClientUploader, error) {
	if storageClient == nil {
		return nil, ErrStorageClientIsNil
	}
	if bucketName == "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if r, e := file.DefaultBucketName(ctx); e != nil {
			return nil, e
		} else {
			bucketName = r
		}
	}
	uploader := &ClientUploader{
		Client:            storageClient,
		Bucket:            storageClient.Bucket(bucketName),
		CurrentBucketName: bucketName,
		UploadPath:        uploadPath,
	}
	return uploader, nil
}

// CreateObject create an object
func (c *ClientUploader) CreateObject(ctx context.Context, name string) (*storage.ObjectAttrs, error) {
	obj := c.UploadPath + "/" + name
	if c.IsObjectExist(obj) {
		return nil, ErrObjectAlreadyExist
	}
	// Upload an object with storage.Writer.
	wc := c.Bucket.Object(obj).NewWriter(ctx)
	wc.ContentType = "text/plain"
	if _, err := wc.Write([]byte("# Create Object\n")); err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}

	return wc.Attrs(), nil
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(ctx context.Context, file multipart.File, filename string) (*storage.ObjectAttrs, error) {
	// Upload an object with storage.Writer.
	obj := c.UploadPath + "/" + filename
	if c.IsObjectExist(obj) {
		return nil, ErrObjectAlreadyExist
	}
	wc := c.Bucket.Object(obj).NewWriter(ctx)
	wc.Metadata = map[string]string{
		"x-google-meta-bucket":    c.CurrentBucketName,
		"x-google-meta-name":      filename,
		"x-google-meta-timestamp": time.Now().String(),
	}
	if _, err := io.Copy(wc, file); err != nil {
		return nil, fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %v", err)
	}

	return wc.Attrs(), nil
}

// UploadFileWithContentType uploads an object
func (c *ClientUploader) UploadFileWithContentType(ctx context.Context, file io.Reader, filename string, contentType string) (*storage.ObjectAttrs, error) {
	// Upload an object with storage.Writer.
	obj := fmt.Sprintf("%s/%s", c.UploadPath, filename)
	if c.IsObjectExist(obj) {
		return nil, ErrObjectAlreadyExist
	}

	wc := c.Bucket.Object(obj).NewWriter(ctx)
	wc.ContentType = contentType
	wc.Metadata = map[string]string{
		"x-google-meta-bucket":    c.CurrentBucketName,
		"x-google-meta-name":      filename,
		"x-google-meta-timestamp": time.Now().String(),
	}

	if _, err := io.Copy(wc, file); err != nil {
		return nil, fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %v", err)
	}

	return wc.Attrs(), nil
}

// List lists objects within specified bucket.
func (c *ClientUploader) List(prefix string) ([]*storage.ObjectAttrs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	it := c.Bucket.Objects(ctx, &storage.Query{
		Prefix: prefix,
	})
	var result []*storage.ObjectAttrs

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %v", c.CurrentBucketName, err)
		}
		result = append(result, attrs)
	}
	return result, nil
}

// DeleteFile removes specified object.
func (c *ClientUploader) DeleteFile(ctx context.Context, object string) (*storage.ObjectAttrs, error) {
	o := c.Bucket.Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("object.Attrs: %v", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return nil, fmt.Errorf("Object(%q).Delete: %v", object, err)
	}
	return attrs, nil
}

// ReadFile reads the named file in Google Cloud Storage.
func (c *ClientUploader) ReadFile(ctx context.Context, fileName string) ([]byte, string, error) {
	rc, err := c.Bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("readFile: unable to open file from bucket %q, file %q: %v", c.CurrentBucketName, fileName, err)
	}
	defer rc.Close()

	slurp, err := io.ReadAll(rc)
	if err != nil {

		return nil, "", fmt.Errorf("readFile: unable to read data from bucket %q, file %q: %v", c.CurrentBucketName, fileName, err)
	}
	return slurp, rc.Attrs.ContentType, nil
}

// Attrs read file metadata in Google Cloud Storage.
func (c *ClientUploader) Attrs(ctx context.Context, fileName string) (*storage.ObjectAttrs, error) {
	obj := c.Bucket.Object(fileName)
	return obj.Attrs(ctx)
}

// IsObjectExist check the named object in Google Cloud Storage.
func (c *ClientUploader) IsObjectExist(object string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rc, err := c.Bucket.Object(object).NewReader(ctx)
	if err != nil {
		return false
	}
	defer rc.Close()
	return true
}

// DeleteFiles This function can delete multiple files on a bucket based on their path and object type.
func (c *ClientUploader) DeleteFiles(ctx context.Context, listFile []Item) ([]*storage.ObjectAttrs, error) {
	var result []*storage.ObjectAttrs
	deleteObj := func(name string) (*storage.ObjectAttrs, error) {
		fi, errDelete := c.DeleteFile(ctx, name)
		if errDelete != nil {
			return nil, errDelete
		}
		return fi, errDelete
	}

	for _, f := range listFile {
		if f.Type == File {
			if fi, err := deleteObj(f.Path); err != nil {
				return result, err
			} else {
				result = append(result, fi)
			}
		}
		if f.Type == Dir {
			attrs, err := c.List(f.Path + "/")
			if err != nil {
				return nil, err
			}
			for _, v := range attrs {
				fi, err := deleteObj(v.Name)
				if err != nil {
					return nil, err
				} else {
					result = append(result, fi)
				}
			}
		}
	}

	return result, nil
}

func (c *ClientUploader) CopyObject(ctx context.Context, srcBucket, srcObject, dstBucket, dstObject string) (*storage.ObjectAttrs, error) {
	src := c.Bucket.Object(fmt.Sprintf("%s/%s", srcBucket, srcObject))
	dst := c.Bucket.Object(fmt.Sprintf("%s/%s", dstBucket, dstObject))

	//dst = dst.If(storage.Conditions{DoesNotExist: true})

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return nil, fmt.Errorf("Object(%q).CopierFrom(%q).Run: %v", srcObject, srcObject, err)
	}

	return dst.Attrs(ctx)
}
