package msutil

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type (
	Document map[string]any
	Documents []Document
)

func NewDocument() *Document {
    return &Document{}
}

func (d *Document) Set(key string, value any) *Document {
    (*d)[key] = value
    return d
}

func (d *Document) Get(key string) any {
    return (*d)[key]
}

func (d *Document) Delete(key string) *Document {
	delete(*d, key)
	return d
}

func (d *Document) Keys() []string {
	keys := make([]string, 0, len(*d))
	for k := range *d {
		keys = append(keys, k)
	}
	return keys
}

func (d *Document) Types() map[string]reflect.Type {
	types := make(map[string]reflect.Type)
	for k, v := range *d {
		types[k] = reflect.TypeOf(v)
	}
	return types
}

func (d *Document) Serialize() ([]byte, error) {
	ds := NewDocuments().Add(*d)
	return json.Marshal(ds)
}

func (d *Document) Deserialize(data []byte) error {
	ds := NewDocuments()
	err := json.Unmarshal(data, ds)
	if err != nil {
		return err
	}
	if len(*ds) == 1 {
		*d = (*ds)[0]
	}
	return nil
}

func NewDocuments() *Documents {
	return &Documents{}
}

func (d *Documents) Add(doc ...Document) *Documents {
	*d = append(*d, doc...)
	return d
}

func (d *Documents) Len() int {
	return len(*d)
}

type IndexOutOfRangeError struct {
	Length int
	Index int
}

func (e *IndexOutOfRangeError) Error() string {
	return "Index out of range: " + strconv.FormatInt(int64(e.Index), 10) + " >= " + strconv.FormatInt(int64(e.Length), 10)
}

func (d *Documents) Get(index int) (Document, error) {
	if index < 0 || index >= len(*d) {
		return nil, &IndexOutOfRangeError{Length: len(*d), Index: index}
	}
	return (*d)[index], nil
}

func (d *Documents) Delete(index int) *Documents {
	*d = append((*d)[:index], (*d)[index+1:]...)
	return d
}

func (d *Documents) Range(f func(key int, value any) bool) {
	for k, v := range *d {
		if !f(k, v) {
			break
		}
	}
}

func (d *Documents) Serialize() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Documents) Deserialize(data []byte) error {
	return json.Unmarshal(data, d)
}
