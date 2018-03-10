//This is a joint work of Mengxin ZHANG and Jiachen Liu
//The HTML part is written by Jiachen Liu and the remaining is finished by Mengxin ZHANG

package main

//The menu bar srtuct
type MainMenu struct{}

//The about page struct
type About struct{}

//The help page struct
type Help struct{}

//display the menu bar
func (m *MainMenu) Render() string {
	return `
<menu label = "2048">
	<menu label="Close">
		<menuitem label="Close" selector="performClose:" shortcut="meta+w" />
		<menuitem label="Quit" selector="terminate:" shortcut="meta+q" />
	</menu>
	<menu label="Options">
		<menuitem label="About" onclick="OpenAbout" shortcut="meta+a" />
		<menuitem label="Help" onclick="OpenHelp" shortcut="meta+h" />
		<menuitem label="Score board" onclick="OpenScoreBoard" shortcut="meta+s"/>
	</menu>
</menu>
	`
}

//shortcut="meta+s"
//display the about page
func (a *About) Render() string {
	return `
<div class = "Aboutheader">ABOUT</div>
<p class = "AboutP">This 2048 GUI is the part of the group project "2048" for class 02-601
    Programming for Scientists. Specifically, this GUI is made by Jiachen Liu
    and Mengxin ZHANG. We consulted the original game 2048 and added a few
    more functions to it.
</p>
<button class = "MainButton Aboutbutton" onclick = "CloseWin">CLOSE</button>
`
}

//display the about page
func (h *Help) Render() string {
	return `
<div class = "Aboutheader">HELP</div>
<p class = "AboutP">Move up or down, left or right trying to join two neighboring
    equal numbers. You will win once you get to 2048! Current we don't support
    keyboard control. Please use the directional buttons on the left of the game board.
</p>
<button class = "MainButton Aboutbutton" onclick = "CloseWin">CLOSE</button>
`
}

//Open the score board
func (m *MainMenu) OpenScoreBoard() {
	ScoreBoard.UpdateRecords()
	ScoreWindow = newWindow(497.0, 405.0, "Score Board", "#EDC22F")
	ScoreWindow.Mount(ScoreBoard)
}

//Close the ScoreRecord
func (s *ScoreRecord) CloseWin() {
	ScoreWindow.Close()
}

//Open about page
func (m *MainMenu) OpenAbout() {
	a := new(About)
	AboutWindow = newWindow(478.0, 319.0, "About", "#FFEBCD")
	AboutWindow.Mount(a)
}

//Close the About window
func (a *About) CloseWin() {
	AboutWindow.Close()
}

//Open about page
func (m *MainMenu) OpenHelp() {
	h := new(Help)
	HelpWindow = newWindow(478.0, 319.0, "Help", "#FFEBCD")
	HelpWindow.Mount(h)
}

//Close the About window
func (h *Help) CloseWin() {
	HelpWindow.Close()
}
