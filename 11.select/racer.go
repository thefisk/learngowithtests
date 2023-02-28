package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select allows you to wait on multiple channels. first one to return a value, 'wins'
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	//time.After() can be a useful case to stop potential code that might block forever
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// function returns a chan struct{} because chan struct{} is the smallest data type
// the return type/contents don't matter in this case - it's purely to signal the func has completed
func ping (url string) chan struct{} {
	// use make to create a channel.  if using var, it is initialised with zero-value, which is nil for a chan and it will block forever
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}