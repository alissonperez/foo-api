package client

import (
	"fmt"
	"github.com/alissonperez/api-foo/auth"
	"github.com/alissonperez/api-foo/repository"
	"github.com/alissonperez/api-foo/resources"
	"net/http"
)

type ClientService struct {
	repo repository.ClientRepository
	auth auth.Auth
}

func (s ClientService) GetClientFromRequest(r *http.Request) (resources.Client, error) {
	authData, err := s.auth.FromRequest(r)
	if err != nil {
		return resources.Client{}, err
	}

	client, ok := s.repo.GetById(authData.ClientId)
	if !ok {
		return resources.Client{}, fmt.Errorf("client not found")
	}

	return client, nil
}

func NewService(repo repository.ClientRepository, authObj auth.Auth) ClientService {
	return ClientService{repo: repo, auth: authObj}
}
