package main

import (
	"log"
	"time"

	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/boot"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	authClient "github.com/ppopgi-pang/ppopgipang-spine/auth/client"
	authController "github.com/ppopgi-pang/ppopgipang-spine/auth/controller"
	authInterceptor "github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	authRoutes "github.com/ppopgi-pang/ppopgipang-spine/auth/routes"
	authService "github.com/ppopgi-pang/ppopgipang-spine/auth/service"

	careerEntity "github.com/ppopgi-pang/ppopgipang-spine/careers/entities"
	certificationEntity "github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	gamificationEntity "github.com/ppopgi-pang/ppopgipang-spine/gamification/entities"
	moderationEntity "github.com/ppopgi-pang/ppopgipang-spine/moderation/entities"
	notificationEntity "github.com/ppopgi-pang/ppopgipang-spine/notifications/entities"
	proposalEntity "github.com/ppopgi-pang/ppopgipang-spine/proposals/entities"
	reviewEntity "github.com/ppopgi-pang/ppopgipang-spine/reviews/entities"
	storeEntity "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	tradeEntity "github.com/ppopgi-pang/ppopgipang-spine/trades/entities"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"

	careerRepository "github.com/ppopgi-pang/ppopgipang-spine/careers/repository"
	certificationRepository "github.com/ppopgi-pang/ppopgipang-spine/certifications/repository"
	gamificationRepository "github.com/ppopgi-pang/ppopgipang-spine/gamification/repository"
	moderationRepository "github.com/ppopgi-pang/ppopgipang-spine/moderation/repository"
	notificationRepository "github.com/ppopgi-pang/ppopgipang-spine/notifications/repository"
	proposalRepository "github.com/ppopgi-pang/ppopgipang-spine/proposals/repository"
	reviewRepository "github.com/ppopgi-pang/ppopgipang-spine/reviews/repository"
	storeRepository "github.com/ppopgi-pang/ppopgipang-spine/stores/repository"
	tradeRepository "github.com/ppopgi-pang/ppopgipang-spine/trades/repository"
	userRepository "github.com/ppopgi-pang/ppopgipang-spine/users/repository"

	careerService "github.com/ppopgi-pang/ppopgipang-spine/careers/service"
	certificationService "github.com/ppopgi-pang/ppopgipang-spine/certifications/service"
	gamificationService "github.com/ppopgi-pang/ppopgipang-spine/gamification/service"
	moderationService "github.com/ppopgi-pang/ppopgipang-spine/moderation/service"
	notificationService "github.com/ppopgi-pang/ppopgipang-spine/notifications/service"
	proposalService "github.com/ppopgi-pang/ppopgipang-spine/proposals/service"
	reviewService "github.com/ppopgi-pang/ppopgipang-spine/reviews/service"
	storeService "github.com/ppopgi-pang/ppopgipang-spine/stores/service"
	tradeService "github.com/ppopgi-pang/ppopgipang-spine/trades/service"
	userService "github.com/ppopgi-pang/ppopgipang-spine/users/service"

	careerController "github.com/ppopgi-pang/ppopgipang-spine/careers/controller"
	certificationController "github.com/ppopgi-pang/ppopgipang-spine/certifications/controller"
	gamificationController "github.com/ppopgi-pang/ppopgipang-spine/gamification/controller"
	moderationController "github.com/ppopgi-pang/ppopgipang-spine/moderation/controller"
	notificationController "github.com/ppopgi-pang/ppopgipang-spine/notifications/controller"
	proposalController "github.com/ppopgi-pang/ppopgipang-spine/proposals/controller"
	reviewController "github.com/ppopgi-pang/ppopgipang-spine/reviews/controller"
	storeController "github.com/ppopgi-pang/ppopgipang-spine/stores/controller"
	tradeController "github.com/ppopgi-pang/ppopgipang-spine/trades/controller"
	userController "github.com/ppopgi-pang/ppopgipang-spine/users/controller"

	careerRoutes "github.com/ppopgi-pang/ppopgipang-spine/careers/routes"
	certificationRoutes "github.com/ppopgi-pang/ppopgipang-spine/certifications/routes"
	gamificationRoutes "github.com/ppopgi-pang/ppopgipang-spine/gamification/routes"
	moderationRoutes "github.com/ppopgi-pang/ppopgipang-spine/moderation/routes"
	notificationRoutes "github.com/ppopgi-pang/ppopgipang-spine/notifications/routes"
	proposalRoutes "github.com/ppopgi-pang/ppopgipang-spine/proposals/routes"
	reviewRoutes "github.com/ppopgi-pang/ppopgipang-spine/reviews/routes"
	storeRoutes "github.com/ppopgi-pang/ppopgipang-spine/stores/routes"
	tradeRoutes "github.com/ppopgi-pang/ppopgipang-spine/trades/routes"
	userRoutes "github.com/ppopgi-pang/ppopgipang-spine/users/routes"

	"github.com/joho/godotenv"
)

func NewDB() *gorm.DB {
	dsn := "root:test1234@tcp(127.0.0.1:3306)/ppopgipang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("데이터베이스 연결 실패: " + err.Error())
	}

	db.AutoMigrate(
		&userEntity.LootTag{},
		&userEntity.User{},
		&userEntity.UserLoot{},
		&userEntity.UserProgress{},
		&userEntity.UserSearchHistory{},
		&tradeEntity.Trade{},
		&tradeEntity.TradeChatMessage{},
		&tradeEntity.TradeChatRoom{},
		&storeEntity.Store{},
		&storeEntity.StoreAnalytics{},
		&storeEntity.StoreFacility{},
		&storeEntity.StoreOpeningHour{},
		&storeEntity.StorePhoto{},
		&storeEntity.StoreType{},
		&reviewEntity.Review{},
		&proposalEntity.Proposal{},
		&notificationEntity.Notification{},
		&notificationEntity.PushSubscription{},
		&moderationEntity.ContentReport{},
		&moderationEntity.ModerationAction{},
		&gamificationEntity.Achievement{},
		&gamificationEntity.Stamp{},
		&gamificationEntity.UserAchievement{},
		&gamificationEntity.UserStamp{},
		&certificationEntity.Certification{},
		&certificationEntity.CertificationPhoto{},
		&certificationEntity.CertificationReason{},
		&certificationEntity.CertificationTag{},
		&certificationEntity.CheckinReasonPreset{},
		&certificationEntity.LootCommentPreset{},
		&certificationEntity.LootLike{},
		&certificationEntity.UserStoreStat{},
		&careerEntity.Application{},
		&careerEntity.JobPosting{},
	)

	log.Println("[Database] MySQL 연결 성공")
	return db
}

func main() {
	_ = godotenv.Load()

	app := spine.New()

	app.Constructor(
		NewDB,
		careerRepository.NewCareerRepository,
		certificationRepository.NewCertificationRepository,
		gamificationRepository.NewGamificationRepository,
		moderationRepository.NewModerationRepository,
		notificationRepository.NewNotificationRepository,
		proposalRepository.NewProposalRepository,
		reviewRepository.NewReviewRepository,
		storeRepository.NewStoreRepository,
		tradeRepository.NewTradeRepository,
		userRepository.NewUserRepository,
		careerService.NewCareerService,
		certificationService.NewCertificationService,
		gamificationService.NewGamificationService,
		moderationService.NewModerationService,
		notificationService.NewNotificationService,
		proposalService.NewProposalService,
		reviewService.NewReviewService,
		storeService.NewStoreService,
		tradeService.NewTradeService,
		userService.NewUserService,
		careerController.NewCareerController,
		certificationController.NewCertificationController,
		gamificationController.NewGamificationController,
		moderationController.NewModerationController,
		notificationController.NewNotificationController,
		proposalController.NewProposalController,
		reviewController.NewReviewController,
		storeController.NewStoreController,
		tradeController.NewTradeController,
		userController.NewUserController,
		authClient.NewKakaoOAuthClient,
		authInterceptor.NewKakaoAuthCallbackInterceptor,
		authController.NewAuthController,
		authService.NewAuthService,
	)

	// 커리어 라우트 등록
	careerRoutes.RegisterUserRoutes(app)
	// 인증 라우트 등록
	certificationRoutes.RegisterCertificationRoutes(app)
	// 게이미피케이션 라우트 등록
	gamificationRoutes.RegisterGamificationRoutes(app)
	// 신고 라우트 등록
	moderationRoutes.RegisterModerationRoutes(app)
	// 알림 라우트 등록
	notificationRoutes.RegisterNotificationRoutes(app)
	// 제안 라우트 등록
	proposalRoutes.RegisterProposalRoutes(app)
	// 리뷰 라우트 등록
	reviewRoutes.RegisterReviewRoutes(app)
	// 가게 라우트 등록
	storeRoutes.RegisterStoreRoutes(app)
	// 중고거래 라우트 등록
	tradeRoutes.RegisterTradeRoutes(app)
	// 유저 라우트 등록
	userRoutes.RegisterUserRoutes(app)
	// 인증 라우트 등록
	authRoutes.RegisterUserRoutes(app)

	app.Run(boot.Options{
		Address:                ":8080",
		EnableGracefulShutdown: true,
		ShutdownTimeout:        10 * time.Second,
	})
}
