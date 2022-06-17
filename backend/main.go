package main

import (
	"context"
	"exam/users"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	//headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	//origins := handlers.AllowedOrigins([]string{"*"})

	userStorage := users.NewInMemoryUserStorage()
	userService := users.NewUserService(userStorage)
	jwtService, err := users.NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/cake", logRequest(getCakeHandler)).Methods(http.MethodGet)

	r.HandleFunc("/user/signup", logRequest(userService.Register)).Methods(http.MethodPost)
	r.HandleFunc("/user/signin", logRequest(users.WrapJwt(jwtService, userService.JWT))).Methods(http.MethodPost)

	r.HandleFunc("/todo/lists", logRequest(jwtService.JWTAuth(userStorage, users.AddList))).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists", logRequest(jwtService.JWTAuth(userStorage, users.GetLists))).Methods(http.MethodGet)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}", logRequest(jwtService.JWTAuth(userStorage, users.UpdateList))).Methods(http.MethodPut)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}", logRequest(jwtService.JWTAuth(userStorage, users.DeleteList))).Methods(http.MethodDelete)

	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", logRequest(jwtService.JWTAuth(userStorage, users.AddTask))).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", logRequest(jwtService.JWTAuth(userStorage, users.GetTasks))).Methods(http.MethodGet)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", logRequest(jwtService.JWTAuth(userStorage, users.UpdateTask))).Methods(http.MethodPut)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", logRequest(jwtService.JWTAuth(userStorage, users.DeleteTask))).Methods(http.MethodDelete)

	handler := cors.Default().Handler(r)

	srv := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Println("Server started")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Error: ", err)
	}
}

func getCakeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cake"))
}
