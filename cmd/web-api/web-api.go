package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"xsolla-sdk-backend/internal/handlers"
	"xsolla-sdk-backend/internal/server/restapi/operations/healthcheck"
	"xsolla-sdk-backend/internal/server/restapi/operations/login"

	"xsolla-sdk-backend/internal/app"
	"xsolla-sdk-backend/internal/config"
	"xsolla-sdk-backend/internal/server/restapi"
	"xsolla-sdk-backend/internal/server/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var (
	logger      *log.Entry
	ProgramName = "Xsolla SDK Backend"
)

const (
	ProgramDescription   = "Example backend for Xsolla SDK"
	AccessControlMaxPage = 60 * 60 * 24 * 7 // 7 days
)

func main() {
	envFlag := flag.Bool("envs", false, "Show configuration option")
	flag.Parse()
	if *envFlag {
		description, err := config.GetConfigDescription()
		if err != nil {
			fmt.Println("Failed out config description. Error: ", err)
			os.Exit(1)
		}
		fmt.Println(description)
		os.Exit(0)
	}

	logger = createLogger(ProgramName)

	c, err := config.Init()
	if err != nil {
		logger.Errorf("Failed init configuration. Error: %s", err)
		os.Exit(1)
	}

	logger = setupLoggerOutput(logger, c.LogPath, ProgramName)

	application := app.NewApplication(
		nil,
		nil,
		&c,
		logger)

	// init handlers
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewXsollaSdkBackendAPI(swaggerSpec)
	api.Logger = logger.Printf

	setupAPI(api, &application)

	// small hack for using custom Handler with middlewares
	server := restapi.NewServer(nil)
	server.SetHandler(setupAPI(api, &application))

	parseArgs(server, api)

	defer server.Shutdown() // nolint

	server.ConfigureAPI()
	server.Port = c.ServerPort
	server.Host = c.ServerHost
	if err := server.Serve(); err != nil {
		logger.Errorf("Failed serve requests. Error: %s", err)
	}
}

func createLogger(programName string) *log.Entry {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyMsg:   "MESSAGE",
			log.FieldKeyLevel: "LEVEL",
		},
	})
	return l.WithField("PROGRAM", programName)
}
func setupLoggerOutput(logEntry *log.Entry, logPath, programName string) *log.Entry {
	le := logEntry.Logger
	if logPath == "" {
		le.SetOutput(os.Stdout)
	} else {
		output, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			logger.Fatalf("logger output: %v", err)
		}
		le.SetOutput(output)
	}
	return le.WithField("PROGRAM", programName)
}
func setupAPI(api *operations.XsollaSdkBackendAPI, application *app.Application) http.Handler {
	healthcheckHandler := handlers.NewHealthcheckHandler(application)

	api.HealthcheckHealthcheckHandler =
		healthcheck.HealthcheckHandlerFunc(healthcheckHandler.CheckHealth)

	loginHandler := handlers.NewLoginHandler(application)

	api.LoginLoginHandler =
		login.LoginHandlerFunc(loginHandler.Login)

	return middlewareCORS(api.Serve(setupMiddlewares))
}
func parseArgs(server *restapi.Server, api *operations.XsollaSdkBackendAPI) {
	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = ProgramName
	parser.LongDescription = ProgramDescription
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
}
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}
func middlewareCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-XSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", strconv.Itoa(AccessControlMaxPage)) // 7 day

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
