package meilisearchnatsconnector

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}
