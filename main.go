package main

import (
	"fmt"
	// "log"
	// "os"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/pkg/errors"
)

const (
	frameMs = 500
	glyph   = '▄'
)

func main() {
	err := initAndDraw()
	if err != nil {
		fmt.Print(err)
	}
}

func initAndDraw() error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		return errors.Wrap(err, "NewScreen()")
	}

	encoding.Register()
	if err := screen.Init(); err != nil {
		return errors.Wrap(err, "screen.Init()")
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()

	var wg sync.WaitGroup
	wg.Add(2)
	end := make(chan bool)
	fig := figureState{coordsFunc: circle}

	go pollEvent(screen, end, &wg, &fig)
	go redrawLoop(screen, end, &wg, &fig)

	screen.Show()

	wg.Wait()
	return nil
}

func pollEvent(screen tcell.Screen, end chan<- bool, wg *sync.WaitGroup, fig *figureState) {
	for {
		event := screen.PollEvent()
		switch event := event.(type) {

		case *tcell.EventKey:
			switch event.Key() {

			case tcell.KeyEscape, tcell.KeyEnter:
				close(end)
				wg.Done()
				return

			case tcell.KeyRune:
				rune := event.Rune()
				if rune == 'q' {
					close(end)
					wg.Done()
					return
				}
				if '1' <= rune && rune <= '9' {
					// rune - '1'

				}
			}
		}
	}
}

func redrawLoop(screen tcell.Screen, end <-chan bool, wg *sync.WaitGroup, fig *figureState) {
	for {
		select {

		case <-end:
			wg.Done()
			return

		default:
			drawScreen(screen, fig)
			time.Sleep(frameMs)
		}
	}
}

func drawScreen(screen tcell.Screen, fig *figureState) {
	w, h := screen.Size()
	h = 2 * h // a cell's height is twice longer than a cell's width
	min := w - 1
	if h < min {
		min = h - 1
	}
	radius := float64(min) / 2.0

	centerX, centerY := w/2, h/2
	msNow := float64(time.Now().UnixNano() / 1e6)
	x, y := fig.getCoords(msNow / frameMs)
	realX, realY := centerX+int(radius*x), centerY+int(radius*y)

	style := tcell.StyleDefault
	if realY%2 == 0 { // set high half of the cell
		style = style.Reverse(true)
	}

	screen.SetContent(realX, realY/2, glyph, nil, style)
	screen.Show()
}
