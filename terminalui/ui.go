package terminalui

import (
	"log"
	"strconv"

	"github.com/brandenc40/2048/game"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	// game dimensions
	width      = 60
	height     = width / 3
	x, y       = 2, 1
	xGap, yGap = 3, 2

	// game cells
	borderXStart = x
	borderXEnd   = borderXStart + width
	borderYStart = y
	borderYEnd   = borderYStart + height
	cellsXStart  = borderXStart + 2
	cellsYStart  = borderYStart + 1
	cellXGap     = xGap
	cellYGap     = yGap
	cellWidth    = (borderXEnd - borderXStart - cellXGap*3) / 4
	cellHeight   = (borderYEnd - borderYStart - cellYGap*3) / 4

	// score
	scoreX = borderXStart + 2
	scoreY = borderYEnd + 2
)

var textMsg = [...]string{
	"HOW TO PLAY: Use your arrow keys to move the",
	"tiles. Tiles with the same number merge into",
	"one when they touch. Add them up to reach 2048!",
	"",
	"Reset the game with 'R' or 'r'",
	"",
	"Quit with ESC or CTRL+C",
}

var logo = [...]string{
	"  .-----.   .----.      .---.    .-----.      ",
	" / ,-.   \\ /  ..  \\    / .  |   /  .-.  \\  ",
	" '-'  |  |.  /  \\  .  / /|  |  |   \\_.' /   ",
	"   . '  / |  |  '  | / / |  |_  /  .-. '.     ",
	"  .'  /__ '  \\  /  '/  '-'    ||  |   |  |   ",
	" |       | \\  `'  / `----|  |-' \\  '-'  /   ",
	"  -------'  `---''       `--'    `----''      ",
	"                                              ",
}

type ui struct {
	isOver      bool
	gc          game.Controller
	colorPalate colorPalate
}

// Run -
func Run(gc game.Controller, options ...Option) {
	u := &ui{
		gc:          gc,
		colorPalate: normalPalate(),
	}
	closeFunc := u.initialize(options...)
	defer closeFunc()

	u.drawGameBoard()
	u.runGameLoop()
}

func (u *ui) initialize(options ...Option) (closeFunc func()) {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	for _, option := range options {
		option.apply(u)
	}
	closeFunc = termbox.Close
	return
}

func (u *ui) drawGameBoard() {
	u.drawGameBackground()
	u.drawGameCells()
	u.drawScore()
	u.drawGuide()
	if err := termbox.Flush(); err != nil {
		log.Fatal(err)
	}
}

func (u *ui) shiftGameController(direction game.Direction) {
	if u.isOver {
		return
	}
	u.gc.Shift(direction)
	u.drawGameCells()
	u.drawScore()
	if u.gc.Won() {
		u.drawOverlayMessage("YOU WIN!")
		u.isOver = true
	}
	if u.gc.Lost() {
		u.drawOverlayMessage("NO MORE MOVES, TRY AGAIN")
		u.isOver = true
	}
	if err := termbox.Flush(); err != nil {
		log.Fatal(err)
	}
}
func (u *ui) resetGameBoard() {
	u.gc.Reset()
	u.isOver = false
	u.drawGameBoard()
}

func (u *ui) runGameLoop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				u.shiftGameController(game.DirectionUp)
			case termbox.KeyArrowDown:
				u.shiftGameController(game.DirectionDown)
			case termbox.KeyArrowRight:
				u.shiftGameController(game.DirectionRight)
			case termbox.KeyArrowLeft:
				u.shiftGameController(game.DirectionLeft)
			case termbox.KeyCtrlC, termbox.KeyEsc:
				return
			default:
				switch ev.Ch {
				case 'r', 'R':
					u.resetGameBoard()
				}
			}
		case termbox.EventResize:
			u.drawGameBoard()
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
	}
}

func (u *ui) drawGameBackground() {
	for x := borderXStart; x <= borderXEnd; x++ {
		for y := borderYStart; y <= borderYEnd; y++ {
			termbox.SetCell(x, y, ' ', u.colorPalate.border, u.colorPalate.border)
		}
	}
}

func (u *ui) drawGameCells() {
	for rowIdx, row := range u.gc.GetCells() {
		for colIdx, col := range row {
			u.drawGameCell(colIdx, rowIdx, col)
		}
	}
}

func (u *ui) drawGameCell(colIdx, rowIdx int, value uint16) {
	var (
		bg     = u.colorPalate.empty
		xStart = cellsXStart + (colIdx * cellWidth) + (colIdx * cellXGap)
		xEnd   = xStart + cellWidth
		yStart = cellsYStart + (rowIdx * cellHeight) + (rowIdx * cellYGap)
		yEnd   = yStart + cellHeight
		xMid   int
		yMid   int
	)
	if value > 0 {
		bg = u.colorPalate.values[value]
		xMid = (xStart + xEnd) / 2
		yMid = (yStart + yEnd) / 2
		if value > 100 {
			xMid -= 1
		}
	}
	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			termbox.SetCell(x, y, ' ', termbox.ColorWhite, bg)
		}
	}
	if value > 0 {
		val := strconv.FormatUint(uint64(value), 10)
		tbPrint(xMid, yMid, u.colorPalate.valueText, bg, val)
	}
}

func (u *ui) drawScore() {
	msg := "Current Score: " + strconv.Itoa((int)(u.gc.GetScore()))
	tbPrint(scoreX, scoreY, u.colorPalate.score, termbox.ColorDefault, msg)
}

func (u *ui) drawGuide() {
	y := borderYStart
	x := borderXEnd + 2
	for _, line := range logo {
		y++
		tbPrint(x, y, u.colorPalate.guide, termbox.ColorDefault, line)
	}
	for _, line := range textMsg {
		y++
		tbPrint(x, y, u.colorPalate.guide, termbox.ColorDefault, line)
	}
}

func (u *ui) drawOverlayMessage(message string) {
	x := borderXStart + (width-len(message))/2
	y := borderYStart + height/2
	tbPrint(x, y, u.colorPalate.overlayText, u.colorPalate.overlayBg, message)
}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
