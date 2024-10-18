package msutil

import (
	"encoding/json"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

type SearchRequest meilisearch.SearchRequest

func NewSearchRequest() *SearchRequest {
    return &SearchRequest{}
}

func (s *SearchRequest) Serialize() ([]byte, error) {
    data, err := json.Marshal(s)
    if err != nil {
        return nil, fmt.Errorf("searchRequest: serialize: %w", err)
    }

    return data, nil
}

func (s *SearchRequest) Deserialize(data []byte) error {
    err := json.Unmarshal(data, s)
    if err != nil {
        return fmt.Errorf("searchRequest: deserialize: %w", err)
    }

    return nil
}

type SearchResponse meilisearch.SearchResponse

func NewSearchResponse() *SearchResponse {
    return &SearchResponse{}
}

func (s *SearchResponse) Serialize() ([]byte, error) {
    data, err := json.Marshal(s)
    if err != nil {
        return nil, fmt.Errorf("searchResponse: serialize: %w", err)
    }

    return data, nil
}

func (s *SearchResponse) Deserialize(data []byte) error {
    err := json.Unmarshal(data, s)
    if err != nil {
        return fmt.Errorf("searchResponse: deserialize: %w", err)
    }

    return nil
}

type SearchGenericResponse[T any] struct {
	Hits               []*T   `json:"hits"`
	EstimatedTotalHits int64         `json:"estimatedTotalHits,omitempty"`
    Offset             int64         `json:"offset,omitempty"`
    Limit              int64         `json:"limit,omitempty"`
    ProcessingTimeMs   int64         `json:"processingTimeMs"`
    Query              string        `json:"query"`
    FacetDistribution  interface{}   `json:"facetDistribution,omitempty"`
    TotalHits          int64         `json:"totalHits,omitempty"`
    HitsPerPage        int64         `json:"hitsPerPage,omitempty"`
    Page               int64         `json:"page,omitempty"`
    TotalPages         int64         `json:"totalPages,omitempty"`
    FacetStats         interface{}   `json:"facetStats,omitempty"`
    IndexUID           string        `json:"indexUid,omitempty"`
}

func NewSearchGenericResponse[T any]() *SearchGenericResponse[T] {
    return &SearchGenericResponse[T]{}
}

func (s *SearchGenericResponse[T]) Serialize() ([]byte, error) {
    data, err := json.Marshal(s)
    if err != nil {
        return nil, fmt.Errorf("searchGenericResponse: serialize: %w", err)
    }

    return data, nil
}

func (s *SearchGenericResponse[T]) Deserialize(data []byte) error {
    err := json.Unmarshal(data, s)
    if err != nil {
        return fmt.Errorf("searchGenericResponse: deserialize: %w", err)
    }

    return nil
}

type IndexConfig meilisearch.IndexConfig

func NewIndexConfig(indexName string, primaryKey string) *IndexConfig {
    return &IndexConfig{
        Uid: indexName,
        PrimaryKey: primaryKey,
    }
}

func (i *IndexConfig) Serialize() ([]byte, error) {
    data, err := json.Marshal(i)
    if err != nil {
        return nil, fmt.Errorf("indexConfig: serialize: %w", err)
    }

    return data, nil
}

func (i *IndexConfig) Deserialize(data []byte) error {
    err := json.Unmarshal(data, i)
    if err != nil {
        return fmt.Errorf("indexConfig: deserialize: %w", err)
    }

    return nil
}
