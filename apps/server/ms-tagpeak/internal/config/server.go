package config

import (
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/controllers"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

var server *echo.Echo

func InitServer() {
	logster.StartFuncLog()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit(dotenv.GetEnv("BODYLIMIT")))
	e.Use(logster.LogsterMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.Use(auth.Principal.Process)

	keycloakMiddleware := &auth.KeycloakMiddleware{}
	e.Pre(keycloakMiddleware.Process)

	logster.Info(fmt.Sprintf("Middlewares: CORST, BodyLimit - %s, Logger, Recover, Secure (XSS), Custome auth keycloak", dotenv.GetEnv("BODYLIMIT")))

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Debug = dotenv.GetEnv("ENV") == "dev"
	logster.Info(fmt.Sprintf("Debug mode: %t", e.Debug))

	e.Server.ReadTimeout = 20 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	server = e
	initRoutes()
	logster.Fatal(e.Start(":8097"), "Error starting server")
}

/**
 * Define all API routes here
 */
func initRoutes() {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	//Example for group routes
	userGroup := server.Group("/auth")
	userGroup.POST("/", controllers.CreateUser)
	userGroup.POST("/migration", controllers.CreateUserFromMigration)
	userGroup.GET("/me", controllers.GetOwnUser)
	userGroup.GET("/:id", controllers.GetUserById)
	userGroup.PATCH("/:id", controllers.UpdateUser)
	userGroup.POST("/reset", controllers.ResetPassword)
	userGroup.POST("/email-verified", controllers.EmailVerified)
	userGroup.POST("/send-verification-email", controllers.SendVerificationEmail)
	userGroup.GET("/email-verification-status", controllers.CheckEmailVerificationStatus)
	userGroup.GET("/:id/referral", controllers.GetAllReferralByUserUuid)
	userGroup.GET("/users", controllers.GetAllUsers)
	userGroup.POST("/:id/profile-picture", controllers.UploadUserProfilePicture)
	userGroup.DELETE("/:id/profile-picture/:pictureUuid", controllers.DeleteUserProfilePicture)
	userGroup.GET("/:id/referral/clicks", controllers.GetReferralsInfoByUserUuid)
	userGroup.GET("/:id/referral/revenue", controllers.GetReferralsRevenueInfoByUserUuid)
	userGroup.GET("/:id/referral/revenue/users-info", controllers.GetUsersReferralsRevenueInfoByReferrerUuid)
	userGroup.GET("/:id/referral/revenue/info", controllers.GetRevenuesInfoByUserUuid)
	userGroup.GET("/:id/referral/revenue/total-revenue", controllers.GetAllReferralRevenueByReferrerUuid)
	userGroup.GET("/:id/stats", controllers.GetUserStats)
	userGroup.GET("/me/check", controllers.CheckUserState)
	userGroup.GET("/valid", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"valid": "true"})
	})
	userGroup.PATCH("/:id/currency", controllers.SetCurrencyUser)
	userGroup.GET("/callback", controllers.AuthCallback)
	userGroup.POST("/socials/finish-profile", controllers.SocialsFinishProfile)

	//Country routes
	countryGroup := server.Group("/countries")
	countryGroup.POST("", controllers.CreateCountry)
	countryGroup.GET("/:id", controllers.GetCountry)
	countryGroup.GET("", controllers.GetCountries)
	countryGroup.DELETE("/:id", controllers.DeleteCountry)
	countryGroup.PATCH("/:id", controllers.PatchCountry)

	//Category routes
	categoryGroup := server.Group("/category")
	categoryGroup.POST("", controllers.CreateCategory)
	categoryGroup.GET("/:id", controllers.GetCategory)
	categoryGroup.GET("", controllers.GetAllCategories)
	categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	categoryGroup.PATCH("/:id", controllers.PatchCategory)

	//Partner routes
	partnerGroup := server.Group("/partner")
	partnerGroup.POST("", controllers.CreatePartner)
	partnerGroup.GET("/:id", controllers.GetPartner)
	partnerGroup.GET("", controllers.GetAllPartners)
	partnerGroup.DELETE("/:id", controllers.DeletePartner)
	partnerGroup.PATCH("/:id", controllers.PatchPartner)

	//Language routes
	languageGroup := server.Group("/language")
	languageGroup.POST("", controllers.CreateLanguage)
	languageGroup.GET("/:id", controllers.GetLanguage)
	languageGroup.GET("", controllers.GetAllLanguages)
	languageGroup.DELETE("/:id", controllers.DeleteLanguage)
	languageGroup.PATCH("/:id", controllers.PatchLanguage)

	//Store routes
	storeGroup := server.Group("/store")
	storeGroup.POST("", controllers.CreateStore)
	storeGroup.POST("/upload-excel", controllers.UploadStoreByExcel)
	storeGroup.GET("/:id", controllers.GetStore)
	storeGroup.POST("/:id/logo", controllers.UploadStoreLogo)
	storeGroup.DELETE("/:id/logo", controllers.DeleteStoreLogo)
	storeGroup.POST("/:id/banner", controllers.UploadStoreBanner)
	storeGroup.DELETE("/:id/banner", controllers.DeleteStoreBanner)
	storeGroup.GET("/:id/partner", controllers.GetAffiliatePartnerByStore)
	storeGroup.GET("/:id/category", controllers.GetCategoryByStore)
	storeGroup.GET("/:id/language", controllers.GetLanguageByStore)
	storeGroup.GET("", controllers.GetAllStores)
	storeGroup.GET("/admin", controllers.GetAllStoresAdmin)
	storeGroup.GET("/:id/country", controllers.GetCountryByStore)
	storeGroup.POST("/:id/country", controllers.AddCountryToStore)
	storeGroup.DELETE("/:id/country", controllers.RemoveCountryFromStore)
	storeGroup.GET("/:id/category", controllers.GetCategoryByStore)
	storeGroup.POST("/:id/category", controllers.AddCategoryToStore)
	storeGroup.DELETE("/:id/category", controllers.RemoveCategoryFromStore)
	storeGroup.DELETE("/:id", controllers.DeleteStore)
	storeGroup.PATCH("/:id", controllers.PatchStore)
	storeGroup.GET("/export-csv", controllers.ExportStoreCSV)
	storeGroup.GET("/:id/redirect", controllers.GetStoreRedirectUrl)
	storeGroup.GET("/approvals", controllers.GetStoresForApproval)
	storeGroup.GET("/position/:position", controllers.IsUniquePosition)

	//Store Visit Routes
	storeVisitGroup := server.Group("/store-visit")
	storeVisitGroup.POST("", controllers.CreateStoreVisit)
	storeVisitGroup.GET("/:id", controllers.GetStoreVisit)
	storeVisitGroup.GET("/:id/store", controllers.GetStoreByStoreVisit)
	storeVisitGroup.POST("/reference/:reference", controllers.ValidateReference)
	storeVisitGroup.GET("", controllers.GetAllStoreVisits)
	storeVisitGroup.GET("/admin", controllers.GetStoreVisitsAdmin)
	storeVisitGroup.GET("/user", controllers.GetAllStoreVisitsByUserUUID)
	storeVisitGroup.DELETE("/:id", controllers.DeleteStoreVisit)
	storeVisitGroup.PATCH("/:id", controllers.PatchStoreVisit)
	storeVisitGroup.GET("/stores", controllers.GetDistinctStoresVisitedByUserUUID)

	//Currency Exchange Rate Routes
	currencyExchangeRateGroup := server.Group("/currency-exchange-rate")
	currencyExchangeRateGroup.POST("", controllers.CreateCurrencyExchangeRate)
	currencyExchangeRateGroup.GET("/:id", controllers.GetCurrencyExchangeRate)
	currencyExchangeRateGroup.GET("", controllers.GetAllCurrencyExchangeRates)
	currencyExchangeRateGroup.DELETE("/:id", controllers.DeleteCurrencyExchangeRate)
	currencyExchangeRateGroup.PATCH("/:id", controllers.PatchCurrencyExchangeRate)

	//Transaction Routes
	transactionGroup := server.Group("/transaction")
	transactionGroup.POST("", controllers.CreateTransaction)
	transactionGroup.GET("", controllers.GetAllTransactions)
	transactionGroup.GET("/me", controllers.GetMyTransactions)
	transactionGroup.GET("/:id", controllers.GetTransaction)
	transactionGroup.DELETE("/:id", controllers.DeleteTransaction)
	transactionGroup.PATCH("/:id", controllers.PatchTransaction)
	transactionGroup.GET("/:id/reward", controllers.GetRewardTransaction)
	transactionGroup.GET("/:id/store-visit", controllers.GetStoreVisitByTransaction)
	transactionGroup.GET("/:id/store", controllers.GetStoreByTransaction)
	transactionGroup.GET("/:id/currency-exchange-rate", controllers.GetCurrencyExchangeRateByTransaction)
	transactionGroup.PATCH("/bulk/edit", controllers.BulkEditTransaction)

	//Payment Method Routes
	paymentMethodGroup := server.Group("/payment-method")
	paymentMethodGroup.POST("", controllers.CreatePaymentMethod)
	paymentMethodGroup.GET("/:id", controllers.GetPaymentMethod)
	paymentMethodGroup.GET("", controllers.GetAllPaymentMethods)
	paymentMethodGroup.DELETE("/:id", controllers.DeletePaymentMethod)
	paymentMethodGroup.PATCH("/:id", controllers.PatchPaymentMethod)

	//Withdrawal Routes
	withdrawalGroup := server.Group("/withdrawal")
	withdrawalGroup.POST("", controllers.CreateWithdrawal)
	withdrawalGroup.GET("/:id", controllers.GetWithdrawal)
	withdrawalGroup.GET("", controllers.GetAllWithdrawals)
	withdrawalGroup.GET("/me", controllers.GetMyWithdrawals)
	withdrawalGroup.GET("/me/stats", controllers.GetMyWithdrawalsStats)
	withdrawalGroup.DELETE("/:id", controllers.DeleteWithdrawal)
	withdrawalGroup.PATCH("/:id", controllers.PatchWithdrawal)
	withdrawalGroup.PATCH("/bulk", controllers.BulkUpdateStateWithdrawal)

	//Reward Routes
	rewardGroup := server.Group("/reward")
	rewardGroup.POST("", controllers.CreateReward)
	rewardGroup.GET("/:id", controllers.GetReward)
	rewardGroup.GET("", controllers.GetAllRewards)
	rewardGroup.DELETE("/:id", controllers.DeleteReward)
	rewardGroup.PATCH("/:id", controllers.PatchReward)
	rewardGroup.GET("/:id/transaction", controllers.GetTransactionByReward)
	rewardGroup.GET("/:id/currency-exchange-rate", controllers.GetCurrencyExchangeRateByReward)
	rewardGroup.PATCH("/:id/stop", controllers.StopReward)
	rewardGroup.PATCH("/:id/finish", controllers.FinishReward)
	rewardGroup.POST("/:id/verify", controllers.VerifyReward)
	rewardGroup.PATCH("/bulk/edit", controllers.RewardBulkEdit)
	rewardGroup.POST("/bulk/create", controllers.CreateRewardBulk)
	rewardGroup.GET("/:userUUID/sum-live", controllers.GetSumRewardsLiveByUserUuid)

	//Reward History Routes
	rewardGroup.GET("/:id/history", controllers.GetRewardHistory)
	rewardGroup.GET("/:id/history/graph", controllers.GetRewardHistoryGraph)
	rewardGroup.POST("/:id/history", controllers.CreateRewardHistory)

	userPaymentMethod := server.Group("/user/payment-method")
	userPaymentMethod.POST("", controllers.CreateUserPaymentMethod)
	userPaymentMethod.GET("", controllers.GetUserPaymentMethodsByUserUuid)
	userPaymentMethod.GET("/:id", controllers.GetUserPaymentMethodById)
	userPaymentMethod.DELETE("/:id", controllers.DeleteUserPaymentMethod)
	userPaymentMethod.POST("/file", controllers.UploadIbanStatement)
	userPaymentMethod.GET("/file/:id", controllers.DownloadFile)

	// Public Routes
	publicGroup := server.Group("/public")
	publicGroup.GET("/ip", controllers.GetInfoFromIp)
	publicGroup.GET("/store", controllers.GetPublicAllStores)
	publicGroup.GET("/store/:id", controllers.GetPublicStore)
	publicGroup.GET("/countries", controllers.GetPublicCountries)
	publicGroup.GET("/category", controllers.GetPublicCategories)
	publicGroup.GET("/countries/code/:code", controllers.GetCountryByCode)
	publicGroup.GET("/category/code/:code", controllers.GetCategoryByCode)
	publicGroup.GET("/auth/validate-action", controllers.ValidateUserAction)
	//publicGroup.POST("/migrated-users-balance", controllers.SetMigratedUserBalance)
	publicShopifyGroup := publicGroup.Group("/shopify")
	publicShopifyGroup.GET("/shop/:url/exists", controllers.CheckIfShopExists)
	publicShopifyGroup.POST("/shop/auth", controllers.CreateUserShopifyShop)

	//Aux Routes
	auxGroup := server.Group("/aux")
	auxGroup.GET("/logo", controllers.GetLogo)
	auxGroup.POST("/vat", controllers.CheckVatValidity)
	auxGroup.POST("/upload-file", controllers.UploadFile)

	//Configuration Routes
	configurationGroup := server.Group("/configuration")
	//configurationGroup.POST("", controllers.CreateConfiguration)
	configurationGroup.GET("", controllers.GetConfigurations)
	configurationGroup.GET("/latest", controllers.GetLatestConfigurations)
	configurationGroup.GET("/:code", controllers.GetConfiguration)
	configurationGroup.PATCH("/:id", controllers.UpdateConfiguration)
	//configurationGroup.DELETE("/:id", controllers.DeleteConfiguration)

	//Referral Routes
	referralGroup := server.Group("/referral")
	referralGroup.POST("/click", controllers.CreateReferralClick)

	dashboardGroup := server.Group("/dashboard")
	dashboardGroup.GET("", controllers.GetValuesDashboard)
	dashboardGroup.GET("/statistics", controllers.GetStatisticsByMonth)
	dashboardGroup.GET("/transactions", controllers.GetTransactionsDashboard)
	dashboardGroup.GET("/rewards/currencies/count", controllers.GetRewardCountByCurrencies)

	//Admin processes
	adminGroup := server.Group("/admin")
	adminGroup.POST("/reward-history", controllers.RunRewardHistory)

	shopifyGroup := server.Group("/shopify")
	shopifyGroup.GET("/valid", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"valid": "true"})
	})
	shopifyGroup.Any("/relay/*", controllers.RelayMSShopify)
	shopifyGroup.POST("", controllers.CreateShopifyStore)
	shopifyGroup.GET("/user/:id", controllers.GetShopifyStoreByUser)
	shopifyGroup.GET("/shop/:id", controllers.GetShopifyStore)
	shopifyGroup.GET("/shop/owner", controllers.CheckIfUserIsOwnerOfShop)
	shopifyGroup.POST("/shop/:id", controllers.UpdateShopifyStore)
	shopifyGroup.GET("/shop/:id/stats", controllers.GetShopifyStoreStats)
	shopifyGroup.GET("/store/:id/transaction", controllers.GetAllShopTransactions)
	shopifyGroup.PATCH("/store/bulk/edit", controllers.BulkEditTransaction)

	server.POST("/bug-report", controllers.CreateBugReport)

	notificationsGroup := server.Group("/notifications")
	notificationsGroup.POST("/token", controllers.SaveUserFirebaseToken)
	notificationsGroup.DELETE("/token", controllers.DeleteFirebaseToken)
	notificationsGroup.GET("/topics", controllers.GetTopics)
}
