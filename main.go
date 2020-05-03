package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/pkg/errors"
)

const (
	frameMs     = 500
	glyph       = '▄'
	startNRunes = 45
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
	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()

	var wg sync.WaitGroup
	wg.Add(2)
	end := make(chan bool)

	fig := newFigure(0)
	hist := &history{maxLen: startNRunes}

	go pollEvent(screen, end, &wg, fig, hist)
	go redrawLoop(screen, end, &wg, fig, hist)

	screen.Show()

	wg.Wait()
	return nil
}

func pollEvent(screen tcell.Screen, end chan<- bool, wg *sync.WaitGroup, fig *figureState, hist *history) {
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
					idx := int(rune - '1')
					fig.change(idx)
				}
				if rune == '+' {
					hist.incMaxLen()
				}
				if rune == '-' {
					hist.decMaxLen()
				}
			}
		}
	}
}

func redrawLoop(screen tcell.Screen, end <-chan bool, wg *sync.WaitGroup, fig *figureState, hist *history) {
	for {
		select {

		case <-end:
			wg.Done()
			return

		default:
			drawScreen(screen, fig, hist)
			time.Sleep(frameMs)
		}
	}
}

func drawScreen(screen tcell.Screen, fig *figureState, hist *history) {
	w, h := screen.Size()
	h = 2 * h // a cell's height is twice longer than a cell's width
	min := w - 1
	if h < min {
		min = h - 1
	}

	msNow := float64(time.Now().UnixNano() / 1e6)
	t := msNow / frameMs
	points := fig.getCoords(t)
	center, radius := intPoint{w / 2, h / 2}, float64(min)/2.0

	screenPoints := []intPoint{}
	for _, p := range points {
		intP := p.toScreen(center, radius)
		screenPoints = append(screenPoints, intP)
	}
	hist.add(screenPoints)

	screen.Clear()
	for _, points := range hist.points {
		for _, p := range points {
			style := tcell.StyleDefault
			if p.y%2 == 0 { // set high half of the cell
				style = style.Reverse(true)
			}
			screen.SetContent(p.x, p.y/2, glyph, nil, style)
		}
	}
	screen.Show()
}
