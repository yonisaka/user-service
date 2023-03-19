package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/yonisaka/user-service/cmd"
	"github.com/yonisaka/user-service/config"
	"github.com/yonisaka/user-service/domain/service"
	"github.com/yonisaka/user-service/grpc/client"
	"github.com/yonisaka/user-service/infrastructure/persistence"
	"github.com/yonisaka/user-service/rest"
	"github.com/yonisaka/user-service/rest/route"
)

// main is a main function
func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.New()

	db, errConn := persistence.NewDBConnection(conf.DBConfig)
	if errConn != nil {
		log.Fatalf("unable connect to database, %v", errConn)
	}

	repo := service.NewDBService(db)

	command := cmd.NewCommand(
		cmd.WithConfig(conf),
		cmd.WithRepo(repo),
	)

	app := cmd.NewCLI()
	app.Commands = command.Build()

	clientConn, errClient := client.NewGRPCConn(conf)
	if errClient != nil {
		log.Fatalf("grpc client unable connect to server, %v", errClient)
	}

	grpcClient := client.NewGRPCClient(clientConn)
	app.Action = func(ctx *cli.Context) error {
		router := route.NewRouter(
			route.WithConfig(conf),
			route.WithRepository(repo),
			route.WithGRPCClient(grpcClient),
		).Init()

		shutdownTimeout := 10 * time.Second

		err := rest.RunHTTPServer(router, strconv.Itoa(conf.AppPort), shutdownTimeout)
		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
