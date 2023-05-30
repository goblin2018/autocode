package syncx

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	f := NewSingleFlight()

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f.Do("key1", func() (any, error) {
				time.Sleep(10 * time.Millisecond)
				fmt.Println(i)
				return "val", nil
			})
		}()
	}

	wg.Wait()

	// for i := 100; i < 200; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		f.Do("key", func() (any, error) {
	// 			time.Sleep(10 * time.Millisecond)
	// 			fmt.Println(i)
	// 			return "val", nil
	// 		})
	// 	}()
	// }

	// wg.Wait()

	fmt.Println("done")

}
