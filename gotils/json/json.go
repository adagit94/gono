package json

import (
	"github.com/adagit94/gono/gotils/fs"
	"github.com/bytedance/sonic"
	"os"
)

func Json[T any](v *T) ([]byte, error) {
	return sonic.Marshal(v)
}

func JsonToFile(v *any, perm os.FileMode, pathSegments ...string) error {
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

func JsonStr(v *any) (string, error) {
	return sonic.MarshalString(v)
}

func JsonParse[T any](source []byte, target *T) error {
	return sonic.Unmarshal(source, target)
}

func JsonParseStr[T any](source string, target *T) error {
	return sonic.UnmarshalString(source, target)
}

