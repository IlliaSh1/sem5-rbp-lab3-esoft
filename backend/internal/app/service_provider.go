package app

import (
	"github.com/IlliaSh1/backend/configs"
	repos_mysql_client "github.com/IlliaSh1/backend/internal/repos/mysql/client"
	repos_mysql_deal "github.com/IlliaSh1/backend/internal/repos/mysql/deal"
	repos_mysql_need "github.com/IlliaSh1/backend/internal/repos/mysql/need"
	repos_mysql_offer "github.com/IlliaSh1/backend/internal/repos/mysql/offer"
	repos_mysql_real_estate_object "github.com/IlliaSh1/backend/internal/repos/mysql/real_estate_object"
	repos_mysql_realtor "github.com/IlliaSh1/backend/internal/repos/mysql/realtor"
	"github.com/IlliaSh1/backend/internal/storage/transactor"
	"gorm.io/gorm"
)

type serviceProvider struct {
	db *gorm.DB

	transactor *transactor.Transactor

	jwtConfig *configs.JWTConfig
	// jwtService controllers.IJWTService

	// userService controllers.IUserService
	// userRepo    services.IUserRepo

	clientRepo *repos_mysql_client.ClientRepo

	realtorRepo *repos_mysql_realtor.RealtorRepo

	realEstateObjectRepo *repos_mysql_real_estate_object.RealEstateObjectRepo

	offerRepo *repos_mysql_offer.OfferRepo

	needRepo *repos_mysql_need.NeedRepo

	dealRepo *repos_mysql_deal.DealRepo
	// ...
}

func newServiceProvider(db *gorm.DB, jwtConfig *configs.JWTConfig) *serviceProvider {
	return &serviceProvider{
		db: db,

		jwtConfig: jwtConfig,
	}
}

func (s *serviceProvider) Transactor() *transactor.Transactor {
	if s.transactor == nil {
		s.transactor = transactor.NewTransactor(s.db)
	}
	return s.transactor
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

func (s *serviceProvider) RealtorRepo() *repos_mysql_realtor.RealtorRepo {
	if s.realtorRepo == nil {
		s.realtorRepo = repos_mysql_realtor.NewRealtorRepo(s.db)
	}
	return s.realtorRepo
}

func (s *serviceProvider) RealEstateObjectRepo() *repos_mysql_real_estate_object.RealEstateObjectRepo {
	if s.realEstateObjectRepo == nil {
		s.realEstateObjectRepo = repos_mysql_real_estate_object.NewRealEstateObjectRepo(s.db)
	}
	return s.realEstateObjectRepo
}

func (s *serviceProvider) OfferRepo() *repos_mysql_offer.OfferRepo {
	if s.offerRepo == nil {
		s.offerRepo = repos_mysql_offer.NewOfferRepo(s.db)
	}
	return s.offerRepo
}

func (s *serviceProvider) NeedRepo() *repos_mysql_need.NeedRepo {
	if s.needRepo == nil {
		s.needRepo = repos_mysql_need.NewNeedRepo(s.db)
	}
	return s.needRepo
}

func (s *serviceProvider) DealRepo() *repos_mysql_deal.DealRepo {
	if s.dealRepo == nil {
		s.dealRepo = repos_mysql_deal.NewDealRepo(s.db)
	}
	return s.dealRepo
}
