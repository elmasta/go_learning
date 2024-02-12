package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(door Door) bool {
	PrintStr("Door Closing...\n")
	door.state = false
	return door.state
}

func OpenDoor(door Door) bool {
	PrintStr("Door Opening...\n")
	door.state = true
	return door.state
}

func IsDoorOpen(door Door) bool {
	PrintStr("is the Door opened ?\n")
	door.state = true
	return door.state
}

func IsDoorClose(door Door) bool {
	PrintStr("is the Door closed ?\n")
	door.state = false
	return door.state
}

func main() {
	var door Door
	door.state = true

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if !door.state {
		CloseDoor(door)
	}
}
