//paralellität

package main

import (
	"fmt"
	"time"
)

func expert(band string, bands chan<- string, timeToPick int) {
	for {
		fmt.Println("Ich empfehle:", band)
		bands <- band
		time.Sleep(time.Duration(timeToPick) * time.Second)
	}
}

func buyer(i int, bands <-chan string) {
	for band := range bands {
		fmt.Println("Käuferin", i, ": Jo, die Platte von den", band, "kaufe ich.")
		time.Sleep(time.Second * 2)
	}
}

func hearer(i int, bands <-chan string) {
	for band := range bands {
		fmt.Println("Hörerin", i, ": Wahnsinn, diese", band,".")
		time.Sleep(time.Second * 2)
	}
}

func critic(i int, bands <-chan string) {
	for band := range bands {
		fmt.Println("Kritikerin", i, ": Die", band, "waren früher besser.")
		time.Sleep(time.Second * 2)
	}
}

func main() {
	bands := make(chan string, 1000)

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Millisecond * 600)
		go hearer(i, bands)
		go buyer(i, bands)
		go critic(i, bands)
	}

	go expert("Kinks", bands, 2)
	time.Sleep(time.Millisecond * 300)
	go expert("Beach Boys", bands, 1)
	time.Sleep(time.Millisecond * 600)
	go expert("Stones", bands, 5)
	time.Sleep(time.Millisecond * 600)
	go expert("Beatles", bands, 3)
	time.Sleep(time.Millisecond * 600)


	time.Sleep(time.Second * 2)
}