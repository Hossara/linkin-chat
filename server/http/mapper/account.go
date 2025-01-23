package mapper

import (
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/server/http/types"
)

func RegisterRequestToUserDomain(req types.RegisterRequest) *domain.User {
	return &domain.User{
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
}
