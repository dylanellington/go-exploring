package _200_number_of_islands

var directions = [4][2]int{
	{ -1, 0 }, // North
	{ 1, 0 },  // South
	{ 0, 1 },  // East
	{ 0, -1 }, // West
}

var island = "1"

func numIslands(grid [][]byte) int {
	islandMap := make(map[int]bool)
	numIslands := 0

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for columnIndex := 0; columnIndex < len(grid[0]); columnIndex++ {
			islandMapIndex := (rowIndex * len(grid[0])) + columnIndex
			_, visited := islandMap[islandMapIndex]

			if !visited && string(grid[rowIndex][columnIndex]) == island {
				numIslands++
				islandMap[islandMapIndex] = true
				ExploreIsland(grid, rowIndex, columnIndex, islandMap)
			} else {
				islandMap[islandMapIndex] = true
			}
		}
	}

	return numIslands
}

func ExploreIsland(grid [][]byte, row int, column int, islandMap map[int]bool) {
	gridHeight := len(grid)
	gridWidth := len(grid[0])

	for _, direction := range directions {
		newRow := row + direction[0]
		newColumn := column + direction[1]

		if newRow < 0 ||
			newRow > gridHeight - 1 ||
			newColumn < 0 ||
			newColumn > gridWidth - 1 {
			continue
		}

		islandMapIndex := (newRow * len(grid[0])) + newColumn
		_, visited := islandMap[islandMapIndex]

		if !visited && string(grid[newRow][newColumn]) == island {
			islandMap[islandMapIndex] = true
			ExploreIsland(grid, newRow, newColumn, islandMap)
		}
	}
}
