package user

type Service interface {
	Me(userID string) *MeResponse
}

type service struct {}

func NewService () Service {
	return &service{}
}

func (s *service) Me(userID string) *MeResponse {
	return &MeResponse{
		UserID: userID,
	}
}