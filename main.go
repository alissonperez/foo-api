package main

import (
	"github.com/alissonperez/api-foo/infra/ioc"
	"github.com/alissonperez/api-foo/net"
)

func main() {
	net.SetupServer(ioc.CreateContainer())
}
