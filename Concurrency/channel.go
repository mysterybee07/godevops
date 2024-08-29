package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	userID := 10

	respChannel := make(chan string)
	wg := &sync.WaitGroup{}
	go FetchUserData(userID, respChannel, wg)
	go FetchUserRecommendation(userID, respChannel, wg)
	go FetchUserLikes(userID, respChannel, wg)

	wg.Add(3)

	go func() {
		wg.Wait()
		close(respChannel)
	}()
	// close(respChannel)
	for resp := range respChannel {
		fmt.Println(resp)
	}

	// fmt.Println(data)
	// fmt.Println(recom)
	// fmt.Println(likes)
	fmt.Println(time.Since(now))
}

func FetchUserData(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 80)
	// return "user Data"
	respCh <- "user data"
	wg.Done()
}

func FetchUserRecommendation(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	//return "user recommendation"
	respCh <- "user recommendation"
	wg.Done()
}

func FetchUserLikes(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respCh <- "user Likes"
	wg.Done()
}
