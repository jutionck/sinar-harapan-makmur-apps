package main

import "time"

type EntryLog struct {
	StartTime    time.Time
	EndTime      time.Duration
	StatusCode   int
	ClientAIP    string
	Method       string
	RelativePath string
	UserAgent    string
}
