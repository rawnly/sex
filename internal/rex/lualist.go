package rex

import (
	"bytes"
	"encoding/json"
)

// luaList decodes a list from a Lua script; Lua encodes an empty table as {}.
type luaList[T any] []T

func (l *luaList[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(bytes.TrimSpace(b), []byte("{}")) {
		*l = nil
		return nil
	}
	return json.Unmarshal(b, (*[]T)(l))
}
