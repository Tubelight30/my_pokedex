package main

type config struct {
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {

	cfg := config{
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}
	startRepl(&cfg)

	// ticker := time.NewTicker(500 * time.Millisecond)
	// for {

	// 	t := <-ticker.C
	// 	fmt.Println(t)
	// }

}
