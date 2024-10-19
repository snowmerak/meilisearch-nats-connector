package meilisearchnatsconnector

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type NatsConnectionConfig struct {
	host     string
	username string
	password string
	token    string
}

func NewNatsConnectionConfig() *NatsConnectionConfig {
	return &NatsConnectionConfig{
		host:     "localhost:4222",
		username: "",
		password: "",
		token:    "",
	}
}

func (cfg *NatsConnectionConfig) SetHost(host string) *NatsConnectionConfig {
	cfg.host = host
	return cfg
}

func (cfg *NatsConnectionConfig) SetUsername(username string) *NatsConnectionConfig {
	cfg.username = username
	return cfg
}

func (cfg *NatsConnectionConfig) SetPassword(password string) *NatsConnectionConfig {
	cfg.password = password
	return cfg
}

func (cfg *NatsConnectionConfig) SetToken(token string) *NatsConnectionConfig {
	cfg.token = token
	return cfg
}

type NatsConnection struct {
	conn *nats.Conn
	uuid string
}

func NewNatsConnection(ctx context.Context, config *NatsConnectionConfig) (*NatsConnection, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("nats: uuid: %w", err)
	}

	conn, err := nats.Connect(config.host, nats.UserInfo(config.username, config.password), nats.Token(config.token))
	if err != nil {
		return nil, fmt.Errorf("nats: connect: %w", err)
	}

	context.AfterFunc(ctx, func() {
		conn.Close()
	})

	return &NatsConnection{
		conn: conn,
		uuid: id.String(),
	}, nil
}
