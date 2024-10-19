package msutil

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/snowmerak/meilisearch-nats-connector/gen/model"
)

type Synonyms struct {
	message *model.Synonyms
}

func NewSynonyms() *Synonyms {
	return &Synonyms{}
}

func (s *Synonyms) SetList(key string, synonyms []string) *Synonyms {
	s.message.Synonyms[key] = &model.Words{Words: synonyms}
	return s
}

func (s *Synonyms) Add(key string, synonym ...string) *Synonyms {
	if s.message.Synonyms == nil {
		s.message.Synonyms = make(map[string]*model.Words)
	}
	s.message.Synonyms[key].Words = append(s.message.Synonyms[key].Words, synonym...)
	return s
}

func (s *Synonyms) GetList() map[string]*model.Words {
	return s.message.Synonyms
}

func (s *Synonyms) ToMap() *map[string][]string {
	synonyms := make(map[string][]string)
	for k, ws := range s.message.Synonyms {
		synonyms[k] = ws.Words
	}
	return &synonyms
}

func (s *Synonyms) Drop(key string) *Synonyms {
	delete(s.message.Synonyms, key)
	return s
}

func (s *Synonyms) Serialize() ([]byte, error) {
	data, err := proto.Marshal(s.message)
	if err != nil {
		return nil, fmt.Errorf("synonyms: serialize: %w", err)
	}

	return data, nil
}

func (s *Synonyms) Deserialize(data []byte) error {
	err := proto.Unmarshal(data, s.message)
	if err != nil {
		return fmt.Errorf("synonyms: deserialize: %w", err)
	}

	return nil
}
