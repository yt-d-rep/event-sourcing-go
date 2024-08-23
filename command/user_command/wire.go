//go:build wireinject

package usercommand

import (
	"github.com/google/wire"

	uRepo "yt-d-rep/github.com/event-sourcing-go/infrastructure/repository/user"
)

func Wire() *handler {
	panic(wire.Build(
		UserCommandProviderSet,
		uRepo.UserRepositoryProvideSet,
	))
}
