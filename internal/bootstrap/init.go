package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	http2 "github.com/Ilja-R/library-auth-service/internal/adapter/driving/http"
	"github.com/Ilja-R/library-auth-service/internal/config"
	"github.com/Ilja-R/library-auth-service/internal/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initDB(cfg config.Postgres) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(cfg.ConnectionURL())
	if err != nil {
		return nil, err
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return db, err
	}

	// Connection configuration
	// more info here https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)

	return db, nil
}

func initHTTPService(
	cfg *config.Config,
	uc *usecase.UseCases,
) *http.Server {
	return http2.New(
		cfg,
		uc,
	)
}

func initRabbitMQ(cfg *config.RabbitMQ) (*amqp.Connection, *amqp.Channel, error) {
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	// Connect
	conn, err := amqp.Dial(connURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to open channel: %w", err)
	}

	return conn, ch, nil
}

func initAuthQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		"auth-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare auth queue: %w", err)
	}

	return &queue, nil
}
