package repository

import (
	resources "github.com/alissonperez/api-foo/resources"
)

type ClientRepository interface {
	GetById(id int) (resources.Client, bool)
}
