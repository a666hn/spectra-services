package connection

import (
	// Postgres dialect
	_ "github.com/lib/pq"
)

type Connection interface{}

type connection struct{}
