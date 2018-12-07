package main

import "fmt"

func showGreetings(n []string) {

	for _, c := range n {
		fmt.Println("Name:", c)
	}

}

func showAge(y []string) {

	for _, i := range y {
		fmt.Println("Age:", i)
	}
}

func main() {

	n1 := []string{"Yani", "Vlad", "Krasi", "Ivan", "Kris", "Naso", "Dimi", "Geri", "Lili", "Zoya"}
	showGreetings(n1)

	n2 := []string{"Yani"}
	showGreetings(n2)

	y := []string{"36", "39", "20", "25", "41", "30", "15", "11", "22", "19"}
	showAge(y)

}
