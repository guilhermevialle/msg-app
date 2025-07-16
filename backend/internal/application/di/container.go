package di

import (
	app_services "app/internal/application/services"
	"app/internal/application/use_cases"
	env "app/internal/infra"
	"app/internal/infra/http/controllers"
	"app/internal/infra/http/middlewares"
	"app/internal/infra/repositories"
	infra_services "app/internal/infra/services"
)

type Container struct {
	AuthService     app_services.IAuthService
	AuthController  controllers.IAuthController
	AuthMiddleware  middlewares.IAuthMiddleware
	UserController  controllers.IUserController
	PostController  controllers.IPostController
	DebugController controllers.IDebugController
}

func NewContainer() *Container {
	ur := repositories.NewUserRepository()
	ts := infra_services.NewJwtTokenService(env.TOKEN_SECRET)
	am := middlewares.NewAuthMiddleware(ts)
	hs := infra_services.NewBcryptHashService()
	as := app_services.NewAuthService(ur, ts, hs)
	ac := controllers.NewAuthController(as)
	pr := repositories.NewPostRepository()
	cp := use_cases.NewCreatePost(ur, pr)
	lp := use_cases.NewLikePost(pr, ur)
	actp := use_cases.NewAddCommentToPost(ur, pr)
	gap := use_cases.NewGetAllPosts(pr)
	pc := controllers.NewPostController(cp, lp, actp, gap)
	dc := controllers.NewDebugController(ur, pr)

	// use cases
	gupUc := use_cases.NewGetUserProfile(ur)
	uc := controllers.NewUserController(gupUc)

	return &Container{
		AuthService:     as,
		AuthController:  ac,
		AuthMiddleware:  am,
		UserController:  uc,
		PostController:  pc,
		DebugController: dc,
	}
}
