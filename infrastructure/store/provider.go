package store

import (
	"sync"

	"github.com/google/wire"
)

var (
	wrtr     *writer
	wrtrOnce sync.Once

	StoreProviderSet = wire.NewSet(
		ProvideWriter,
		wire.Bind(new(Writer), new(*writer)),
	)
)

func ProvideWriter() *writer {
	wrtrOnce.Do(func() {
		w, err := newWriter()
		if err != nil {
			panic(err)
		}
		wrtr = w
	})
	return wrtr
}
