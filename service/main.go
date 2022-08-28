package service

import (
	"github.com/alissonperez/api-foo/service/client"
	"github.com/alissonperez/api-foo/service/url"
	"go.uber.org/dig"
)

func Provide(container *dig.Container) {
	container.Provide(client.NewService)
	container.Provide(url.NewService)
}
