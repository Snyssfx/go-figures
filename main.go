package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/pkg/errors"
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
	screen.Clear()
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
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func drawScreen(screen tcell.Screen, fig *figureState) {
	w, h := screen.Size()
	min := w
	if h < min {
		min = h
	}

	const gl = '▄'
	centerX, centerY := w/2, h/2
	x, y := fig.getCoords(float64(time.Now().UnixNano() / 1e6))
	realX, realY := centerX+int(float64(min)*x), centerY+int(float64(min)*y)

	screen.SetContent(realX, realY, gl, nil, tcell.StyleDefault)
	screen.Show()
}
