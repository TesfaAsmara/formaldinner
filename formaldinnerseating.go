package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
	"os"
	"strings"
	"math/rand"
	"time"
)

func main() {
	// Step 1: create empty collections.
	people := []string{}
	position := []string{}

	
	// 2. Open the respective csv file
	csvFile, _ := os.Open("list.csv")
	// 3. Initialize the reader
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// 4. For loop that reads each line in the csv file and other commands explained in further detail below. 
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
			log.Fatal(error)
		}
	// 5. Create another empty collection for the names. Append csv data to empty collection. 
		name := []string{}
	// Line 1 = first names, Line 0 = respective last names.
		name = append(name, line[1],line[0])
	// 6. Take slice "name", join its' elements (first and last names) with a space, and then append it into empty collection "people"
		people = append(people, strings.Join(name[:], " "))
		position = append(position, line[2])

	}

	// 7. Shuffle, randomly, the elements in slice "people".
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })

	// 8. Shuffle, randomly, the elements in slice "position".
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(position), func(i, j int) { position[i], position[j] = position[j], position[i] })

	// 9. A for loop that prints a name with a position.
		for i := range people {
			fmt.Println( people[i], position[i])
		}

	} 
	






