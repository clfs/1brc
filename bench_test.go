package obrc

import (
	"bytes"
	"io"
	"testing"
)

func BenchmarkGenerateCSV1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCSV(io.Discard, 1e3)
	}
}

func BenchmarkGenerateCSV1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCSV(io.Discard, 1e4)
	}
}

func BenchmarkGenerateCSV1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCSV(io.Discard, 1e5)
	}
}

func BenchmarkGenerateCSV1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCSV(io.Discard, 1e6)
	}
}

func BenchmarkTakeRecordings1e3(b *testing.B) {
	var buf bytes.Buffer
	GenerateCSV(&buf, 1e3)
	data := buf.Bytes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TakeRecordings(bytes.NewReader(data))
	}
}

func BenchmarkTakeRecordings1e4(b *testing.B) {
	var buf bytes.Buffer
	GenerateCSV(&buf, 1e4)
	data := buf.Bytes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TakeRecordings(bytes.NewReader(data))
	}
}

func BenchmarkTakeRecordings1e5(b *testing.B) {
	var buf bytes.Buffer
	GenerateCSV(&buf, 1e5)
	data := buf.Bytes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TakeRecordings(bytes.NewReader(data))
	}
}

func BenchmarkTakeRecordings1e6(b *testing.B) {
	var buf bytes.Buffer
	GenerateCSV(&buf, 1e6)
	data := buf.Bytes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TakeRecordings(bytes.NewReader(data))
	}
}
