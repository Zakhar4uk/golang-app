package user_transport_http

type UsersHTTPHandler struct {
	USerService USerService
}

type USerService interface {
}

func NewUsersHTTPHanlder(USerService USerService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		USerService: USerService,
	}
}
