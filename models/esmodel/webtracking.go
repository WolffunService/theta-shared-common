package esmodel

import (
	"time"
)

type WebTrackingMapping struct {
	UserId    string    `json:"user_id"`
	Domain    string    `json:"domain"`
	Path      string    `json:"path"`
	UserAgent string    `json:"user_agent"`
	Ip        string    `json:"ip"`
	Timestamp time.Time `json:"@timestamp"`
}

func (WebTrackingMapping) Index() string {
	return "web-tracking"
}
