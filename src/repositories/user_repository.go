package repositories

import (
	"bm-support/config/database"
	dto "bm-support/src/dto"
	"bm-support/src/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserRepository interface {
	Store(dto dto.UserRegisterDTO) (models.User, error)
	Index() (models.User, error)
	GetByID(userID uint64) (models.User, error)
	Login(dto dto.LoginDTO) (string, uint64, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repository userRepository) Store(dto dto.UserRegisterDTO) (models.User, error) {
	var user models.User
	user.Email = dto.Email
	// user.FirstName = models.JSONB{
	// 	"ar": "محمود",
	// 	"en": "Mahmoud",
	// }
	// user.LastName = models.JSONB{
	// 	"ar": "أحمد",
	// 	"en": "Ahmed",
	// }
	user.SetPassword("123456")

	if err := database.DB.Create(&user).Error; err != nil {
		fmt.Println(err)
	}
	return user, nil
}

func (repository userRepository) Index() (models.User, error) {
	var user models.User
	if err := database.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
	}
	return user, nil
}

func (repository userRepository) Login(dto dto.LoginDTO) (string, uint64, error) {
	var user models.User
	if err := database.DB.Where("email=?", dto.Email).Find(&user).Error; err != nil {
		return "", 0, errors.New("bad request")
	}

	if user.ID == 0 {
		return "", 0, errors.New("unauthorized")
	}

	if err := user.ComparePassword(dto.Password); err != nil {
		return "", 0, errors.New("unauthorized")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", 0, errors.New("bad request")
	}

	return tokenString, user.ID, nil
}

func (repository userRepository) GetByID(userID uint64) (models.User, error) {
	var user models.User
	if err := database.DB.Where("id=?", userID).Find(&user).Error; err != nil {
		return user, errors.New("bad request")
	}

	if user.ID == 0 {
		return user, errors.New("not found")
	}

	return user, nil
}
