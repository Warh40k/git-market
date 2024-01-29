package service

import (
	"github.com/Warh40k/gitmarket/pkg/domain"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"os"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

type SearchResponse struct {
	Items []domain.Project `json:"items"`
}

func (s *ProjectService) SearchRepos(searchString string) ([]domain.Project, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetResult(&SearchResponse{}).
		SetHeaders(map[string]string{
			"Authorization": "Bearer " + os.Getenv("GH_TOKEN"),
			"Accept":        "application/vnd.github+json",
		}).
		SetQueryParam("q", searchString).
		Get(viper.GetString("rest.github.url") + viper.GetString("rest.github.repos.search"))
	if err != nil {
		return nil, err
	}

	data := resp.Result().(*SearchResponse)
	if err != nil {
		return nil, err
	}

	return data.Items, nil
}
