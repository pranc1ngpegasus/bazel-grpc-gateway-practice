package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var _ RDBConnector = (*rdbConnector)(nil)

type (
	RDBConnector interface {
		GetDB() *ent.Client
	}

	rdbConnector struct {
		config configuration.Config
		db     *ent.Client
	}
)

func NewRDBConnector(config configuration.Config) RDBConnector {
	dsn := newDSN(config)

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error().Err(err)
		panic(err)
	}

	drv := entsql.OpenDB(dialect.MySQL, dbConn)

	opt := []ent.Option{
		ent.Driver(drv),
	}
	client := ent.NewClient(opt...)

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Error().Err(err).Msg("failed creating schema resources")
	}

	return &rdbConnector{
		config: config,
		db:     client,
	}
}

func newDSN(config configuration.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
}

func (dbc *rdbConnector) GetDB() *ent.Client {
	return dbc.db
}
