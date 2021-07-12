package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	c := make(chan string)
	people := []string{"han", "ham", "huk", "hahahaha", "heja", "jou"}

	var s1 string = "한글"
	fmt.Println(utf8.RuneCountInString(s1))

	for _, p := range people {
		fmt.Println(p)
	}

	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is sexy"
}
