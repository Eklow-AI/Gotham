package main

import (
	"fmt"
	stdlog "log"

	"github.com/Eklow-AI/Gotham/src/sdk"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		stdlog.Fatal("Error loading .env file")
	}

	//Connect Postgres database
	//models.ConnectDB()
	//Setup SDK
	sdk.SetupSDK()
	contracts := sdk.GetContractsFromCage("6ZP36")
	fmt.Println(len(contracts))

	// /*
	//  * Create and configure router and all its routes
	//  */
	// router := mux.NewRouter()
	// // User management
	// router.HandleFunc("/login", handlers.SignInUser).Methods("POST")
	// router.HandleFunc("/register", handlers.CreateUser).Methods("POST")

	// /*
	//  *  Set up and configure logger
	//  */
	// var logger log.Logger
	// logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// // Direct any attempts to use Go's log package to our structured logger
	// stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// // Log the timestamp (in UTC) and the loc (file + line number) of the logging
	// logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	// loggingMiddleware := middleware.LoggingMiddleware(logger)
	// loggedRouter := loggingMiddleware(router)

	// // Start application
	// if err := http.ListenAndServe(":"+os.Getenv("PORT"), loggedRouter); err != nil {
	// 	logger.Log("status", "fatal", "err", err)
	// 	os.Exit(1)
	// }
}
