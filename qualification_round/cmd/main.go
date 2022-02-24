package main

import (
	"fmt"
	"os"
	"simulation/internal/io"
	"simulation/internal/simulation"
)

func main() {
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		fmt.Println("There are no input files given")
		os.Exit(1)
	}

	for _, filename := range filenames {
		dataset, err := io.ReadFile(filename)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		allIn := simulation.FirstContributorsInOrder{}
		output := allIn.Run(dataset)

		err = io.WriteFile(filename+".out", output)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(`Simulation for "` + filename + `" complete`)
	}

	fmt.Println("Simulation complete ✅")
}
