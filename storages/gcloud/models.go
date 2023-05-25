package gcloud

type FileType string

const (
	Dir  FileType = "dir"
	File FileType = "file"
)

type Item struct {
	Path string   `json:"path"`
	Type FileType `json:"type"`
}
