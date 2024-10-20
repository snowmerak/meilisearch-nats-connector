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

func CreateIndex(c *Client, request *msutil.IndexConfig) error {
	_, err := c.client.CreateIndex((*meilisearch.IndexConfig)(request))
	if err != nil {
		return fmt.Errorf("meilisearch: createIndex: %w", err)
	}

	return nil
}

func DeleteIndex(c *Client, index string) error {
	_, err := c.client.DeleteIndex(index)
	if err != nil {
		return fmt.Errorf("meilisearch: deleteIndex: %w", err)
	}

	return nil
}

func TruncateIndex(c *Client, index string) error {
	_, err := c.client.Index(index).DeleteAllDocuments()
	if err != nil {
		return fmt.Errorf("meilisearch: truncateIndex: %w", err)
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

func DeleteDocuments(c *Client, index string, identifiers *msutil.Identifiers) error {
	_, err := c.client.Index(index).DeleteDocuments(identifiers.GetList())
	if err != nil {
		return fmt.Errorf("meilisearch: deleteDocument: %w", err)
	}

	return nil
}

func UploadSynonyms(c *Client, index string, synonyms *msutil.Synonyms) error {
	_, err := c.client.Index(index).UpdateSynonyms(synonyms.ToMap())
	if err != nil {
		return fmt.Errorf("meilisearch: uploadSynonyms: %w", err)
	}

	return nil
}

func ResetSynonyms(c *Client, index string) error {
	_, err := c.client.Index(index).ResetSynonyms()
	if err != nil {
		return fmt.Errorf("meilisearch: resetSynonyms: %w", err)
	}

	return nil
}

func GetSynonyms(c *Client, index string) (*msutil.Synonyms, error) {
	synonyms, err := c.client.Index(index).GetSynonyms()
	if err != nil {
		return nil, fmt.Errorf("meilisearch: getSynonyms: %w", err)
	}

	value := msutil.NewSynonyms()
	for k, ws := range *synonyms {
		value.SetList(k, ws)
	}

	return value, nil
}
