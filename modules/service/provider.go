package service

import (
	"go-web-template/modules/repository"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

// All Services
var ServiceSet = wire.NewSet(
	userServiceSet,
	contentServiceSet,
)

// UserService
var userServiceSet = wire.NewSet(
	wire.Bind(new(IUserService), new(UserService)),
	UserServiceProvider,
)

func UserServiceProvider(
	mySQLGorm *repository.MySQLGorm,
	cryptTool crypt.PasswordCrypt,
	jwtManager jwt.JwtManager,
) UserService {
	return UserService{
		MySQLGorm:  mySQLGorm,
		CryptTool:  cryptTool,
		JwtManager: jwtManager,
	}
}

// ContentService
var contentServiceSet = wire.NewSet(
	wire.Bind(new(IContentService), new(ContentService)),
	ContentServiceProvider,
)

func ContentServiceProvider() ContentService {
	return ContentService{}
}
