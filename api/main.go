package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/courses"
	"github.com/Serj1c/datalearn/api/pkg/data"
	"github.com/Serj1c/datalearn/api/pkg/handlers"
	"github.com/Serj1c/datalearn/api/pkg/util"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	/* TODO move this file to cmd folder and update Dockerfile accordingly */

	dsn := "root:spartak1@tcp(127.0.0.1:5432/datacademy?charset=utf8&interpolateParams=true)"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to db, err: %v\n", err)
	}

	// read configuration from api.env
	config, err := util.LoadConfig("./") // TODO address of a config file
	if err != nil {
		log.Fatal("unable to read configuration: ", err)
	}

	// init logger
	l := log.New(os.Stdout, "API ", log.LstdFlags)
	//init validation
	v := data.NewValidation()
	// init courses repo
	cr := courses.NewRepo(db)

	// create the handlers
	coursesHandler := handlers.NewCourses(l, v, cr)
	authorsHandler := handlers.NewAuthors(l, v)

	// register handlers
	sm := mux.NewRouter()

	// handlers for API
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/courses", coursesHandler.ListAll)
	getRouter.HandleFunc("/authors", authorsHandler.ListAllAuthors)
	getRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.ListOne)
	getRouter.HandleFunc("/authors/{id:[0-9]+}", authorsHandler.ListSingleAuthor)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.Update)
	putRouter.Use(coursesHandler.MiddlewareValidateCourse)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/courses", coursesHandler.Create)
	postRouter.Use(coursesHandler.MiddlewareValidateCourse)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.Delete)
	deleteRouter.HandleFunc("/authors/{id:[0-9]+}", authorsHandler.DelAuthor)

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// create a new server
	server := &http.Server{
		Addr:         config.ServerPort,
		Handler:      corsHandler(sm),   // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request for the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	// start the server
	go func() {
		l.Printf("Server is listening on port %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s", err)
			os.Exit(1)
		}
	}()

	// gracefully shutdown
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	// block until a signal is received
	sig := <-sigChannel
	l.Println("Command to terminate received, shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	timeoutContext, finish := context.WithTimeout(context.Background(), 30*time.Second)
	defer finish()
	server.Shutdown(timeoutContext)
}
