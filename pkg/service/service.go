package service

import "github.com/Warh40k/gitmarket/pkg/domain"

type Project interface {
	SearchRepos(searchString string) ([]domain.Project, error)
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
