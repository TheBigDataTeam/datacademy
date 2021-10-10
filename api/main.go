package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/authors"
	"github.com/Serj1c/datalearn/api/pkg/courses"
	"github.com/Serj1c/datalearn/api/pkg/handlers"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/session"
	"github.com/Serj1c/datalearn/api/pkg/users"
	"github.com/Serj1c/datalearn/api/pkg/util"
	mgo "github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	/* TODO move this file to cmd folder and update Dockerfile accordingly */

	config, err := util.LoadConfig("./") // TODO address of a config file
	if err != nil {
		log.Fatal("unable to read configuration: ", err)
	}

	// init PostgreSQL
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to db, err: %v\n", err)
	}

	// init MongoDB
	mongo, err := mgo.Dial("mongodb://localhost") /* TODO */
	collection := mongo.DB("datalearn").C("authors")

	// init logger
	l := log.New(os.Stdout, "API ", log.LstdFlags)

	//init validation
	v := middleware.NewValidation() /* TODO: currently not registered */

	// init repos
	cr := courses.NewRepo(db)
	ar := authors.NewRepo(mongo, collection)
	ur := users.NewRepo(db)
	s := session.NewDBSession(db)

	// create the handlers
	coursesHandler := handlers.NewCourses(l, v, cr)
	authorsHandler := handlers.NewAuthors(l, v, ar)
	usersHandler := handlers.NewUsers(l, v, ur, s)

	sm := mux.NewRouter()

	// handlers for API calls
	sm.HandleFunc("/courses", coursesHandler.ListAll).Methods("GET")
	sm.HandleFunc("/courses", coursesHandler.Create).Methods("POST")
	sm.HandleFunc("/authors", authorsHandler.ListAll).Methods("GET")
	sm.HandleFunc("/courses/{id}", coursesHandler.ListOne).Methods("GET")
	sm.HandleFunc("/authors/{id}", authorsHandler.ListOne).Methods("GET")
	sm.HandleFunc("/api/users", usersHandler.Get).Methods("GET") /* currently is not used */

	sm.HandleFunc("/courses/{id}", coursesHandler.Update).Methods("PUT")
	sm.HandleFunc("/courses/{id}", authorsHandler.Update).Methods("PUT")

	sm.HandleFunc("/api/auth/signup", usersHandler.Signup).Methods("POST")
	sm.HandleFunc("/api/auth/login", usersHandler.Login).Methods("POST")
	sm.HandleFunc("/api/auth/user", usersHandler.GetBySessionID).Methods("GET")
	sm.HandleFunc("/api/auth/logout", usersHandler.Logout).Methods("GET")

	//sm.Use(coursesHandler.MiddlewareValidateCourse)

	sm.HandleFunc("/courses/{id}", coursesHandler.Delete).Methods("DELETE")
	sm.HandleFunc("/authors/{id}", authorsHandler.Delete).Methods("DELETE")

	// define middleware to handle CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
		MaxAge:           10000, /* TODO check that it is necessary */
		ExposedHeaders:   []string{"set-cookie"},
	})

	// register middleware
	withAuthHandler := session.AuthMiddleware(s, sm)
	withCorsHandler := c.Handler(withAuthHandler)

	server := &http.Server{
		Addr:         config.ServerPort,
		Handler:      withCorsHandler,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
