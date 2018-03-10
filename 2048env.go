//Game 2048 The following is Mengxin ZHANG's work
package main

import (
	"math/rand"
)

//reset the Game to initial step
func (g *Game2048) Reset() {
	//initialize the Gameboard as an empty board
	g.Gameboard = make([][]int, 4)
	for i := range g.Gameboard {
		g.Gameboard[i] = make([]int, 4)
	}
	g.Score = 0
	//fill in two cells randomly by 2 or 4 (0.9:0.1 probability)
	g.FillACell()
	g.FillACell()
}

//Pick a random empty cell from a list of all empty cells and fill it with 2 or 4
func (g *Game2048) FillACell() {
	cellList := g.GetAllEmpty()
	a := rand.Intn(len(cellList))
	row := cellList[a] / 4
	col := cellList[a] - row*4
	g.Gameboard[row][col] = (rand.Intn(10)/9 + 1) * 2
}

//make a list of all empty cells
func (g *Game2048) GetAllEmpty() []int {
	cellList := make([]int, 0)
	for i := range g.Gameboard {
		for j := range g.Gameboard[i] {
			if g.Gameboard[i][j] == 0 {
				cellList = append(cellList, i*4+j)
			}
		}
	}
	return cellList
}

/*
Take in direction as int: up 0, down 1, left 2, right 3
Rule: Move tiles.
a. If changed == true, then check if 2048 is achieved return false; or fill a new cell, return true;
b. if changed == false, decide if the Game can continue. If so, move again; else return false
Finally, return false if no more steps can be take; return true if Game continue
*/
func (g *Game2048) Step(direction int) bool {
	var changed bool
	if direction == 0 { //case: up
		changed = g.ShiftUp()
	} else if direction == 1 { //case: down
		changed = g.ShiftDown()
	} else if direction == 2 { //case: left
		changed = g.ShiftLeft()
	} else if direction == 3 { //case right
		changed = g.ShiftRight()
	}
	//Rule a. changed == true
	if changed == true {
		if g.Is2048() {
			return false
		} else {
			g.FillACell()
			if g.IsMovable() {
				return true
			} else {
				return false
			}
		}
	} else { //Rule b. changed == false
		return true
	}
}

//Base case
//shift each row and add up same tiles by subroutine function. Then assign back to the Gameboard
func (g *Game2048) ShiftLeft() bool {
	var changed bool
	for i := range g.Gameboard {
		tempRow := g.ShiftRow(i)
		for j := range g.Gameboard[i] {
			if tempRow[j] != g.Gameboard[i][j] {
				changed = true //if there is any change in the tempRow, mark changed as true
				g.Gameboard[i][j] = tempRow[j]
			}
		}
	}
	return changed
}

//Shift each row in the Gameboard and add up same cells
func (g *Game2048) ShiftRow(i int) []int {
	//Extract all non-zero cells first
	row := NotEmpty(g.Gameboard[i])
	// If there are more than one non-cero cells, add up same cells
	if len(row) > 1 {
		for k := range row {
			if (k+1) < len(row) && row[k] == row[k+1] {
				row[k] *= 2
				row[k+1] = 0
				g.Score += row[k]
			}
		}
		//Extract all non-zero cells a second time
		row = NotEmpty(row)
	}
	//Append zero cells to the end of tempRow
	if len(row) != 4 {
		for k := len(row); k < 4; k++ {
			row = append(row, 0)
		}
	}
	return row
}

//Extract all non-zero cells in a slice of cells
func NotEmpty(s []int) []int {
	Row := make([]int, 0)
	for i := range s {
		if s[i] != 0 {
			Row = append(Row, s[i])
		}
	}
	return Row
}

//Transpose the Gameboard, shift left, transpose back
func (g *Game2048) ShiftUp() bool {
	g.Transpose()
	changed := g.ShiftLeft()
	g.Transpose()
	return changed
}

//Transpose, shift right, transpose back
func (g *Game2048) ShiftDown() bool {
	g.Transpose()
	changed := g.ShiftRight()
	g.Transpose()
	return changed
}

//make mirror image, shift left, make mirror image
func (g *Game2048) ShiftRight() bool {
	g.Mirror()
	changed := g.ShiftLeft()
	g.Mirror()
	return changed
}

//Make a transpose of the original Gameboard. (Gameboard[i][j] = Gameboard[j][i])
func (g *Game2048) Transpose() {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			g.Gameboard[i][j], g.Gameboard[j][i] = g.Gameboard[j][i], g.Gameboard[i][j]
		}
	}
}

//make a mirror image of the original Gameboard
func (g *Game2048) Mirror() {
	for i := range g.Gameboard {
		for j := 0; j < 2; j++ {
			g.Gameboard[i][j], g.Gameboard[i][3-j] = g.Gameboard[i][3-j], g.Gameboard[i][j]
		}
	}
}

//Check whether 2048 is achieved
func (g *Game2048) Is2048() bool {
	for i := range g.Gameboard {
		for j := range g.Gameboard[i] {
			if g.Gameboard[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}

//Check whether the Gameboard is still movable
//a. not all filled b. have same neighbouring cells
func (g *Game2048) IsMovable() bool {
	for i := range g.Gameboard {
		for j := range g.Gameboard[i] {
			if g.Gameboard[i][j] == 0 {
				return true
			} else if g.InField(i+1, j) && g.Gameboard[i][j] == g.Gameboard[i+1][j] {
				return true
			} else if g.InField(i, j+1) && g.Gameboard[i][j] == g.Gameboard[i][j+1] {
				return true
			}
		}
	}
	return false
}

//check if specific cell is still in field
func (g *Game2048) InField(a, b int) bool {
	return a < 4 && a >= 0 && b < 4 && b >= 0
}
