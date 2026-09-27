package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"CA2/hotel"
)

func main() {

	// Check command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . input.txt")
		return
	}

	// Get filename from command line
	filename := os.Args[1]

	// Load old data
	err := hotel.LoadRooms(filename)

	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	fmt.Println("Old room data loaded successfully.")

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n===== HOTEL ROOM BOOKING REGISTRY =====")
		fmt.Println("1. Register Room")
		fmt.Println("2. Find Room")
		fmt.Println("3. Toggle Occupancy")
		fmt.Println("4. Remove Room")
		fmt.Println("5. Display All Rooms")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		choice, err := strconv.Atoi(choiceInput)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		switch choice {

		case 1:

			fmt.Print("Enter Room Number: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			roomNumber, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid room number.")
				continue
			}

			fmt.Print("Enter Room Type: ")

			roomType, _ := reader.ReadString('\n')
			roomType = strings.TrimSpace(roomType)

			fmt.Print("Enter Rate Per Night: ")

			rateInput, _ := reader.ReadString('\n')
			rateInput = strings.TrimSpace(rateInput)

			rate, err := strconv.ParseFloat(rateInput, 64)

			if err != nil {
				fmt.Println("Invalid rate.")
				continue
			}

			room := hotel.Room{
				RoomNumber:   roomNumber,
				RoomType:     roomType,
				RatePerNight: rate,
				IsOccupied:   false,
			}

			err = hotel.RegisterRoom(room)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				err = hotel.SaveRooms(filename)

				if err != nil {
					fmt.Println("Room registered, but could not save:", err)
				} else {
					fmt.Println("Room registered successfully!")
				}
			}

		case 2:

			fmt.Print("Enter Room Number to find: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			roomNumber, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid room number.")
				continue
			}

			room, err := hotel.FindRoom(roomNumber)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				fmt.Println("\n===== ROOM FOUND =====")
				fmt.Println("Room Number :", room.RoomNumber)
				fmt.Println("Room Type   :", room.RoomType)
				fmt.Println("Rate/Night  :", room.RatePerNight)
				fmt.Println("Occupied    :", room.IsOccupied)
			}

		case 3:

			fmt.Print("Enter Room Number: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			roomNumber, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid room number.")
				continue
			}

			fmt.Print("Enter occupancy status (true/false): ")

			statusInput, _ := reader.ReadString('\n')
			statusInput = strings.TrimSpace(statusInput)

			occupied, err := strconv.ParseBool(statusInput)

			if err != nil {
				fmt.Println("Enter only true or false.")
				continue
			}

			err = hotel.ToggleOccupancy(roomNumber, occupied)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				err = hotel.SaveRooms(filename)

				if err != nil {
					fmt.Println("Occupancy updated, but could not save:", err)
				} else {
					fmt.Println("Occupancy updated successfully!")
				}
			}

		case 4:

			fmt.Print("Enter Room Number to remove: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			roomNumber, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid room number.")
				continue
			}

			err = hotel.RemoveRoom(roomNumber)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				err = hotel.SaveRooms(filename)

				if err != nil {
					fmt.Println("Room removed, but could not save:", err)
				} else {
					fmt.Println("Room removed successfully!")
				}
			}

		case 5:

			hotel.DisplayRooms()

		case 6:

			err := hotel.SaveRooms(filename)

			if err != nil {
				fmt.Println("Error saving data:", err)
			} else {
				fmt.Println("Data saved successfully.")
			}

			fmt.Println("Thank you!")

			return

		default:

			fmt.Println("Invalid choice. Please select 1-6.")
		}
	}
}
