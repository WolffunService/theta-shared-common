package lbrewardapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// ================================ util
var deadline = 10 * time.Second

func convertRawQuery(model interface{}) string {
	var inInterface map[string]interface{}
	q := url.Values{}
	inrec, _ := json.Marshal(model)
	json.Unmarshal(inrec, &inInterface)

	// iterate through inrecs
	for field, val := range inInterface {
		value := fmt.Sprint(val)
		if value != "" {
			q.Add(field, value)
		}
	}
	return encode(q)
}

func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	//sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

func newGetRequestParams(address string, method string, data interface{}) ([]byte, int, error) {

	base, err := url.Parse(address)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	base.Path = path.Join(base.Path, method)
	base.RawQuery = convertRawQuery(data)
	client := http.Client{
		Timeout: deadline,
	}
	res, errGet := client.Get(base.String())
	if errGet != nil {
		return nil, http.StatusBadRequest, errGet
	}

	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, http.StatusBadRequest, errRead
	}
	return body, http.StatusOK, nil
}

// ================================== mresty
func gETDynamic(absolute string, queryString string, pathParams map[string]string, token string) (*resty.Response, error) {
	client := resty.New()

	return client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(token).
		SetQueryString(queryString).
		SetPathParams(pathParams).
		Get(absolute)
}
