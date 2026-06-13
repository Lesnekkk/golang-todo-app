package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Lesnekkk/golang-todo-app/internal/core/config"
	core_logger "github.com/Lesnekkk/golang-todo-app/internal/core/logger"
	core_postgres "github.com/Lesnekkk/golang-todo-app/internal/core/repository/postgres"
	core_redis "github.com/Lesnekkk/golang-todo-app/internal/core/repository/redis"
	core_http_middleware "github.com/Lesnekkk/golang-todo-app/internal/core/transport/http/middlewear"

	statistics_postgres "github.com/Lesnekkk/golang-todo-app/internal/features/statistics/repository/postgres"
	statistics_service "github.com/Lesnekkk/golang-todo-app/internal/features/statistics/service"
	statistics_http "github.com/Lesnekkk/golang-todo-app/internal/features/statistics/transport/http"

	users_postgres "github.com/Lesnekkk/golang-todo-app/internal/features/users/repository/postgres"
	users_redis "github.com/Lesnekkk/golang-todo-app/internal/features/users/repository/redis"
	users_service "github.com/Lesnekkk/golang-todo-app/internal/features/users/service"
	users_http "github.com/Lesnekkk/golang-todo-app/internal/features/users/transport/http"

	tasks_postgres "github.com/Lesnekkk/golang-todo-app/internal/features/tasks/repository/postgres"
	tasks_redis "github.com/Lesnekkk/golang-todo-app/internal/features/tasks/repository/redis"
	tasks_service "github.com/Lesnekkk/golang-todo-app/internal/features/tasks/service"
	tasks_http "github.com/Lesnekkk/golang-todo-app/internal/features/tasks/transport/http"

	web_fs_repository "github.com/Lesnekkk/golang-todo-app/internal/features/web/repository/file_system"
	web_service "github.com/Lesnekkk/golang-todo-app/internal/features/web/service"
	web_http "github.com/Lesnekkk/golang-todo-app/internal/features/web/transport/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	logger, err := core_logger.NewLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("logger: %v", err)
	}
	defer logger.Close()

	pool, err := core_postgres.NewPool(context.Background(), cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("postgres: %v", err)
	}
	defer pool.Close()

	redisClient, err := core_redis.NewClient(context.Background(), cfg.RedisAddr())
	if err != nil {
		log.Fatalf("redis: %v", err)
	}
	defer redisClient.Close()

	// repositories
	usersRepo := users_postgres.NewUserPostgresRepository(pool)
	usersCache := users_redis.NewUsersRedisRepository(redisClient)
	tasksRepo := tasks_postgres.NewTasksPostgresRepository(pool)
	tasksCache := tasks_redis.NewTasksRedisRepository(redisClient)
	statsRepo := statistics_postgres.NewStatisticsPostgresRepository(pool)
	webRepo := web_fs_repository.NewWebRepository()

	// services
	usersService := users_service.NewUsersService(usersRepo, usersCache)
	tasksService := tasks_service.NewTaskService(tasksRepo, tasksCache)
	statsService := statistics_service.NewStatisticsService(statsRepo)
	webSvc := web_service.NewWebService(webRepo)

	// handlers
	usersHandler := users_http.NewUsersHTTPHandler(usersService)
	tasksHandler := tasks_http.NewTasksHTTPHandler(tasksService)
	statsHandler := statistics_http.NewStatisticsHTTPHandler(statsService)
	webHandler := web_http.NewWebHTTPHandler(webSvc)

	// router
	router := mux.NewRouter()
	router.Use(mux.MiddlewareFunc(core_http_middleware.RequestId()))
	router.Use(mux.MiddlewareFunc(core_http_middleware.Logger(logger.Logger)))

	// web
	router.HandleFunc("/", webHandler.GetMainPage).Methods("GET")

	// api/v1
	api := router.PathPrefix("/api/v1").Subrouter()

	// users
	api.HandleFunc("/users", usersHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users", usersHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", usersHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", usersHandler.PatchUser).Methods("PATCH")
	api.HandleFunc("/users/{id}", usersHandler.DeleteUser).Methods("DELETE")

	// tasks
	api.HandleFunc("/tasks", tasksHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks", tasksHandler.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", tasksHandler.GetTask).Methods("GET")
	api.HandleFunc("/tasks/{id}", tasksHandler.PatchTask).Methods("PATCH")
	api.HandleFunc("/tasks/{id}", tasksHandler.DeleteTask).Methods("DELETE")

	// statistics
	api.HandleFunc("/statistics", statsHandler.GetStatistics).Methods("GET")

	fmt.Printf("server starting on %s\n", cfg.HTTPAddr)
	if err := http.ListenAndServe(cfg.HTTPAddr, router); err != nil {
		log.Fatalf("server: %v", err)
	}
}
