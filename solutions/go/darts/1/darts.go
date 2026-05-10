package darts

func Score(x, y float64) int {
	distSquared := x*x + y*y

    // Check the Inner Circle first
    if distSquared <= 1 {
        return 10
    }
    // Then check the Middle Circle
    if distSquared <= 25 {
        return 5
    }
    // Then check the Outer Circle
    if distSquared <= 100 {
        return 1
    }

    // Otherwise, you missed the board!
    return 0
}
