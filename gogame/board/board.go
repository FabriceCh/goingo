package board

import (
	
)

//Representing the game according to official rules and terminology from wikipedia
//https://en.wikipedia.org/wiki/Go_(game)

type CrossPoint int

const(
	//see https://github.com/golang/go/wiki/Iota for what is iota
	VACANT CrossPoint = iota 
	STONE_P1 
	STONE_P2 
)


