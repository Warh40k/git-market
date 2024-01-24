package model

import "time"

type Release struct {
	App      *Project
	Version  string
	Date     time.Time
	Comments string
}
