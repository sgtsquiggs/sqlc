package compiler

import (
	"github.com/sgtsquiggs/sqlc/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
