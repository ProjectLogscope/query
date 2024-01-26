package main

import "github.com/hardeepnarang10/query/pkg/authorization"

type authorizationConfig struct {
	authorizationAccessDefault []string
	authorizationAccessLimited []string
}

func initAuthorization(ac authorizationConfig) authorization.Authorization {
	return authorization.New(
		ac.authorizationAccessDefault,
		ac.authorizationAccessLimited,
	)
}
