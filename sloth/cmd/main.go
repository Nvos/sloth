package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"net/http"
	"sloth/feature/activity"
	"sloth/infra"
	"sloth/store/db"
	"sloth/transport"
	"sloth/transport/graph"
	"sloth/transport/graph/resolver"
)

func main() {
	cfg := infra.MustNewConfiguration()
	logger := infra.MustNewLogger(cfg)

	defer func() {
		// TODO, 24.05.2020: verify how to handle sync error
		_ = logger.Background().Sync()
	}()

	client := infra.MustNewEntClient(cfg, logger)
	defer func() {
		err := client.Close()
		logger.Background().Warn("failed to gracefully close ent client", zap.Error(err))
	}()

	activityStore := db.NewActivityStore(client)
	activityService := activity.NewService(activityStore)

	gqlHandle := graph.NewHandler(resolver.Dependencies{
		ActivityService: activityService,
		Logger: logger,
	})

	
	router := transport.NewRouter()
	router.Handle("/query", gqlHandle)

	if cfg.Environment.Mode == infra.Develop {
		router.Handle("/graphql", playground.Handler("Sloth playground", "/query"))
	}

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Background().Panic("failed to start http server", zap.Error(err))
	}
}
