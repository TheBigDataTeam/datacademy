package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/config"
	"github.com/Serj1c/datalearn/api/pkg/handlers/author"
	"github.com/Serj1c/datalearn/api/pkg/handlers/course"
	"github.com/Serj1c/datalearn/api/pkg/handlers/user"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/repository"
	"github.com/Serj1c/datalearn/api/pkg/service"
	"github.com/Serj1c/datalearn/api/pkg/session"
	mgo "github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {

	// init logger
	l := log.New(os.Stdout, "API ", log.LstdFlags)

	// read config file
	cfg, err := config.Load("configs")
	if err != nil {
		log.Fatal("unable to read configuration: ", err)
	}

	// init PostgreSQL
	db, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		l.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		l.Fatalf("Cannot connect to postgres, err: %v\n", err)
	}

	// init MongoDB
	mongo, err := mgo.Dial(cfg.MongoSource)
	if err != nil {
		log.Fatalf("Cannot connect to mongo, err: %v\n", err)
	}
	collection := mongo.DB(cfg.MongoDBname).C(cfg.MongoCname)

	//init validation
	v := middleware.NewValidation() /* TODO: currently not registered */

	// init repos
	cr := repository.NewCourseRepository(db)
	ar := repository.NewAuthorRepository(mongo, collection)
	ur := repository.NewUserRepository(db)
	s := session.New(db)

	// init services
	ap := service.NewAuthorProcessor(ar)

	// init handlers
	courseHandler := course.NewCourseHandler(l, v, cr)
	authorHandler := author.NewAuthorHandler(l, v, ap)
	usersHandler := user.NewUserHandler(l, v, ur, s)

	sm := mux.NewRouter()

	// register handler-functions
	sm.HandleFunc("/api/authors", authorHandler.List).Methods("GET")
	sm.HandleFunc("/api/authors/{id}", authorHandler.Get).Methods("GET")
	sm.HandleFunc("/api/authors/name/{name}", authorHandler.GetByName).Methods("GET")

	sm.HandleFunc("/api/courses", courseHandler.List).Methods("GET")
	sm.HandleFunc("/api/courses/{id}", courseHandler.Get).Methods("GET")

	sm.HandleFunc("/api/users", usersHandler.Get).Methods("GET") /* currently is not used */

	sm.HandleFunc("/api/auth/signup", usersHandler.Signup).Methods("POST")
	sm.HandleFunc("/api/auth/login", usersHandler.Login).Methods("POST")
	sm.HandleFunc("/api/auth/logout", usersHandler.Logout).Methods("GET")
	sm.HandleFunc("/api/auth/user", usersHandler.GetBySessionID).Methods("GET")

	/* Administration endpoints */
	sm.HandleFunc("/api/admin/add/author", authorHandler.Create).Methods("POST")
	sm.HandleFunc("/api/admin/add/course", courseHandler.Create).Methods("POST")
	sm.HandleFunc("/courses/{id}", courseHandler.Delete).Methods("DELETE")
	sm.HandleFunc("/courses/{id}", courseHandler.Update).Methods("PUT")
	sm.HandleFunc("/authors/{id}", authorHandler.Update).Methods("PUT")
	sm.HandleFunc("/authors/{id}", authorHandler.Delete).Methods("DELETE")

	// define middleware to handle CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.AppURL},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
		MaxAge:           10000, /* TODO check that it is necessary */
		ExposedHeaders:   []string{"set-cookie"},
	})

	// register middlewares
	withAuthHandler := session.AuthMiddleware(s, sm)
	withCorsHandler := c.Handler(withAuthHandler)

	// init server
	server := &http.Server{
		Addr:         cfg.ServerPort,
		Handler:      withCorsHandler,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start server
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
