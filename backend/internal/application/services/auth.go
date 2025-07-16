package app_services

import (
	dtos "app/backend/internal"
	"app/backend/internal/domain/entities"
	env "app/backend/internal/infra"
	"app/backend/internal/infra/repositories"
	infra_services "app/backend/internal/infra/services"
	"errors"
)

type userInfo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type authInfo struct {
	Token string `json:"token"`
}

type loginResult struct {
	UserInfo userInfo `json:"user"`
	AuthInfo authInfo `json:"auth"`
}

type IAuthService interface {
	Register(data *dtos.RegisterDto) error
	Authenticate(data *dtos.LoginDto) (*loginResult, error)
}

var _ IAuthService = (*AuthService)(nil)

type AuthService struct {
	userRepository repositories.IUserRepository
	tokenService   infra_services.ITokenService
	hashService    infra_services.IHashService
}

func NewAuthService(
	ur repositories.IUserRepository,
	ts infra_services.ITokenService,
	hs infra_services.IHashService,
) *AuthService {
	return &AuthService{
		userRepository: ur,
		tokenService:   ts,
		hashService:    hs,
	}
}

func (s *AuthService) Register(data *dtos.RegisterDto) error {
	for _, check := range []struct {
		value string
		fn    func(string) *entities.User
	}{
		{data.Username, s.userRepository.FindByUsername},
		{data.Email, s.userRepository.FindByEmail},
	} {
		if check.fn(check.value) != nil {
			return errors.New("user already exists")
		}
	}

	ps, err := s.hashService.Hash(data.Password)
	if err != nil {
		return err
	}

	user, err := entities.NewUser(data.Username, data.Email, ps)
	if err != nil {
		return err
	}

	s.userRepository.Save(user)

	return nil
}

func (s *AuthService) Authenticate(data *dtos.LoginDto) (*loginResult, error) {
	user := s.userRepository.FindByEmail(data.Login)
	if user == nil {
		user = s.userRepository.FindByUsername(data.Login)
	}

	if user == nil {
		return nil, errors.New("invalid credentials [user not found]")
	}

	isValid := s.hashService.Compare(data.Password, user.Password)

	if !isValid {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.tokenService.Generate(user.Id, env.ACCESS_TOKEN_EXP)
	if err != nil {
		return nil, err
	}

	return &loginResult{
		UserInfo: userInfo{
			Id:       user.Id,
			Username: user.Username,
		},
		AuthInfo: authInfo{
			Token: token,
		},
	}, nil
}
