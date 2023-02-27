package concurrency

type WebsiteChecker func(string) bool

// define structure of what we want to return over the channel
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// create channel before go routine
	resultChannel := make(chan result)

	for _, url := range urls {
		// specify param 'u' to ensure each go routine gets a unique url
		go func(u string) {
			// return result of function call to channel
			resultChannel <- result{u, wc(u)}
		// pass url in to anonymous function as 'u' as explained above
		}(url)
	}
	
	// in separate for loop with length of urls, take each result and add to results map
	// this makes orderly writes in memory and avoids race condition
	for i := 0; i < len(urls); i++ {
		r := <- resultChannel
		results[r.string] = r.bool
	}

	return results
}