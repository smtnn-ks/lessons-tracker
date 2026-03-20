package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sethvargo/go-envconfig"
	"github.com/smtnn-ks/lessons-tracker/internal/app/pages/course_list"
	course_list_view "github.com/smtnn-ks/lessons-tracker/internal/app/pages/course_list/view"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var config Config
	if err := envconfig.Process(ctx, &config); err != nil {
		slog.New(slog.NewJSONHandler(os.Stderr, nil)).Error("failed to parse config", "error", err)
		os.Exit(1)
	}

	mustInitLogger(config)
	defer zap.L().Sync()

	courseListHandler := course_list.NewCourseListHandler(course_list_view.CourseList)

	// TODO pass zap.Logger to logger middleware
	r := chi.NewRouter()

	r.Route("/courses", courseListHandler.Mount)

	zap.L().Info("server is started", zap.Int("port", config.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r); err != nil {
		zap.L().Fatal("fatal server error", zap.Error(err))
	}
}

func mustInitLogger(config Config) {
	var newLoggerFunc func(options ...zap.Option) (*zap.Logger, error)
	switch config.Env {
	case "local":
		newLoggerFunc = zap.NewDevelopment
	case "production":
		newLoggerFunc = zap.NewProduction
	default:
		slog.New(slog.NewJSONHandler(os.Stderr, nil)).Error("invalid env: " + config.Env)
		os.Exit(1)
	}
	logger, err := newLoggerFunc()
	if err != nil {
		zap.L().Fatal("failed to get dev logger", zap.Error(err))
	}

	zap.ReplaceGlobals(logger)
}
