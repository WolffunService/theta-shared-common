package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func ConvertRawQuery(model interface{}) string {
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
