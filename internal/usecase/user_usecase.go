package usecase

import (
	"time"

	"github.com/brunoseba/event-api/internal/entity"
	"github.com/brunoseba/event-api/internal/repository"
	"github.com/brunoseba/event-api/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserUseCase struct {
	userRepo repository.UserRepositoryImp
}

func NewUserUseCase(userRepo repository.UserRepositoryImp) UserUseCase {
	return UserUseCase{userRepo}
}

func (us *UserUseCase) CreateToken(name string) (string, error) {

	user, err := us.userRepo.GetUserRepo(name)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": user.Nickname,
		"userType": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	})

	tokenString, err := token.SignedString([]byte(utils.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (us *UserUseCase) CreateUserService(user *entity.User) (*entity.UserResponse, string, error) {
	userRes, err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, "", err
	}
	tokenString, _ := us.CreateToken(userRes.Nickname)

	return userRes, tokenString, nil
}

// busca usr por nickname
func (us *UserUseCase) GetUser(name string) (*entity.UserResponse, error) {

	user, err := us.userRepo.GetUserRepo(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserUseCase) GetUserIdService(id string) (*entity.UserResponse, error) {
	user, err := us.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserUseCase) UpdateUserService(id string, user *entity.UserUpdate) (*entity.UserResponse, error) {
	userUpdate, err := us.userRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return userUpdate, nil
}
