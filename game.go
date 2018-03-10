package main

import (
	"github.com/murlokswarm/app"
)

//A 2048 object: a. the Gameboard as 2D slice. b. total Score. c. best score
//Render does not support slice, so we need to assign each cell number to individual fields
//CellClass are also assigned to individual fields for better visualization (change of class)
type Game2048 struct {
	Gameboard [][]int
	Score     int
	Best      int
}

//Display the Gameboard(only non-0 tiles), current score and best score by divs
//add buttons that can triger function upon click (new game and direction button)
//This game does not support keyboard control (explained in writeup)
func (g *Game2048) Render() string {
	return `
<button class = "MainButton NewGame" onclick = "NewGame">NewGame</button>
<button class = "MainButton Direction UpButton" onclick = "MoveUp"> &#9650; </button>
<button class = "MainButton Direction DownButton" onclick = "MoveDown"> &#9660; </button>
<button class = "MainButton Direction LeftButton" onclick = "MoveLeft"> &#9668; </button>
<button class = "MainButton Direction RightButton" onclick = "MoveRight"> &#9658; </button>
<div class = "MainHeader">2048</div>
<div class = "scoretitle main1">SCORE</div>
<div class = "scoretitle main2">BEST</div>
<div class = "MainP">Join the numbers and get to the <strong>2048 tile!</strong></div>
<div class = "boardbackground"></div>
<div class = "cell cell11 cell_{{index .Gameboard 0 0}}"><span>{{if index .Gameboard 0 0}}{{index .Gameboard 0 0}}{{end}}</span></div>
<div class = "cell cell12 cell_{{index .Gameboard 0 1}}"><span>{{if index .Gameboard 0 1}}{{index .Gameboard 0 1}}{{end}}</span></div>
<div class = "cell cell13 cell_{{index .Gameboard 0 2}}"><span>{{if index .Gameboard 0 2}}{{index .Gameboard 0 2}}{{end}}</span></div>
<div class = "cell cell14 cell_{{index .Gameboard 0 3}}"><span>{{if index .Gameboard 0 3}}{{index .Gameboard 0 3}}{{end}}</span></div>
<div class = "cell cell21 cell_{{index .Gameboard 1 0}}"><span>{{if index .Gameboard 1 0}}{{index .Gameboard 1 0}}{{end}}</span></div>
<div class = "cell cell22 cell_{{index .Gameboard 1 1}}"><span>{{if index .Gameboard 1 1}}{{index .Gameboard 1 1}}{{end}}</span></div>
<div class = "cell cell23 cell_{{index .Gameboard 1 2}}"><span>{{if index .Gameboard 1 2}}{{index .Gameboard 1 2}}{{end}}</span></div>
<div class = "cell cell24 cell_{{index .Gameboard 1 3}}"><span>{{if index .Gameboard 1 3}}{{index .Gameboard 1 3}}{{end}}</span></div>
<div class = "cell cell31 cell_{{index .Gameboard 2 0}}"><span>{{if index .Gameboard 2 0}}{{index .Gameboard 2 0}}{{end}}</span></div>
<div class = "cell cell32 cell_{{index .Gameboard 2 1}}"><span>{{if index .Gameboard 2 1}}{{index .Gameboard 2 1}}{{end}}</span></div>
<div class = "cell cell33 cell_{{index .Gameboard 2 2}}"><span>{{if index .Gameboard 2 2}}{{index .Gameboard 2 2}}{{end}}</span></div>
<div class = "cell cell34 cell_{{index .Gameboard 2 3}}"><span>{{if index .Gameboard 2 3}}{{index .Gameboard 2 3}}{{end}}</span></div>
<div class = "cell cell41 cell_{{index .Gameboard 3 0}}"><span>{{if index .Gameboard 3 0}}{{index .Gameboard 3 0}}{{end}}</span></div>
<div class = "cell cell42 cell_{{index .Gameboard 3 1}}"><span>{{if index .Gameboard 3 1}}{{index .Gameboard 3 1}}{{end}}</span></div>
<div class = "cell cell43 cell_{{index .Gameboard 3 2}}"><span>{{if index .Gameboard 3 2}}{{index .Gameboard 3 2}}{{end}}</span></div>
<div class = "cell cell44 cell_{{index .Gameboard 3 3}}"><span>{{if index .Gameboard 3 3}}{{index .Gameboard 3 3}}{{end}}</span></div>
<div class = "MainScore MainBestScore">{{ .Best}}</div>
<div class = "MainScore MainCurrentScore">{{ .Score}}</div>
	`
}

//Restart the game. Called upon clicking "New Game" button
//record the score every time before a new game
func (g *Game2048) NewGame() {
	ScoreBoard.AddRecord(g.Score)
	g.Reset()
	app.Render(g)
}

//Move Up the board. Called upon clicking "Up" button
//con records whether to continue the game
func (g *Game2048) MoveUp() {
	Con := g.Step(0)
	//Update the best score
	if g.Best < g.Score {
		g.Best = g.Score
	}
	app.Render(g)
	//if the game can no longer continue, call stop function
	if Con == false {
		g.StopGame()
	}
}

//Move down. Called upon clicking "Down" button
func (g *Game2048) MoveDown() {
	Con := g.Step(1)
	//Update the best score
	if g.Best < g.Score {
		g.Best = g.Score
	}
	app.Render(g)
	if Con == false {
		g.StopGame()
	}
}

//Move left. Called upon clicking "Left" button
func (g *Game2048) MoveLeft() {
	Con := g.Step(2)
	//Update the best score
	if g.Best < g.Score {
		g.Best = g.Score
	}
	app.Render(g)
	if Con == false {
		g.StopGame()
	}
}

//Move right. Called upon clicking "Right" button
func (g *Game2048) MoveRight() {
	Con := g.Step(3)
	//Update the best score
	if g.Best < g.Score {
		g.Best = g.Score
	}
	app.Render(g)
	if Con == false {
		g.StopGame()
	}
}

//Will stop the function and open notifiers (in notifier.go)
func (g *Game2048) StopGame() {
	ScoreBoard.AddRecord(g.Score)
	for i := range g.Gameboard {
		for j := range g.Gameboard[i] {
			if g.Gameboard[i][j] == 2048 { //Get to 2048
				WinProcess()
				return
			}
		}
	}
	LossProcess()
}
