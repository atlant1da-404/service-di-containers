package main

import (
	"atlant1da-404/service-di-containers/auth/app/internal/application"
	"atlant1da-404/service-di-containers/auth/app/internal/infrastructure/lite"
	"atlant1da-404/service-di-containers/auth/app/internal/interface/ahttp"
	"atlant1da-404/service-di-containers/auth/app/internal/pkg/auth"
	"atlant1da-404/service-di-containers/pkg/public/config"
	"atlant1da-404/service-di-containers/pkg/public/database"
	"atlant1da-404/service-di-containers/pkg/public/ghttp"
	"atlant1da-404/service-di-containers/pkg/public/hash"
	"fmt"
	"github.com/sarulabs/di"
	"log"
)

func main() {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(di.Def{
		Name: "config",
		Build: func(ctn di.Container) (interface{}, error) {
			return config.GetConfig(), nil
		},
		Scope: di.App,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(di.Def{
		Name: "database",
		Build: func(ctn di.Container) (interface{}, error) {
			cfg, ok := ctn.Get("config").(*config.Config)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of config: %w", err)
			}
			db, err := database.NewSQLDatabase(cfg.Address)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to database: %w", err)
			}

			return db, nil
		},
		Scope: di.App,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(di.Def{
		Name: "storage",
		Build: func(ctn di.Container) (interface{}, error) {
			db, ok := ctn.Get("database").(database.Database)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of database: %w", err)
			}
			return lite.NewAuthStorage(db), nil

		},
		Scope: di.App,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(di.Def{
		Name: "service",
		Build: func(ctn di.Container) (interface{}, error) {
			storage, ok := ctn.Get("storage").(application.AuthStorage)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of storage: %w", err)
			}
			cfg, ok := ctn.Get("config").(*config.Config)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of config: %w", err)
			}

			return application.NewAuthService(storage, hash.NewHash(), auth.NewAuth(cfg.SecretKey)), nil

		},
		Scope: di.App,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(di.Def{
		Name: "controller",
		Build: func(ctn di.Container) (interface{}, error) {
			service, ok := ctn.Get("service").(application.AuthService)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of service: %w", err)
			}
			cfg, ok := ctn.Get("config").(*config.Config)
			if !ok {
				return nil, fmt.Errorf("failed to conversion type of config: %w", err)
			}

			mux := ghttp.NewHTTPRouter(cfg.Address)
			router := ahttp.NewController(service)

			mux.Post("/api/auth/login", mux.ErrorWrapper(router.Login))
			mux.Post("/api/auth/refresh", mux.ErrorWrapper(router.Refresh))

			return mux, nil
		},
		Scope: di.App,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Fatal(builder.Build().Get("controller").(ghttp.HTTTPRouter).ListenAndServe())

}
