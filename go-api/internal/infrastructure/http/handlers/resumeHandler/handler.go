package resumeHandler

type Handler struct {
	service Service
}

func NewResumeHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
