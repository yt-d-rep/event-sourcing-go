package user

import (
	"sync"

	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"

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

func ProvideUserRepository() *userRepository {
	uRepoOnce.Do(func() {
		uRepo = &userRepository{}
	})
	return uRepo
}
