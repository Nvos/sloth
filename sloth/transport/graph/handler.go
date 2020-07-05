package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"net/http"
	"sloth/infra"
	"sloth/transport/graph/generated"
	"sloth/transport/graph/resolver"
)

func NewHandler(deps resolver.Dependencies) http.Handler {
	root := resolver.New(deps)

	cfg := generated.Config{
		Resolvers: root,
	}

	exec := generated.NewExecutableSchema(cfg)
	srv := handler.NewDefaultServer(exec)
	//srv.SetErrorPresenter(errorPresenter(deps.Logger))

	return srv
}

// Useless
func errorPresenter(log infra.Logger) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		if gqlerr, ok := err.(*gqlerror.Error); ok {
			return gqlerr
		}

		log.Background().Error("graphql", zap.Error(err))
		hidden := graphql.DefaultErrorPresenter(ctx, err)
		hidden.Message = "error redacted"

		return hidden
	}
}