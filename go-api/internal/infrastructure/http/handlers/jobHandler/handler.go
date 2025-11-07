package jobHandler

type Handler struct {
	service Service
}

func NewJobHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
