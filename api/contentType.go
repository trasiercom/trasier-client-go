package api

type ContentType string

const (
	TEXT      = ContentType("text/plain")
	XML       = ContentType("application/xml")
	JSON      = ContentType("application/json")
	SQL       = ContentType("application/sql")
	BINARY    = ContentType("application/octet-stream")
	ENCRYPTED = ContentType("application/pgp-encrypted")
)
