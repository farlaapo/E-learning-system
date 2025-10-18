package main

import (
	"database/sql"
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/gateway"
	"e-learning-system/internal/api/routes"
	"e-learning-system/internal/config"
	"e-learning-system/internal/domain/service"
	"fmt"

	// utils "kaabe-app/pkg/config"

	"log"
	// "net/http"
	// "os"
	// "time"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, relying on system environment variables")
	}
}

func main() {
	// Load configuration
	config.LoadEnv()
	appCfg, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
	}
	

	dbCfg := config.LoadDBConfig()
	db := config.InitDB(dbCfg)
	if db == nil {
		log.Fatal("Failed to initialize the database")
	}

	
	//---MIGRATIONS ----
	DB  := dbCfg.GetDatabaseURL()
	m, err := migrate.New(
		"file://migrations", DB,
	)
	if err != nil {
		log.Fatalf("migrations failed: %v", err)
	}

	if err := m.Up(); err != nil  && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("database migrated successfully")
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()

	
	fmt.Printf("Server running on port %s in %s mode\n", appCfg.App.Port, appCfg.App.Env)
	var dbConn *sql.DB = db

	userRepo := gateway.NewUserRepositry(dbConn)
	tokenRepo := gateway.NewTokenRepository(dbConn)
	courseRepo := gateway.NewCourserRepositry(dbConn)
	enrollmentRepo := gateway.NewEnrollmentRepository(dbConn)
	moduleRepo := gateway.NewModuleRepository(dbConn)
	lessonRepo := gateway.NewLessonRepository(dbConn)
	orgRepo := gateway.NewOrganizationRepository(dbConn)
	paymentRepo := gateway.NewPaymentRepository(dbConn)
	adminRepo := gateway.NewAdminRepository(dbConn)



	// Initialize Services
	userService := service.NewUserService(userRepo, tokenRepo)
	courseService := service.NewCourseService(courseRepo, tokenRepo)
	enrollmnetsService := service.NewEnrollmentService(enrollmentRepo, tokenRepo)
	moduleService := service.NewModuleService(moduleRepo, tokenRepo)
	lessonService := service.NewlessonService(lessonRepo, tokenRepo)
	orgService := service.NewOrganizationService(orgRepo, tokenRepo)
	paymentService := service.NewpaymentService(paymentRepo, tokenRepo)
	adminService :=    service.NewAdminService(adminRepo, tokenRepo)


	// Initialize Controllers
	userController := controller.NewUserController(userService)
	courseController := controller.NewCourseController(courseService)
	enrollmentControlller := controller.NewEnrollmentController(enrollmnetsService)
	moduleController := controller.NewModuleController(moduleService)
	lessonController := controller.NewLessonController(lessonService)
	orgController := controller.NeworgsController(orgService)
	paymentController := controller.NewPaymentController(paymentService)
	adminController := controller.NewAdminController(adminService)


	// Setup Gin HTTP Server
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register API Routes
	routes.RegisterUserRoutes(r, userController, tokenRepo)
	routes.RegisterCourseRoutes(r, courseController, tokenRepo)
	routes.RegisterEnrollmentRoutes(r, *enrollmentControlller, tokenRepo)
	routes.RegisterModuleRoutes(r, *moduleController, tokenRepo )
	routes.RegisterLessonRoutes(r, *lessonController, tokenRepo)
	routes.RegisterOrganizationRoutes(r, *orgController, tokenRepo)
	routes.RegisterpaymentRoutes(r, *paymentController, tokenRepo)
	routes.RegisterAdminRoutes(r, adminController, tokenRepo)


	// Start Gin server (blocks here, keeps container alive)
	if err := r.Run(fmt.Sprintf(":%s", appCfg.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
