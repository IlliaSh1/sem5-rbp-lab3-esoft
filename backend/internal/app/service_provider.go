package app

import (
	"github.com/IlliaSh1/backend/configs"
	"gorm.io/gorm"
)

type serviceProvider struct {
	db *gorm.DB

	jwtConfig *configs.JWTConfig
	// jwtService controllers.IJWTService

	// userService controllers.IUserService
	// userRepo    services.IUserRepo
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
