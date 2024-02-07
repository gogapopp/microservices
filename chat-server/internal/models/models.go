package models

import "time"

type Message struct {
	From      string
	Text      string
	Timestamp time.Time
}
