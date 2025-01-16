package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/IlliaSh1/backend/configs"
	"github.com/IlliaSh1/backend/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	config          *configs.Config
	Storage         *storage.Storage
	serviceProvider *serviceProvider
	fiberApp        *fiber.App
}

func NewApp() (*App, error) {
	app := &App{}

	app.fiberApp = fiber.New()

	err := app.initDependencies()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	err := app.fiberApp.Listen(fmt.Sprintf(":%d", app.config.ServerConfig.Port))
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initDependencies() error {
	inits := []func() error{
		app.initConfigs,
		app.initStorage,
		app.initServiceProvider,
		app.initControllers,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initConfigs() error {

	args_arr := os.Args

	args_map := make(map[string]string)
	for _, arg := range args_arr {
		key_val := strings.Split(arg, "=")
		if len(key_val) < 2 {
			continue
		}

		arg_key, arg_val := key_val[0], key_val[1]
		args_map[arg_key] = arg_val
	}

	var path_to_config string
	var ok bool
	if path_to_config, ok = args_map["--config_path"]; !ok {
		return fmt.Errorf("--config-path argument is empty or not provided")
	}

	cfg := configs.Config{}
	err := cleanenv.ReadConfig(path_to_config, &cfg)
	if err != nil {
		return fmt.Errorf("falied to load config: %v", err)
	}

	app.config = &cfg

	log.Info("server config: ", app.config)

	return nil
}

func (app *App) initStorage() error {
	storage, err := storage.NewStorage(&app.config.DBConfig)
	if err != nil {
		return err
	}
	app.Storage = storage
	return nil
}

func (app *App) initServiceProvider() error {
	app.serviceProvider = newServiceProvider(
		app.Storage.DB,
		&app.config.JWTConfig,
	)

	return nil
}
