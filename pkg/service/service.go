package service

import "github.com/Warh40k/gitmarket/pkg/model"

type Project interface {
	SearchRepos(searchString string) ([]model.Project, error)
}

type Release interface {
}

type File interface {
}

type Service struct {
	Project
	Release
	File
}

func NewService() *Service {
	return &Service{
		Project: NewProjectService(),
	}
}
