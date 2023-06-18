//go:build wireinject
// +build wireinject

// https://github.com/golang/vscode-go/issues/2075
package main

import (
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/favorite"
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
	)
	return nil, nil
}

func wireQuranManagerImpl(csvDir string) (*managers.QuranManagerImpl, error) {
	wire.Build(managers.NewQuranManagerImpl)
	return nil, nil
}

func wireUserManagerImpl() (*managers.UserManagerImpl, error) {
	wire.Build(
		managers.NewUserManagerImpl,
		wire.Bind(new(daos.UserDao), new(*daos.UserDaoImpl)),
		daos.NewUserDaoImpl,
		NewDb,
	)
	return nil, nil
}

func wireFavDaoImpl() (*favorite.FavDaoImpl, error) {
	wire.Build(favorite.NewFavDaoImpl, NewDb)
	return nil, nil
}

func wireFavManagerImpl() (*favorite.FavManagerImpl, error) {
	wire.Build(
		wire.Bind(new(daos.UserDao), new(*daos.UserDaoImpl)),
		wire.Bind(new(favorite.FavDao), new(*favorite.FavDaoImpl)),
		favorite.NewFavManagerImpl,
		daos.NewUserDaoImpl,
		favorite.NewFavDaoImpl,
		NewDb,
	)
	return nil, nil
}
