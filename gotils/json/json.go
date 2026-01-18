package json

import (
	"github.com/adagit94/gono/gotils/fs"
	"github.com/bytedance/sonic"
	"os"
)

func Json[V any](v *V) ([]byte, error) {
	return sonic.Marshal(v)
}

func JsonToFile[V any](v *V, perm os.FileMode, pathSegments ...string) error {
	s, err := Json(v)

	if err != nil {
		return err
	}

	err2 := fs.WriteFile(s, perm, pathSegments...)

	if err2 != nil {
		return err2
	}

	return nil
}

func JsonStr[V any](v *V) (string, error) {
	return sonic.MarshalString(v)
}

func JsonParse[T any](source []byte, target *T) error {
	return sonic.Unmarshal(source, target)
}

func JsonParseStr[T any](source string, target *T) error {
	return sonic.UnmarshalString(source, target)
}
