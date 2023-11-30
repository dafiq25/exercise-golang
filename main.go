package main

import (
	"fmt"
)

type Mobil struct {
	tire       [4]string
	rightDoor string
	leftDoor  string
}

func main() {
	mobil := Mobil{
		tire:       [4]string{"karet", "karet", "karet", "karet"},
		rightDoor: "Knock",
		leftDoor:  "beep",
	}

	fmt.Println("Bunyi pintu kanan ketika diketuk:", mobil.rightDoor)
	fmt.Println("Bunyi pintu kiri ketika diketuk:", mobil.bunyileftDoor())

	tireBaru := [4]string{"kayu", "besi", "karet", "besi"}
	mobil.gantitire(tireBaru)

	fmt.Println("Bunyi pintu kanan ketika dibuka:", mobil.rightDoor)
	fmt.Println("Bunyi pintu kiri ketika dibuka:", mobil.bunyileftDoor())
}

func (m *Mobil) gantitire(tireBaru [4]string) {
	if len(tireBaru) == 4 {
		m.tire = tireBaru
		fmt.Println("tire mobil telah diganti.")
	} else {
		fmt.Println("Jumlah tire tidak sesuai. Harap ganti dengan 4 tire.")
	}
}

func (m Mobil) bunyileftDoor() string {
	if m.rightDoor == "Knock" {
		return "beep"
	}
	return "Knock"
}
