package infra

import (
	"context"
	"go.uber.org/zap"
	"sloth/ent"
	"sloth/ent/migrate"
)

func MustNewEntClient(cfg *Configuration, logger Logger) *ent.Client {
	client, err := ent.Open("sqlite3", "file:slodb?mode=memory&cache=shared&_fk=1")
	if err != nil {
		logger.Background().Panic("failed to open ent client", zap.Error(err))
	}

	if cfg.Environment.Mode == Develop {
		client = client.Debug()

		if err := client.Schema.Create(
			context.Background(),
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
			migrate.WithGlobalUniqueID(true),
		); err != nil {
			logger.Background().Panic("schema creation failed", zap.Error(err))
		}
	}

	return client
}
