package ioc

import (
	"github.com/gorilla/mux"
	"github.com/alissonperez/api-foo/auth"
	"github.com/alissonperez/api-foo/config"
	"github.com/alissonperez/api-foo/infra/plog"
	"github.com/alissonperez/api-foo/infra/teardown"
	"github.com/alissonperez/api-foo/repository"
	"github.com/alissonperez/api-foo/service"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	c := dig.New()

	c.Provide(mux.NewRouter)

	teardown.Provide(c)
	config.Provide(c)
	auth.Provide(c)
	repository.Provide(c)
	service.Provide(c)
	plog.Provide(c)

	return c
}
