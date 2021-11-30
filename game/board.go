package game

import (
	"math/rand"
	"time"
)

const (
	_boardSize = 4
	_emptyCell = 0
	_wonCell   = 2048
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type board struct {
	cells Cells
	score uint32
	won   bool
}

//
// Controller interface implementation
//

var _ Controller = (*board)(nil)

func (b *board) Shift(direction Direction) bool { return b.shift(direction) }
func (b *board) Won() bool                      { return b.won }
func (b *board) Lost() bool                     { return b.noMovesRemaining() }
func (b *board) GetScore() uint32               { return b.score }
func (b *board) GetCells() Cells                { return b.cells }
func (b *board) Reset()                         { *b = initNewBoard() }

//
// internal methods
//

func initNewBoard() board {
	b := board{}
	// add two random cells
	b.fillRandom()
	b.fillRandom()
	return b
}

// shift cells in the given direction and fill a random cell if the board has changed
func (b *board) shift(direction Direction) (hasChanged bool) {
	switch direction {
	case DirectionLeft:
		hasChanged = b.shiftLeft()
	case DirectionRight:
		hasChanged = b.shiftRight()
	case DirectionUp:
		hasChanged = b.shiftUp()
	case DirectionDown:
		hasChanged = b.shiftDown()
	}
	if hasChanged {
		b.fillRandom()
	}
	return
}

// for each row, merge each column left
func (b *board) shiftLeft() (hasChanged bool) {
	for rowIdx := 0; rowIdx < _boardSize; rowIdx++ {
		var (
			newColIdx      int
			lastCellMerged bool
		)
		for colIdx := 0; colIdx < _boardSize; colIdx++ {
			curCell := b.getCell(rowIdx, colIdx)
			if curCell != _emptyCell {
				if newColIdx > 0 && b.getCell(rowIdx, newColIdx-1) == curCell && !lastCellMerged {
					b.doubleCell(rowIdx, newColIdx-1)
					lastCellMerged = true
					hasChanged = true
				} else {
					if b.setCell(rowIdx, newColIdx, curCell) {
						hasChanged = true
					}
					newColIdx++
					lastCellMerged = false
				}
			}
		}
		for ; newColIdx < _boardSize; newColIdx++ {
			b.setCell(rowIdx, newColIdx, _emptyCell)
		}
	}
	return
}

// for each row, merge each column right
func (b *board) shiftRight() (hasChanged bool) {
	for rowIdx := _boardSize - 1; rowIdx >= 0; rowIdx-- {
		var (
			newColIdx      = _boardSize - 1
			lastCellMerged bool
		)
		for colIdx := _boardSize - 1; colIdx >= 0; colIdx-- {
			curCell := b.getCell(rowIdx, colIdx)
			if curCell != _emptyCell {
				if newColIdx < _boardSize-1 && b.getCell(rowIdx, newColIdx+1) == curCell && !lastCellMerged {
					b.doubleCell(rowIdx, newColIdx+1)
					lastCellMerged = true
					hasChanged = true
				} else {
					if b.setCell(rowIdx, newColIdx, curCell) {
						hasChanged = true
					}
					newColIdx--
					lastCellMerged = false
				}
			}
		}
		for ; newColIdx >= 0; newColIdx-- {
			b.setCell(rowIdx, newColIdx, _emptyCell)
		}
	}
	return
}

// for each column, merge each row up
func (b *board) shiftUp() (hasChanged bool) {
	for colIdx := 0; colIdx < _boardSize; colIdx++ {
		var (
			newRowIdx      int
			lastCellMerged bool
		)
		for rowIdx := 0; rowIdx < _boardSize; rowIdx++ {
			curCell := b.getCell(rowIdx, colIdx)
			if curCell != _emptyCell {
				if newRowIdx > 0 && b.getCell(newRowIdx-1, colIdx) == curCell && !lastCellMerged {
					b.doubleCell(newRowIdx-1, colIdx)
					lastCellMerged = true
					hasChanged = true
				} else {
					if b.setCell(newRowIdx, colIdx, curCell) {
						hasChanged = true
					}
					newRowIdx++
					lastCellMerged = false
				}
			}
		}
		for ; newRowIdx < _boardSize; newRowIdx++ {
			b.cells[newRowIdx][colIdx] = _emptyCell
		}
	}
	return
}

// for each column, merge each row down
func (b *board) shiftDown() (hasChanged bool) {
	for colIdx := _boardSize - 1; colIdx >= 0; colIdx-- {
		var (
			newRowIdx      = _boardSize - 1
			lastCellMerged bool
		)
		for rowIdx := _boardSize - 1; rowIdx >= 0; rowIdx-- {
			curCell := b.getCell(rowIdx, colIdx)
			if curCell != _emptyCell {
				if newRowIdx < _boardSize-1 && b.getCell(newRowIdx+1, colIdx) == curCell && !lastCellMerged {
					b.doubleCell(newRowIdx+1, colIdx)
					lastCellMerged = true
					hasChanged = true
				} else {
					if b.setCell(newRowIdx, colIdx, curCell) {
						hasChanged = true
					}
					lastCellMerged = false
					newRowIdx--
				}
			}
		}
		for ; newRowIdx >= 0; newRowIdx-- {
			b.setCell(newRowIdx, colIdx, _emptyCell)
		}
	}
	return
}

func (b *board) getCell(row, col int) uint16 {
	return b.cells[row][col]
}

// doubleCell double the cell value, mark if it's a winning cell, and update the score
func (b *board) doubleCell(row, col int) {
	b.cells[row][col] <<= 1 // double the uint16 using bitshift left
	if b.cells[row][col] == _wonCell {
		b.won = true
	}
	b.score += (uint32)(b.cells[row][col])
}

// setCell sets a cell value, if the cell is already set to the given value, the boolean returned will be false
func (b *board) setCell(row, col int, val uint16) bool {
	if b.cells[row][col] != val {
		b.cells[row][col] = val
		return true
	}
	return false
}

func (b *board) noMovesRemaining() bool {
	for row := 0; row < _boardSize; row++ {
		for col := 0; col < _boardSize; col++ {
			curCell := b.getCell(row, col)
			if curCell == _emptyCell {
				return false
			}
			// up
			if row > 0 && b.getCell(row-1, col) == curCell {
				return false
			}
			// down
			if row < _boardSize-1 && b.getCell(row+1, col) == curCell {
				return false
			}
			// left
			if col > 0 && b.getCell(row, col-1) == curCell {
				return false
			}
			// right
			if col < _boardSize-1 && b.getCell(row, col+1) == curCell {
				return false
			}
		}
	}
	return true
}

func (b *board) fillRandom() {
	emptyCells := b.getEmptyCells()
	randomEmpty := emptyCells[rand.Intn(len(emptyCells))]
	b.cells[randomEmpty[0]][randomEmpty[1]] = randomStartCell()
}

func (b *board) getEmptyCells() [][2]uint8 {
	emptyCells := make([][2]uint8, 0, _boardSize*_boardSize)
	for rowIdx := uint8(0); rowIdx < _boardSize; rowIdx++ {
		for colIdx := uint8(0); colIdx < _boardSize; colIdx++ {
			if b.cellIsEmpty(rowIdx, colIdx) {
				emptyCells = append(emptyCells, [2]uint8{rowIdx, colIdx})
			}
		}
	}
	return emptyCells
}

func (b *board) cellIsEmpty(row, col uint8) bool {
	return b.cells[row][col] == _emptyCell
}

func randomStartCell() uint16 {
	return [2]uint16{2, 4}[rand.Intn(2)]
}
