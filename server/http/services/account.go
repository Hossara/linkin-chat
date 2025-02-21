package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/server/http/mapper"
	"log"
	"time"

	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	jwt2 "github.com/Hossara/linkin-chat/pkg/jwt"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/types"

	userService "github.com/Hossara/linkin-chat/internal/user"
	userPort "github.com/Hossara/linkin-chat/internal/user/port"
)

var (
	ErrUserOnCreate      = userService.ErrUserOnCreate
	ErrUserNotFound      = userService.ErrUserNotFound
	ErrUserAlreadyExists = userService.ErrUserAlreadyExists
	ErrCreatingToken     = errors.New("cannot create token")
)

type AccountService struct {
	svc               userPort.Service
	authSecret        *ecdsa.PrivateKey
	expMin, otpTtlMin uint
}

func NewAccountService(
	svc userPort.Service,
	authSecret *ecdsa.PrivateKey,
	expMin uint,
) *AccountService {
	return &AccountService{
		svc:        svc,
		authSecret: authSecret,
		expMin:     expMin,
	}
}

func AccountServiceGetter(appContainer app.App, cfg config.Server, secret *ecdsa.PrivateKey) helpers.ServiceGetter[*AccountService] {
	return func(ctx context.Context) *AccountService {
		return NewAccountService(
			appContainer.UserService(),
			secret,
			cfg.AuthExpirationMinutes,
		)
	}
}

func (as *AccountService) generateToken(user *domain.User) (string, error) {
	var authExp = time.Now().Add(time.Duration(as.expMin) * time.Minute)

	accessToken, err := jwt2.CreateToken(as.authSecret, jwt2.GenerateUserClaims(user, authExp))

	if err != nil {
		log.Printf("%s: %s", ErrCreatingToken, err)
		return "", ErrCreatingToken
	}

	return accessToken, nil
}

func (as *AccountService) Login(c context.Context, req types.LoginRequest) (*types.AuthResponse, error) {
	user, err := as.svc.GetUserByUsernamePassword(c, req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	token, err := as.generateToken(user)

	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		Token: token,
	}, nil
}

func (as *AccountService) Register(c context.Context, req types.RegisterRequest) (*types.AuthResponse, error) {
	newU := mapper.RegisterRequestToUserDomain(req)

	userId, err := as.svc.CreateUser(c, newU)

	if err != nil {
		return nil, err
	}

	token, err := as.generateToken(&domain.User{
		ID:        userId,
		Username:  newU.Username,
		Password:  newU.Password,
		FirstName: newU.FirstName,
		LastName:  newU.LastName,
	})

	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		Token: token,
	}, nil
}
