package msutil

import "github.com/snowmerak/meilisearch-nats-connector/gen/model"

type Identifiers struct {
	message model.Identifiers
}

func NewIdentifiers() *Identifiers {
	return &Identifiers{}
}

func (i *Identifiers) SetList(identifiers []string) *Identifiers {
	i.message.Identifiers = identifiers
	return i
}

func (i *Identifiers) Add(identifier string) *Identifiers {
	i.message.Identifiers = append(i.message.Identifiers, identifier)
	return i
}

func (i *Identifiers) GetList() []string {
	return i.message.Identifiers
}

func (i *Identifiers) Remove(identifier string) *Identifiers {
	var identifiers []string
	for _, id := range i.message.Identifiers {
		if id != identifier {
			identifiers = append(identifiers, id)
		}
	}
	i.message.Identifiers = identifiers
	return i
}