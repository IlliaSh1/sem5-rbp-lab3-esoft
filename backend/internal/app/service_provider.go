package app

import (
	"github.com/IlliaSh1/backend/configs"
	repos_mysql_client "github.com/IlliaSh1/backend/internal/repos/mysql/client"
	"gorm.io/gorm"
)

type serviceProvider struct {
	db *gorm.DB

	jwtConfig *configs.JWTConfig
	// jwtService controllers.IJWTService

	// userService controllers.IUserService
	// userRepo    services.IUserRepo

	clientRepo *repos_mysql_client.ClientRepo
	// ...
}

func newServiceProvider(db *gorm.DB, jwtConfig *configs.JWTConfig) *serviceProvider {
	return &serviceProvider{
		db: db,

		jwtConfig: jwtConfig,
	}
}

// func (s *ServiceProvider) JWTService() controllers.IJWTService {
// 	if s.jwtService == nil {
// 		s.jwtService = services_jwt.NewJWTService(s.jwtConfig)
// 	}

// 	return s.jwtService
// }

// func (s *ServiceProvider) AuthMiddleware() controllers.IAuthMiddleware {
// 	return middlewares_auth.NewAuthMiddleware(s.JWTService())
// }

// func (s *ServiceProvider) UserService() controllers.IUserService {
// 	if s.userService == nil {
// 		s.userService = services_user.NewUserService(s.UserRepo())
// 	}
// 	return s.userService
// }

// func (s *ServiceProvider) UserRepo() services.IUserRepo {
// 	if s.userRepo == nil {
// 		s.userRepo = repos_mysql_user.NewUserRepo(s.db)
// 	}
// 	return s.userRepo
// }

func (s *serviceProvider) ClientRepo() *repos_mysql_client.ClientRepo {
	if s.clientRepo == nil {
		s.clientRepo = repos_mysql_client.NewClientRepo(s.db)
	}
	return s.clientRepo
}
