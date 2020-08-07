package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1) // 运行时长
	ch := make(chan int)

	go func() {
		for {
			select {
				case <-ticker.C:
				//logic
			}
		}
		ticker.Stop()
		ch <- 0
	}()
	<-ch
}

// @ Chenyu Zhao
// function:

func function() {

}


