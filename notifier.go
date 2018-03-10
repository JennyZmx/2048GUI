//This is a joint work of Mengxin ZHANG and Jiachen Liu
//The HTML part is written by Jiachen Liu and the remaining is finished by Mengxin ZHANG

package main

import (
	"github.com/murlokswarm/app"
)

//The win notifier struct
type WinPage struct {
	Score int
	Best  int
}

//The loss notifier struct
type LossPage struct {
	Score int
	Best  int
}

//Display the win page
func (w *WinPage) Render() string {
	return `
<div class = "WinImage"></div>
<div class = "Winheader">Awsome! YOU WIN!</div>
<div class = "Winscore Ws1">Score: {{.Score}}</div>
<div class = "Winscore Ws2">Best: {{.Best}}</div>
<button class = "MainButton Wb1" onclick = "CloseWin">New Game</button>
<button class = "MainButton Wb2" onclick = "ShowScoreBoard">Score Board</button>
`
}

//display the loss page
func (l *LossPage) Render() string {
	return `
<div class = "LossImage"></div>
<div class = "Lossheader">Sorry! You Loss T^T </div>
<div class = "Winscore Ls1">Score: {{.Score}}</div>
<div class = "Winscore Ls2">Best: {{.Best}}</div>
<button class = "MainButton Lb Lb1" onclick = "CloseWin">New Game</button>
<button class = "MainButton Lb Lb2" onclick = "ShowScoreBoard">Score Board</button>
`
}

//Create the notifier window (need to start new game upon close)
func NewNotifierWindow(w, h float64, title, color string) app.Windower {
	return app.NewWindow(app.Window{
		Title:           title,
		TitlebarHidden:  true,
		Width:           w,
		Height:          h,
		X:               500.00,
		Y:               300.00,
		BackgroundColor: color,
		OnClose: func() bool {
			Game.Reset()
			if MainWiner != nil {
				app.Render(Game)
			}
			return true
		},
	})
}

//Will open a win notifier
func WinProcess() {
	w := new(WinPage)
	w.Score, w.Best = Game.Score, Game.Best
	WinWindow = NewNotifierWindow(458.0, 492.0, "Win", "#F67C5F")
	WinWindow.Mount(w)
}

//Close the win notifier
func (w *WinPage) CloseWin() {
	WinWindow.Close()
}

//Show the score board
func (w *WinPage) ShowScoreBoard() {
	ScoreBoard.UpdateRecords()
	ScoreWindow = newWindow(497.0, 405.0, "Score Board", "#EDC22F")
	ScoreWindow.Mount(ScoreBoard)
}

//Will open a loss notifier
func LossProcess() {
	l := new(LossPage)
	l.Score, l.Best = Game.Score, Game.Best
	LossWindow = NewNotifierWindow(438.0, 492.0, "Loss", "#696969")
	LossWindow.Mount(l)
}

//Close the loss notifier
func (l *LossPage) CloseWin() {
	LossWindow.Close()
}

//Show the score board
func (l *LossPage) ShowScoreBoard() {
	ScoreBoard.UpdateRecords()
	ScoreWindow = newWindow(497.0, 405.0, "Score Board", "#EDC22F")
	ScoreWindow.Mount(ScoreBoard)
}
