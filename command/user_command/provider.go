package usercommand

import (
	"sync"

	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"

	"github.com/google/wire"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	uscs                   *usecase
	uscsOnce               sync.Once
	UserCommandProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideUsecase,
		// Bind the interface to the implementation
		wire.Bind(new(Handler), new(*handler)),
		wire.Bind(new(Usecase), new(*usecase)),
	)
)

func ProvideHandler(usecase Usecase) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			usecase: usecase,
		}
	})
	return hdl
}

func ProvideUsecase(uRepo domainUser.UserRepository) *usecase {
	uscsOnce.Do(func() {
		uscs = &usecase{
			uRepo: uRepo,
		}
	})
	return uscs
}
