package service

type GitRepo interface {
}

type Release interface {
}

type File interface {
}

type Service struct {
	GitRepo
	Release
	File
}

func NewService() *Service {
	return &Service{}
}
