package main

import (
	"github.com/murlokswarm/app"
	"math/rand"
	"time"
)

//Declear all windowers and components here as globle variables for easy manipulate
var (
	MainWiner   app.Windower //Main window
	ScoreWindow app.Windower //Score board window
	AboutWindow app.Windower //About Page
	HelpWindow  app.Windower //Help Page
	WinWindow   app.Windower //Win notifier
	LossWindow  app.Windower //Loss notifier
	ScoreBoard  *ScoreRecord //The ScoreBoard component
	Game        *Game2048    //The game component
)

//all component need to be registered before use
func init() {
	app.RegisterComponent(&Game2048{})
	app.RegisterComponent(&MainMenu{})
	app.RegisterComponent(&ScoreRecord{})
	app.RegisterComponent(&About{})
	app.RegisterComponent(&Help{})
	app.RegisterComponent(&WinPage{})
	app.RegisterComponent(&LossPage{})
}

//This function will create a new main window
//Due to restriction of OnClose, this function cannot be as a Subroutine
func newMainWindow() app.Windower {
	return app.NewWindow(app.Window{
		Title:           "2048",
		TitlebarHidden:  true,
		Width:           838.0,
		Height:          638.0,
		X:               500.00,
		Y:               300.00,
		BackgroundColor: "#FAF8EF",
		OnClose: func() bool {
			MainWiner = nil
			return true
		},
	})
}

//This function will create all windows except for the main window and notifiers
func newWindow(w, h float64, title, color string) app.Windower {
	return app.NewWindow(app.Window{
		Title:           title,
		TitlebarHidden:  true,
		Width:           w,
		Height:          h,
		X:               500.00,
		Y:               300.00,
		BackgroundColor: color,
	})
}

//do all initialization of the game
func InitializeGame(g *Game2048, s *ScoreRecord) {
	rand.Seed(time.Now().UTC().UnixNano())
	g.Reset()
	s.ReadRecordFile()
	g.Best = s.Score1
}

func main() {
	//initialization
	Game, ScoreBoard = new(Game2048), new(ScoreRecord)
	InitializeGame(Game, ScoreBoard)
	//On launch of the app
	app.OnLaunch = func() {
		//Open a window
		MainWiner = newMainWindow()
		MainWiner.Mount(Game)
		//Create a menu bar
		AppMenu := new(MainMenu)
		if m, ok := app.MenuBar(); ok {
			m.Mount(AppMenu)
		}
	}
	//On reopen of the app will open the main window if the main window does not exist
	app.OnReopen = func() {
		if MainWiner != nil {
			return
		}
		MainWiner = newMainWindow()
		MainWiner.Mount(Game)
	}
	//When close the game, write the record to the file
	app.OnFinalize = func() {
		ScoreBoard.WriteRecord()
	}
	//Run the app!
	app.Run()
}
