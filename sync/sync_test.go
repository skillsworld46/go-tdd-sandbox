package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.Value(), 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCnt := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCnt)

		for i := 0; i < wantedCnt; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter.Value(), wantedCnt)
	})
}

func assertCounter(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
