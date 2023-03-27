package query

import _ "embed"

var (
	//go:embed user.insert_user.sql
	InsertUserRegistration string
)
