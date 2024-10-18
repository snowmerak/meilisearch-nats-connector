package meilisearchnatsconnector

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

type MeilisearchConfig struct {
	host string
	apiKey string
	timeout time.Duration
}

func NewMeiliSearchConfig() *MeilisearchConfig {
	return &MeilisearchConfig{
		host: "http://localhost:7700",
		apiKey: "",
		timeout: 5 * time.Second,
	}
}

func (cfg *MeilisearchConfig) SetHost(host string) *MeilisearchConfig {
	cfg.host = host
	return cfg
} 

func (cfg *MeilisearchConfig) SetAPIKey(apiKey string) *MeilisearchConfig {
	cfg.apiKey = apiKey
	return cfg
}

func (cfg *MeilisearchConfig) SetTimeout(timeout time.Duration) *MeilisearchConfig {
	cfg.timeout = timeout
	return cfg
}

type Client struct {
	client meilisearch.ServiceManager
}

func New(config *MeilisearchConfig) *Client {
	hc := &http.Client{Timeout: config.timeout}
	tc := &tls.Config{}

	client := meilisearch.New(config.host, meilisearch.WithAPIKey(config.apiKey), meilisearch.WithCustomClient(hc), meilisearch.WithCustomClientWithTLS(tc))

	return &Client{client: client}
}


