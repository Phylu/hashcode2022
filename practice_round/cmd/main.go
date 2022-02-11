package main

import (
	"errors"
	"fmt"
	"os"
	"simulation/internal/customer"
	"simulation/internal/simulation"
	"strconv"
	"strings"
)

func main() {
	filenames := os.Args[1:]

	// todo: iterate later
	filename := filenames[0]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	customers, err := parseFile(string(data))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// allIn := simulation.AddEverything{}
	// output := allIn.Run(customers)

	// noDislikes := simulation.NoDislikes{}
	// output := noDislikes.Run(customers)

	// likesOverDislikes := simulation.LikesOverDislikes{}
	// output := likesOverDislikes.Run(customers)

	likesOverOrEqualDislikes := simulation.LikesOverOrEqualDislikes{}
	output := likesOverOrEqualDislikes.Run(customers)

	err = writeFile(filename+".out", output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func parseFile(data string) (customer.CustomerDataset, error) {
	rows := strings.Split(data, "\n")
	dataset := customer.CustomerDataset{}

	for index, row := range rows {
		if index == 0 {
			count, err := strconv.Atoi(row)

			if err != nil {
				return customer.CustomerDataset{}, err
			}

			dataset.Count = count
			continue
		}

		if row == "" {
			continue
		}

		if index%2 != 0 {
			// likes
			likes := parseLine(row)

			customer := customer.Customer{
				Likes: likes,
			}

			dataset.Customers = append(dataset.Customers, customer)
		} else {
			// dislikes

			customerIndex := len(dataset.Customers) - 1
			customer := dataset.Customers[customerIndex]

			customer.Dislikes = parseLine(row)

			dataset.Customers[customerIndex] = customer
		}
	}

	if len(dataset.Customers) != dataset.Count {
		return customer.CustomerDataset{}, errors.New("Number of parsed customers " + strconv.Itoa(len(dataset.Customers)) + " does not equal to total number " + strconv.Itoa(dataset.Count))
	}

	return dataset, nil
}

func parseLine(line string) []string {
	return strings.Split(line, " ")[1:]
}

func writeFile(filename string, output customer.PerfectPizzaDataset) error {
	line := strconv.Itoa(len(output.Ingredients)) + " " + strings.Join(output.Ingredients, " ")

	return os.WriteFile(filename, []byte(line), 0644)
}
