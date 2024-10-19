package serializer

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}
