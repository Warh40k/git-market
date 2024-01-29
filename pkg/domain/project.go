package domain

type Project struct {
	Name        string `json:"name"`
	Owner       Owner  `json:"owner"`
	Description string `json:"description"`
}
