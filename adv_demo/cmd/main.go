package main

import (
	"fmt"
	"net/http"

	"adv_demo/configs"
	"adv_demo/internal/auth"
	"adv_demo/internal/link"
	"adv_demo/internal/stat"
	"adv_demo/internal/user"
	"adv_demo/pkg/db"
	"adv_demo/pkg/event"
	"adv_demo/pkg/jwt"
	"adv_demo/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// JWT
	jwt := jwt.NewJWT(conf.Auth.Secret)

	// Services
	authService := auth.NewAuthService(userRepository, jwt)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
