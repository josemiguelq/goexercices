package main

const (
	URL = "url"
)

type Input struct {
	Message string `json:"message"`
	Secret  string `json:"secret"`
}

type Secret struct {
	Hash    string `json:"hashe" `
	Message string `json:"message"`
	Secret  string `json:"secret"`
}
