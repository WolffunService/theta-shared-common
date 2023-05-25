package esmodel

import "time"

type UniversalUserStatMapping struct {
	User      interface{} `json:"user"`
	StatName  string      `json:"stat_name"`
	StatValue float64     `json:"stat_value"`
	Timestamp time.Time   `json:"@timestamp"`
}
