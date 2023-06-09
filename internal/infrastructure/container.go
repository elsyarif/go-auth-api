package infrastructure

import (
	"github.com/elsyarif/go-auth-api/internal/applications/usecases"
	"github.com/elsyarif/go-auth-api/internal/domain/services"
	"github.com/elsyarif/go-auth-api/internal/infrastructure/persistence/postgresql/repositories"
	"github.com/elsyarif/go-auth-api/internal/interface/http/handler"
	"github.com/elsyarif/go-auth-api/pkg/encryption"
	"github.com/elsyarif/go-auth-api/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Container(db *sqlx.DB, app *gin.Engine) {
	idGenerator := uid.NewNanoId()
	hash := encryption.PasswordHash()
	// User
	userRepository := repositories.NewUserRepositoryPostgres(db)
	userService := services.NewUserService(userRepository, idGenerator, hash)
	userUseCase := usecases.NewUserUseCase(userService)
	userHandler := handler.NewUserHandler(userUseCase)

	// authentication
	authRepository := repositories.NewAuthRepositoryPostgres(db)
	authService := services.NewAuthService(authRepository, userRepository, hash)
	authUseCae := usecases.NewAuthUseCase(authService, userService)
	authHandler := handler.NewAuthHandler(authUseCae)

	// routes
	userHandler.Routes(app)
	authHandler.Routes(app)
}
