package events

import "testing"

func TestLoad(t *testing.T) {
	t.Run("Test Load", func(t *testing.T) {
		db, err := Get()

		if err != nil {
			t.Fatalf("Load: %v", err)
		}

		if len(db) == 0 {
			t.Fatal("database empty")
		}
	})
}

func BenchmarkLoad(b *testing.B) {
	b.Run("Benchmark Load", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := Get(); err != nil {
				b.Fatalf("Load: %v", err)
			}
		}
	})
}
