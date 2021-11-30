package game

import (
	"reflect"
	"testing"
)

const a = 35 << 1

func TestBoard_shift(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	hasChanged := b.shiftDown()
	equal(t, [4]uint16{2, 0, 0, 0}, b.cells[0])
	equal(t, [4]uint16{4, 0, 0, 0}, b.cells[1])
	equal(t, [4]uint16{8, 2, 16, 0}, b.cells[2])
	equal(t, [4]uint16{4, 4, 16, 2}, b.cells[3])
	equal(t, true, hasChanged)

	hasChanged = b.shiftRight()
	equal(t, [4]uint16{0, 0, 0, 2}, b.cells[0])
	equal(t, [4]uint16{0, 0, 0, 4}, b.cells[1])
	equal(t, [4]uint16{0, 8, 2, 16}, b.cells[2])
	equal(t, [4]uint16{0, 8, 16, 2}, b.cells[3])
	equal(t, true, hasChanged)

	hasChanged = b.shiftUp()
	equal(t, [4]uint16{0, 16, 2, 2}, b.cells[0])
	equal(t, [4]uint16{0, 0, 16, 4}, b.cells[1])
	equal(t, [4]uint16{0, 0, 0, 16}, b.cells[2])
	equal(t, [4]uint16{0, 0, 0, 2}, b.cells[3])
	equal(t, true, hasChanged)

	hasChanged = b.shiftLeft()
	equal(t, [4]uint16{16, 4, 0, 0}, b.cells[0])
	equal(t, [4]uint16{16, 4, 0, 0}, b.cells[1])
	equal(t, [4]uint16{16, 0, 0, 0}, b.cells[2])
	equal(t, [4]uint16{2, 0, 0, 0}, b.cells[3])
	equal(t, true, hasChanged)
}

func TestBoard_won(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1024, 0, 0, 0},
		{1024, 0, 0, 0},
	}
	equal(t, true, b.shiftDown())
	equal(t, [4]uint16{0, 0, 0, 0}, b.cells[0])
	equal(t, [4]uint16{0, 0, 0, 0}, b.cells[1])
	equal(t, [4]uint16{0, 0, 0, 0}, b.cells[2])
	equal(t, [4]uint16{2048, 0, 0, 0}, b.cells[3])
	equal(t, true, b.won)
	equal(t, uint32(2048), b.score)
}

func TestBoard_noMovesRemaining(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}
	equal(t, true, b.noMovesRemaining())

	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	equal(t, false, b.noMovesRemaining())
}

func TestBoard_shiftUp(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	equal(t, true, b.shiftUp())
	equal(t, [4]uint16{2, 4, 16, 2}, b.cells[0])
	equal(t, [4]uint16{4, 2, 16, 0}, b.cells[1])
	equal(t, [4]uint16{8, 0, 0, 0}, b.cells[2])
	equal(t, [4]uint16{4, 0, 0, 0}, b.cells[3])
	equal(t, true, b.shiftUp())
	equal(t, [4]uint16{2, 4, 32, 2}, b.cells[0])
	equal(t, [4]uint16{4, 2, 0, 0}, b.cells[1])
	equal(t, [4]uint16{8, 0, 0, 0}, b.cells[2])
	equal(t, [4]uint16{4, 0, 0, 0}, b.cells[3])
	equal(t, false, b.shiftUp())
}

func TestBoard_shiftDown(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	equal(t, true, b.shiftDown())
	equal(t, [4]uint16{2, 0, 0, 0}, b.cells[0])
	equal(t, [4]uint16{4, 0, 0, 0}, b.cells[1])
	equal(t, [4]uint16{8, 2, 16, 0}, b.cells[2])
	equal(t, [4]uint16{4, 4, 16, 2}, b.cells[3])
	equal(t, true, b.shiftDown())
	equal(t, [4]uint16{2, 0, 0, 0}, b.cells[0])
	equal(t, [4]uint16{4, 0, 0, 0}, b.cells[1])
	equal(t, [4]uint16{8, 2, 0, 0}, b.cells[2])
	equal(t, [4]uint16{4, 4, 32, 2}, b.cells[3])
	equal(t, false, b.shiftDown())
}

func TestBoard_shiftLeft(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	equal(t, true, b.shiftLeft())
	equal(t, [4]uint16{4, 8, 0, 0}, b.cells[0])
	equal(t, [4]uint16{4, 2, 8, 0}, b.cells[1])
	equal(t, [4]uint16{16, 2, 0, 0}, b.cells[2])
	equal(t, [4]uint16{4, 2, 8, 0}, b.cells[3])
	equal(t, false, b.shiftLeft())
	equal(t, [4]uint16{4, 8, 0, 0}, b.cells[0])
	equal(t, [4]uint16{4, 2, 8, 0}, b.cells[1])
	equal(t, [4]uint16{16, 2, 0, 0}, b.cells[2])
	equal(t, [4]uint16{4, 2, 8, 0}, b.cells[3])
}

func TestBoard_shiftRight(t *testing.T) {
	b := initNewBoard()
	b.cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	equal(t, true, b.shiftRight())
	equal(t, [4]uint16{0, 0, 4, 8}, b.cells[0])
	equal(t, [4]uint16{0, 4, 2, 8}, b.cells[1])
	equal(t, [4]uint16{0, 0, 16, 2}, b.cells[2])
	equal(t, [4]uint16{0, 4, 2, 8}, b.cells[3])
	equal(t, false, b.shiftRight())
	equal(t, [4]uint16{0, 0, 4, 8}, b.cells[0])
	equal(t, [4]uint16{0, 4, 2, 8}, b.cells[1])
	equal(t, [4]uint16{0, 0, 16, 2}, b.cells[2])
	equal(t, [4]uint16{0, 4, 2, 8}, b.cells[3])
}

func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkBoard_Shift(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.Shift(DirectionDown)
	}
}

func BenchmarkBoard_noMovesRemaining(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.noMovesRemaining()
	}
}

func BenchmarkBoard_shiftLeft(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.shiftLeft()
	}
}

func BenchmarkBoard_shiftRight(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.shiftRight()
	}
}

func BenchmarkBoard_shiftUp(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.shiftUp()
	}
}

func BenchmarkBoard_shiftDown(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.shiftDown()
	}
}

func BenchmarkBoard_fillRandom(b *testing.B) {
	board := initNewBoard()
	var cells = [4][4]uint16{
		{2, 2, 8, 0},
		{4, 2, 8, 0},
		{8, 0, 8, 2},
		{4, 2, 8, 0},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.cells = cells
		board.fillRandom()
	}
}

func BenchmarkBoard_randomStartCell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randomStartCell()
	}
}
