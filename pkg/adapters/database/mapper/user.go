package mapper

import (
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/models"
	"github.com/Hossara/linkin-chat/pkg/hash"
	"gorm.io/gorm"
)

func ToDomainUser(m *models.User) *domain.User {
	if m == nil {
		return nil
	}

	return &domain.User{
		ID:        domain.UserID(m.ID),
		Username:  m.Username,
		Password:  m.Password,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		CreatedAt: m.CreatedAt,
	}
}

func ToModelUser(d *domain.User) *models.User {
	if d == nil {
		return nil
	}

	bcrypt := hash.NewBcryptHasher()

	password, err := bcrypt.HashPassword(d.Password)

	if err != nil {
		return nil
	}

	return &models.User{
		Model: gorm.Model{
			ID:        uint(d.ID),
			CreatedAt: d.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(d.DeletedAt)),
		},

		ID:        uint(d.ID),
		Username:  d.Username,
		Password:  password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
}
