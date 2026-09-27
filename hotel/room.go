package hotel

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Room struct {
	RoomNumber   int
	RoomType     string
	RatePerNight float64
	IsOccupied   bool
}

var rooms []Room

func RegisterRoom(r Room) error {
	for _, room := range rooms {
		if room.RoomNumber == r.RoomNumber {
			return fmt.Errorf("room number already exists")
		}
	}

	if r.RatePerNight <= 0 {
		return fmt.Errorf("rate must be greater than 0")
	}

	rooms = append(rooms, r)

	return nil
}

func FindRoom(roomNum int) (*Room, error) {
	for i := range rooms {
		if rooms[i].RoomNumber == roomNum {
			return &rooms[i], nil
		}
	}

	return nil, fmt.Errorf("room not found")
}

func ToggleOccupancy(roomNum int, occupied bool) error {
	for i := range rooms {
		if rooms[i].RoomNumber == roomNum {
			rooms[i].IsOccupied = occupied
			return nil
		}
	}

	return fmt.Errorf("room not found")
}

func RemoveRoom(roomNum int) error {
	for i := range rooms {
		if rooms[i].RoomNumber == roomNum {
			rooms = append(rooms[:i], rooms[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("room not found")
}

func DisplayRooms() {
	if len(rooms) == 0 {
		fmt.Println("No rooms registered.")
		return
	}

	fmt.Println("\n===== ALL ROOMS =====")

	for _, room := range rooms {
		fmt.Println("-------------------------")
		fmt.Println("Room Number :", room.RoomNumber)
		fmt.Println("Room Type   :", room.RoomType)
		fmt.Println("Rate/Night  :", room.RatePerNight)
		fmt.Println("Occupied    :", room.IsOccupied)
	}
}

// LoadRooms reads old room data from the input file.
func LoadRooms(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		// If file does not exist, start with empty rooms.
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// Remove anything currently in memory.
	rooms = nil

	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		if len(record) != 4 {
			continue
		}

		roomNumber, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		rate, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}

		occupied, err := strconv.ParseBool(record[3])
		if err != nil {
			continue
		}

		room := Room{
			RoomNumber:   roomNumber,
			RoomType:     record[1],
			RatePerNight: rate,
			IsOccupied:   occupied,
		}

		rooms = append(rooms, room)
	}

	return nil
}

// SaveRooms saves all rooms into the input file.
func SaveRooms(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for _, room := range rooms {
		record := []string{
			strconv.Itoa(room.RoomNumber),
			room.RoomType,
			strconv.FormatFloat(room.RatePerNight, 'f', 2, 64),
			strconv.FormatBool(room.IsOccupied),
		}

		err := writer.Write(record)

		if err != nil {
			return err
		}
	}

	writer.Flush()

	return writer.Error()
}
