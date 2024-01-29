package rest

import "github.com/go-resty/resty/v2"

type RestClient struct {
	token  string
	client *resty.Client
}
