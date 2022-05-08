package problems

const (
	INF  = 2147483647
	Gate = 0
	Wall = -1
)

type Room struct {
	row int
	column int
}

var directions = [4][2]int{
	{ -1, 0 }, // North
	{ 1, 0 },  // South
	{ 0, 1 },  // East
	{ 0, -1 }, // West
}

// wallsAndGates takes a grid of ints representing gates, walls, and empty rooms.
// Grid types are represented as follows: gates = 0, walls = -1, empty rooms = (INF) 2147483647.
// Empty rooms will be set to an int denoting the distance to the nearest gate, or INF if not reachable.
// Processing of rooms is done by BFS.
func wallsAndGates(rooms [][]int)  {
	roomQueue := GetGateRooms(rooms)
	roomsVisited := make(map[Room]bool)
	distance := 1

	for len(roomQueue) > 0 {
		queueSize := len(roomQueue)

		for index := 0; index < queueSize; index++ {
			room := roomQueue[index]
			emptyNeighboringRooms := room.GetEmptyNeighboringRooms(rooms)

			for _, emptyRoom := range emptyNeighboringRooms {
				_, visited := roomsVisited[emptyRoom]

				if !visited {
					roomQueue = append(roomQueue, emptyRoom)
					rooms[emptyRoom.row][emptyRoom.column] = distance
					roomsVisited[emptyRoom] = true
				}
			}
		}

		distance++
		roomQueue = roomQueue[queueSize:]
	}
}

func GetGateRooms(rooms [][]int) []Room {
	gateRooms := make([]Room, 0, 0)

	for rowIndex := 0; rowIndex < len(rooms); rowIndex++ {
		for columnIndex := 0; columnIndex < len(rooms[rowIndex]); columnIndex++ {
			if rooms[rowIndex][columnIndex] == Gate {
				gateRooms = append(gateRooms, Room {
					row: rowIndex,
					column: columnIndex,
				})
			}
		}
	}

	return gateRooms
}

func (currentRoom Room) GetEmptyNeighboringRooms(rooms [][]int) []Room {
	neighboringRooms := make([]Room, 0, 0)

	for _, direction := range directions {
		nextRoom := Room {
			row: currentRoom.row + direction[0],
			column: currentRoom.column + direction[1],
		}

		gridHeight := len(rooms)
		gridWidth := len(rooms[0])

		if nextRoom.row < 0 ||
			nextRoom.row > gridHeight - 1 ||
			nextRoom.column < 0 ||
			nextRoom.column > gridWidth - 1 ||
			rooms[nextRoom.row][nextRoom.column] != INF {
			continue
		}

		neighboringRooms = append(neighboringRooms, nextRoom)
	}

	return neighboringRooms
}
