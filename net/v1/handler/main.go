package handler

import (
	"go.uber.org/dig"

	"github.com/alissonperez/api-foo/infra/plog"
	"github.com/alissonperez/api-foo/service/client"
	"github.com/alissonperez/api-foo/service/url"
)

type HandlerDependencies struct {
	dig.In

	ClientService      client.ClientService
	UrlService         url.UrlService
	Logger             plog.Log
}
