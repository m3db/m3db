package storage

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestThrottler(t *testing.T) {
	throttler := &Throttler{
		keyState:        make(map[string]*keyContext, 0),
		keyQueue:        list.New(),
		globalMaxClaims: 2,
	}

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			// Simulate distinct small user.
			u1 := fmt.Sprintf("user_%d", i)
			claim1, err := throttler.Acquire(u1)
			require.NoError(t, err)
			time.Sleep(time.Millisecond * 200)
			claim1.Release()
			wg.Done()
		}()

		go func() {
			// Simulate distinct small user.
			u2 := "user_bad"
			claim2, err := throttler.Acquire(u2)
			require.NoError(t, err)
			time.Sleep(time.Millisecond * 200)
			claim2.Release()
			wg.Done()
		}()
	}

	wg.Wait()
}