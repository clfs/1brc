package obrc

import (
	"bytes"
	"testing"
)

func FuzzTakeRecordings(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = TakeRecordings(NewReader(bytes.NewReader(data)))
	})
}
