package models

import "time"

type Notification struct {
	Data   interface{} `json:"data"`
	Header interface{} `json:"header"`
	Method string      `json:"method"`
	Time   time.Time   `json:"time"`
}
