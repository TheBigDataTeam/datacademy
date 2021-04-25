package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/datalearn/api/data"
	"github.com/Serj1c/datalearn/api/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// create a logger for a server
	l := log.New(os.Stdout, "API ", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	coursesHandler := handlers.NewCourses(l, v)
	authorsHandler := handlers.NewAuthors(l, v)

	// register handlers
	sm := mux.NewRouter()

	// handlers for API
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/courses", coursesHandler.ListAllCourses)
	getRouter.HandleFunc("/authors", authorsHandler.ListAllAuthors)
	getRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.ListSingleCourse)
	getRouter.HandleFunc("/authors/{id:[0-9]+}", authorsHandler.ListSingleAuthor)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.UpdCourse)
	putRouter.Use(coursesHandler.MiddlewareValidateCourse)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/courses", coursesHandler.CreateCourse)
	postRouter.Use(coursesHandler.MiddlewareValidateCourse)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/courses/{id:[0-9]+}", coursesHandler.DelCourse)

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// create a new server
	server := &http.Server{
		Addr:         ":3100",
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
	timeoutContext, cancelContext := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelContext()
	server.Shutdown(timeoutContext)
}
