package msutil

import "github.com/meilisearch/meilisearch-go"

type SearchRequest meilisearch.SearchRequest

type SearchResponse meilisearch.SearchResponse

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

type IndexConfig meilisearch.IndexConfig

func NewIndexConfig(indexName string, primaryKey string) *IndexConfig {
    return &IndexConfig{
        Uid: indexName,
        PrimaryKey: primaryKey,
    }
}
