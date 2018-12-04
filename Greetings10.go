package main

import "fmt"

func showGreetings(n []string) {

	for _, c := range n {
		fmt.Println("Name:", c)
	}

}

func main() {

	n1 := []string{"Yani", "Vlad", "Krasi", "Ivan", "Kris", "Naso", "Dimi", "Geri", "Lili", "Zoya"}
	showGreetings(n1)

	n2 := []string{"Yani"}
	showGreetings(n2)
}
