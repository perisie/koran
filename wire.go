//go:build wireinject
// +build wireinject

// https://github.com/golang/vscode-go/issues/2075
package main

import (
	"github.com/arikama/koran-backend/managers"
	"github.com/google/wire"
)

func InitializeQuranManagerImpl(csvDir string) (*managers.QuranManagerImpl, error) {
	wire.Build(managers.NewQuranManagerImpl)
	return &managers.QuranManagerImpl{}, nil
}

func InitializeGoogleAuthManagerImpl(
	clientId managers.ClientId,
	clientSecret managers.ClientSecret,
	redirectUrl managers.RedirectUrl,
) (*managers.GoogleAuthManagerImpl, error) {
	wire.Build(managers.NewGoogleAuthManagerImpl)
	return &managers.GoogleAuthManagerImpl{}, nil
}
