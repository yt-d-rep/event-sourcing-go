package usercommand

type (
	Handler interface {
		Create() error
	}
	handler struct {
		usecase Usecase
	}
)

func (h *handler) Create() error {
	return h.usecase.Create()
}
