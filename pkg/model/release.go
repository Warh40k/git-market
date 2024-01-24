package model

import "time"

type Release struct {
	App      *Gitrepo
	Version  string
	Date     time.Time
	Comments string
}
