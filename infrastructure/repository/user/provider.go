package user

import (
	"sync"

	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
	"yt-d-rep/github.com/event-sourcing-go/infrastructure/store"

	"github.com/google/wire"
)

var (
	uRepo     *userRepository
	uRepoOnce sync.Once

	UserRepositoryProvideSet wire.ProviderSet = wire.NewSet(
		ProvideUserRepository,
		// Bind the interface to the implementation
		wire.Bind(new(domainUser.UserRepository), new(*userRepository)),
	)
)

func ProvideUserRepository(writer store.Writer) *userRepository {
	uRepoOnce.Do(func() {
		uRepo = &userRepository{
			writer: writer,
		}
	})
	return uRepo
}
