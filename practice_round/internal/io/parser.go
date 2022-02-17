package io

import (
	"errors"
	"simulation/internal/models"
	"strconv"
	"strings"
)

func parseFile(data string) (models.CustomerDataset, error) {
	rows := strings.Split(data, "\n")
	dataset := models.CustomerDataset{}

	for index, row := range rows {
		if index == 0 {
			count, err := strconv.Atoi(row)

			if err != nil {
				return models.CustomerDataset{}, err
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

			customer := models.Customer{
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
		return models.CustomerDataset{}, errors.New("Number of parsed customers " + strconv.Itoa(len(dataset.Customers)) + " does not equal to total number " + strconv.Itoa(dataset.Count))
	}

	return dataset, nil
}

func parseLine(line string) []string {
	return strings.Split(line, " ")[1:]
}
