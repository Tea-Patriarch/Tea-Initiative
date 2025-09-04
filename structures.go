package main

type chara struct {
	name       string
	initiative int
	isVating   bool
}

type model struct {
	//characters and cursor
	characters []chara
	cursor     int

	// size of the window
	width  int
	height int

	// Input
	inputMode   bool
	inputBuffer string
}
