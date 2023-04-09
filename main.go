package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bxn0/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose convert, average or exit, You have to choose one of these three")

		Venligst velg convert, average eller exit:

		if !scanner.Scan() {
			break
		}
		input = scanner.Text()

		switch input {
		case "q", "exit":
			fmt.Println("exit")
			return

		case "convert": 
			fmt.Println("Converting C values to Fahreinheit")
			
		
			yr.ConvertTemperature()

		case "average":
			fmt.Println("Calculation of the average")

			for {
				
				yr.AverageTemperature()

				var input2 string
				scanjn := bufio.NewScanner(os.Stdin)
				fmt.Println("Do you want to go back to the main? (j/n)")
				for scanjn.Scan() {
					input2 = scanjn.Text()
					if input2 == "j" {
						break
					} else if input2 == "n" {
						break
					}
				}
				if input2 == "j" {
					break
				}
			}
		}
	}
	fmt.Println("Terminated...")
}
