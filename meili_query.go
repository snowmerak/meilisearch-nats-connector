package meilisearchnatsconnector

import (
	"encoding/json"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/snowmerak/meilisearch-nats-connector/msutil"
)

func Search(c *Client, index string, query string, searchRequest *msutil.SearchRequest) (*json.RawMessage, error) {
	res, err := c.client.Index(index).SearchRaw(query, (*meilisearch.SearchRequest)(searchRequest))
	if err != nil {
		return nil, fmt.Errorf("meilisearch: searchRaw: %w", err)
	}

	return res, nil
}

func Create(c *Client, request *msutil.IndexConfig) error {
	_, err := c.client.CreateIndex((*meilisearch.IndexConfig)(request))
	if err != nil {
		return fmt.Errorf("meilisearch: createIndex: %w", err)
	}

	return nil
}

func AddDocuments(c *Client, index string, document []byte) error {
	_, err := c.client.Index(index).AddDocumentsNdjson(document)
	if err != nil {
		return fmt.Errorf("meilisearch: addDocument: %w", err)
	}

	return nil
}

func UpdateDocuments(c *Client, index string, document []byte) error {
	_, err := c.client.Index(index).UpdateDocumentsNdjson(document)
	if err != nil {
		return fmt.Errorf("meilisearch: updateDocument: %w", err)
	}

	return nil
}

func DeleteDocuments(c *Client, index string, identifiers []string) error {
	_, err := c.client.Index(index).DeleteDocuments(identifiers)
	if err != nil {
		return fmt.Errorf("meilisearch: deleteDocument: %w", err)
	}

	return nil
}