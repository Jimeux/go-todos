package auth

import (
	"gin-todos/app/user"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"errors"
)

const TokenKey = "user-token:"

type Service interface {
	Authenticate(username string, password string) (string, error)
	VerifyToken(token string) (*user.Model, error)
}

func NewService(cache *redis.Pool, userRepository user.Repository) Service {
	return &ServiceImpl{cache, userRepository}
}

type ServiceImpl struct {
	cache          *redis.Pool
	userRepository user.Repository
}

func (s *ServiceImpl) Authenticate(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("get out")
	}

	userModel, err := s.userRepository.FindByCredentials(username, password)

	if err != nil {
		return "", err
	} else if userModel == nil {
		return "", errors.New("get out")
	}


	token := generateToken(userModel)
	key := TokenKey + token
	marsh, _ := json.Marshal(userModel)
	s.cache.Get().Do("SET", key, marsh)
	return token, nil
}

func (s *ServiceImpl) VerifyToken(token string) (*user.Model, error) {
	key := TokenKey + token

	reply, err := redis.String(s.cache.Get().Do("GET", key))

	if err != nil {
		return nil, err
	}

	u := new(user.Model)

	jerr := json.Unmarshal([]byte(reply), u)

	if jerr != nil {
		return nil, jerr
	}

	return u, nil
}

func generateToken(u *user.Model) string {
	return "aToken" + u.Username
}
