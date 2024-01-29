package domain

import "time"

type Release struct {
	App      *Project
	Version  string
	Date     time.Time
	Comments string
}
