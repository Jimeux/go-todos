package auth

import (
	"gin-todos/app/user"
	"encoding/json"
	"errors"
	"gin-todos/app"
)

const TokenKey = "user-token:"

type Service interface {
	Authenticate(username string, password string) (string, error)
	VerifyToken(token string) (*user.Model, error)
	RevokeToken(token string) error
}

func NewService(cache app.Cache, userRepository user.Repository) Service {
	return &ServiceImpl{cache, userRepository}
}

type ServiceImpl struct {
	cache          app.Cache
	userRepository user.Repository
}

func (s *ServiceImpl) Authenticate(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("parameters cannot be empty")
	}

	userModel, err := s.userRepository.FindByCredentials(username, password)

	if err != nil {
		return "", err
	} else if userModel == nil {
		return "", errors.New("credentials do not exist")
	}

	token := generateToken(userModel)
	key := TokenKey + token
	marsh, _ := json.Marshal(userModel)
	cacheErr := s.cache.Set(key, string(marsh))

	if cacheErr != nil {
		return "", cacheErr
	}

	return token, nil
}

func (s *ServiceImpl) VerifyToken(token string) (*user.Model, error) {
	key := TokenKey + token
	reply, err := s.cache.Get(key)

	if err != nil {
		return nil, err
	}

	userModel := new(user.Model)
	jsonErr := json.Unmarshal([]byte(reply), userModel)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return userModel, nil
}

func (s *ServiceImpl) RevokeToken(token string) error {
	key := TokenKey + token
	return s.cache.Del(key)
}

func generateToken(u *user.Model) string {
	return "aToken" + u.Username
}
