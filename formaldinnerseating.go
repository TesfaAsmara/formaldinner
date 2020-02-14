package main

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	FirstName string
	LastName string
}

const numberOfStudents = 290
const numberOfTables = 31
const numberOfWaiters = 31
const numberOfWaiterAlt = 8
const numberOfKitchenCrew = 8
const numberOfShuffles = 3


func Shuffle(students []Student) {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    // We start at the end of the slice, inserting our random
    // values one at a time.
    for n := len(students); n > 0; n-- {
        randIndex := r.Intn(n)
        // We swap the value at index n-1 and the random index
        // to move our randomly chosen value to the end of the
        // slice, and to move the value that was at n-1 into our
        // unshuffled portion of the slice.
        students[n-1], students[randIndex] = students[randIndex], students[n-1]
    }
}

func main() {
		// Get list of students into this slice
		unshuffledStudents := make([]Student, 0, numberOfStudents)

		// Open the respective csv file
		csvFile, _ := os.Open("list.csv")
	
		// Initialize the reader
		reader := csv.NewReader(bufio.NewReader(csvFile))
	
		// For loop that reads lines 0 and 1, first and last names, from the csv file. 
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			unshuffledStudents = append(unshuffledStudents, Student{
				FirstName: line[1],
				LastName:  line[0],
			})
		}

		for i := 1; i <= numberOfShuffles; i++ {

		Shuffle(unshuffledStudents)

		fmt.Println("------------------------------------------This is seating", i)

			for index, student := range unshuffledStudents {
				if index < numberOfWaiters {
					 fmt.Printf("%s %s Waiter %d\n", student.FirstName, student.LastName, index + 1)
				} 
				if index < numberOfWaiters + numberOfWaiterAlt && index > numberOfWaiters {
					 fmt.Printf("%s %s Waiter Alt\n", student.FirstName, student.LastName)
				}
				if index < numberOfWaiters + numberOfWaiterAlt + numberOfKitchenCrew && index > numberOfWaiters + numberOfWaiterAlt {
					 fmt.Printf("%s %s Kitchen Crew\n", student.FirstName, student.LastName)
				}
				if index > numberOfWaiters + numberOfWaiterAlt + numberOfKitchenCrew {
					 fmt.Printf("%s %s %d\n", student.FirstName, student.LastName, (index - (numberOfWaiters + numberOfWaiterAlt + numberOfKitchenCrew)) % numberOfTables + 1)
				}
		}
}
}