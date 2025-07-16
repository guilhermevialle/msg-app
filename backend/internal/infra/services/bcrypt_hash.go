package infra_services

import (
	"golang.org/x/crypto/bcrypt"
)

type IHashService interface {
	Hash(plain string) (string, error)
	Compare(plain string, hash string) bool
}

type HashService struct{}

var _ IHashService = (*HashService)(nil)

func NewBcryptHashService() *HashService {
	return &HashService{}
}

func (h *HashService) Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash), err
}

func (h *HashService) Compare(plain string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
