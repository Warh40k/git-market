package model

import "time"

type File struct {
	Title        string
	Size         int
	DateUpload   *time.Time
	DownloadLink string
}
