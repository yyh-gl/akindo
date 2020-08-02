package main

type AkindoClient struct{}

func NewAkindoClient() *AkindoClient {
	return &AkindoClient{}
}

func (ac AkindoClient) GetAccount() (string, error) {
	return "", nil
}
