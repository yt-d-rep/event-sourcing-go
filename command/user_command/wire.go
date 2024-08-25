//go:build wireinject

package usercommand

import (
	"github.com/google/wire"

	uRepo "yt-d-rep/github.com/event-sourcing-go/infrastructure/repository/user"
	"yt-d-rep/github.com/event-sourcing-go/infrastructure/store"
)

func Wire() *handler {
	panic(wire.Build(
		UserCommandProviderSet,
		uRepo.UserRepositoryProvideSet,
		store.StoreProviderSet,
	))
}
