package repositories

import (
	"bm-support/config/database"
	"bm-support/src/models"
	"errors"
)

type TokenRepository interface {
	Store(token string, userID uint64) error
	GetByToken(token string) (models.Token, error)
	Delete(userID uint64) error
}

type tokenRepository struct{}

func NewTokenRepository() TokenRepository {
	return &tokenRepository{}
}

func (repository tokenRepository) Store(token string, userID uint64) error {
	var tokenModel models.Token
	tokenModel.Token = token
	tokenModel.UserID = userID

	if err := database.DB.Create(&tokenModel).Error; err != nil {
		return errors.New("bad request")
	}
	return nil
}

func (repository tokenRepository) Delete(userID uint64) error {
	var tokenModel models.Token

	if err := database.DB.Where("user_id=?", userID).Delete(&tokenModel).Error; err != nil {
		return errors.New("bad request")
	}

	return nil
}

func (repository tokenRepository) GetByToken(token string) (models.Token, error) {
	var tokenModel models.Token
	if err := database.DB.Where("token=?", token).Find(&tokenModel).Error; err != nil {
		return tokenModel, errors.New("bad request")
	}

	if tokenModel.ID == 0 {
		return tokenModel, errors.New("unauthorized")
	}

	return tokenModel, nil
}
