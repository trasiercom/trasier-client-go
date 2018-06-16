package api

type ContentType string

const (
	TEXT      = ContentType("TEXT")
	XML       = ContentType("XML")
	JSON      = ContentType("JSON")
	SQL       = ContentType("SQL")
	BINARY    = ContentType("OCTET-STREAM")
	ENCRYPTED = ContentType("PGP-ENCRYPTED")
)
