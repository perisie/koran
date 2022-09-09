//go:build wireinject
// +build wireinject

// https://github.com/golang/vscode-go/issues/2075
package main

import (
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/services"
	"github.com/google/wire"
)

func wireUserDaoImpl() (*daos.UserDaoImpl, error) {
	wire.Build(daos.NewUserDaoImpl, NewDb)
	return nil, nil
}

func wireGoogleAuthServiceImpl() (*services.GoogleAuthServiceImpl, error) {
	wire.Build(
		services.NewGoogleAuthServiceImpl,
		wire.Bind(new(daos.UserDao), new(*daos.UserDaoImpl)),
		daos.NewUserDaoImpl,
		NewDb,
	)
	return nil, nil
}

func wireQuranManagerImpl(csvDir string) (*managers.QuranManagerImpl, error) {
	wire.Build(managers.NewQuranManagerImpl)
	return nil, nil
}
