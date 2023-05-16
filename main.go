package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	r, err := NewRoom("https://live.douyin.com/405518163654")
	if err != nil {
		panic(err)
	}
	logrus.Info(r.Connect())
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
