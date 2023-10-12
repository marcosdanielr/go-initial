package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// METHODS USING CAMELCASE ARE PRIVATE AND PASCALCASE ARE PUBLIC

type Animal struct {
	Name   string `json:"name"`
	Specie string `json:"specie"`
	Age    int    `json:"age"`
}

func (a Animal) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Specie: %s, Age: %d", a.Name, a.Specie, a.Age)
}

// func counter() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func worker(workerID int, data chan int) {
// 	for x := range data {
// 		fmt.Printf("Worker %d received %d\n", workerID, x)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	channel := make(chan int)
// 	// go worker(1, channel) // T2
// 	// go worker(2, channel) // T3
// 	// go worker(3, channel) // T4

// 	for i := 1; i <= 100000; i++ {
// 		go worker(i, channel) // Ti, 2k
// 	}

// 	for i := 0; i < 200000; i++ {
// 		channel <- i
// 	}
// }

// T1
// func main() {
// 	channel := make(chan string)

// 	// T2
// 	go func() {
// 		channel <- "Aoba"
// 	}()

// 	fmt.Println(<-channel)
// }

// go routine
// T0
// func main() {
// 	go counter() // T
// 	go counter() // T2
// 	counter()

// 	http.HandleFunc("/", home)
// 	http.ListenAndServe(":3002", nil)
// }

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3002", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	animal := Animal{
		Name:   "Ludi",
		Specie: "Dog",
		Age:    2,
	}

	json.NewEncoder(w).Encode(animal)
}
