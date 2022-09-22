package models

import "time"

type Subscription struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	CreatedDate     time.Time `json:"created_date"`
	Email_Templates int64     `json:"email_template"`
	Capture_Pages   int64     `json:"email_pagess"`
}
