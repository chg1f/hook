package hook

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestHook(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()

	startAt := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer Start(startAt, func(c time.Time) {
			fmt.Println("startAt:" + c.Format(time.RFC3339Nano))
			assert.Equal(t, startAt, c)
		}).Stop(func(c time.Time) {
			stopAt := time.Now()
			elapsed := stopAt.Sub(c)
			fmt.Println("stopAt:" + stopAt.Format(time.RFC3339Nano))
			fmt.Println("elapsed:" + elapsed.String())
			assert.Equal(t, startAt, c)
			assert.Equal(t, int64(elapsed.Seconds()), int64(time.Second.Seconds()))
		})
		time.Sleep(time.Second)
	}()
}

func TestEmpty(t *testing.T) {
	IgnoreHook = true
	var wg sync.WaitGroup
	defer wg.Wait()

	startAt := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer Start(startAt, func(c time.Time) {
			fmt.Println("startAt:" + c.Format(time.RFC3339Nano))
			assert.Equal(t, startAt, c)
		}).Stop(func(c time.Time) {
			stopAt := time.Now()
			elapsed := stopAt.Sub(c)
			fmt.Println("stopAt:" + stopAt.Format(time.RFC3339Nano))
			fmt.Println("elapsed:" + elapsed.String())
			assert.Equal(t, startAt, c)
			assert.Equal(t, elapsed.Seconds(), time.Second)
		})
		time.Sleep(time.Second)
	}()
}
