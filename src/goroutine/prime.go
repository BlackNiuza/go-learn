package main

import "fmt"

func main() {
	var ch chan int = make(chan int)
	go geneNum(ch)

	for {
		if prime, ok := <-ch; !ok {
			break
		} else {
			fmt.Printf("%d ", prime)
			ch1 := make(chan int)

			go filter(ch, ch1, prime)
			ch = ch1
		}
	}
}

func geneNum(ch chan int) {
	for i := 2; i < 20; i++ {
		ch <- i
	}
	chClose(ch)
}

func filter(in <-chan int, out chan int, prime int) {
	for {
		if num, ok := <-in; !ok {
			break
		} else {
			if num%prime != 0 {
				out <- num
			}
		}
	}
	chClose(out)
}

func chClose(ch chan int) {
	close(ch)
	fmt.Printf("close ch=%v\n", ch)
}
