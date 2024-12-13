package main

const (
	// Directions
	Up    = '^'
	Right = '>'
	Down  = 'v'
	Left  = '<'
	// Obstacles
	Obstacle = '#'
	// Visited mark
	Visited = 'X'
)

type Position struct {
	y         int
	x         int
	direction rune
}

func NewPosition(y int, x int, direction rune) Position {
	return Position{
		y:         y,
		x:         x,
		direction: direction,
	}
}

type Cell struct {
	Char    rune
	Visited bool
}

func NewCell(char rune) Cell {
	return Cell{
		Char:    char,
		Visited: false,
	}
}

func (c Cell) IsObstacle() bool {
	return c.Char == Obstacle
}

func (c *Cell) MarkVisited() {
	(*c).Char = Visited
	(*c).Visited = true
}

type Map struct {
	Cells [][]Cell
	Pos   Position
}

func (m Map) String() string {
	str := ""
	for row := 0; row < len(m.Cells); row++ {
		for col := 0; col < len(m.Cells[row]); col++ {
			str += string(m.Cells[row][col].Char)
		}
		str += "\n"
	}
	return str
}

// func (m Map) PrintMap() {
// 	fmt.Println(m)
// }

func (m *Map) GetPosition() *Cell {
	return &(*m).Cells[(*m).Pos.y][(*m).Pos.x]
}

func (m *Map) SetDirection(direction rune) {
	switch direction {
	case Up, Right, Down, Left:
		(*m).Pos.direction = direction
		(*m).GetPosition().Char = direction
	}
}

func (m *Map) NextUp() Cell {
	return (*m).Cells[(*m).Pos.y-1][(*m).Pos.x]
}

func (m *Map) NextRight() Cell {
	return (*m).Cells[(*m).Pos.y][(*m).Pos.x+1]
}

func (m *Map) NextDown() Cell {
	return (*m).Cells[(*m).Pos.y+1][(*m).Pos.x]
}

func (m *Map) NextLeft() Cell {
	return (*m).Cells[(*m).Pos.y][(*m).Pos.x-1]
}
