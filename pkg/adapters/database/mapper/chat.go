package mapper

import (
	"github.com/Hossara/linkin-chat/internal/chat/domain"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/models"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"gorm.io/gorm"
)

func ToModelChat(d *domain.ChatRoom) *models.Chat {
	users := utils.Map(d.Users, func(t userDomain.User) models.User {
		return *ToModelUser(&t)
	})

	return &models.Chat{
		Model: gorm.Model{
			CreatedAt: d.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(d.DeletedAt)),
		},
		ID:      uint(d.ID),
		Title:   d.Title,
		Code:    string(d.Code),
		Users:   users,
		OwnerID: uint(d.OwnerID),
		Owner:   *ToModelUser(&d.Owner),
	}
}

func ToDomainChat(m *models.Chat) *domain.ChatRoom {
	users := utils.Map(m.Users, func(t models.User) userDomain.User {
		return *ToDomainUser(&t)
	})

	return &domain.ChatRoom{
		ID:        domain.ChatRoomID(m.ID),
		Title:     m.Title,
		Code:      domain.ChatRoomCode(m.Code),
		Users:     users,
		OwnerID:   userDomain.UserID(m.OwnerID),
		Owner:     *ToDomainUser(&m.Owner),
		CreatedAt: m.CreatedAt,
	}
}
