package main

import (
	"fmt"
	"time"
)

type GameState struct {
	Id int
	Param chan int
}

const (
	Up = iota + 1
	Down
	Left
	Right
	Over = -1
)

var GlobalGameCtrl = make(map[int]*GameState)

func NewGameThread(id int) {
	gs := &GameState{Id : id, Param: make(chan int)}
	GlobalGameCtrl[id] = gs
	go GameServer(gs)
}

func GameServer(gs *GameState) {
	var cmd int

	for {
		select {
		case cmd = <- gs.Param:
			switch cmd {
			case Up:
				fmt.Println("game", gs.Id, "Up")
			case Down:
				fmt.Println("game", gs.Id, "Down")
			case Left:
				fmt.Println("game", gs.Id, "Left")
			case Right:
				fmt.Println("game", gs.Id, "Right")
			case Over:
				fmt.Println("game", gs.Id, "Over")
				return
			default:
				fmt.Println("game", gs.Id, "undefined")
			}
		//default:
		//	fmt.Println("game", gs.Id, "do nothing")
		//	time.Sleep(time.Second*1)
		}
	}
}

func main() {
	/*
	gameId1 := 1
	NewGameThread(gameId1)

	gsHandle := GlobalGameCtrl[gameId1]
	gsHandle.Param <- Up
	gsHandle.Param <- Down
	gsHandle.Param <- Left
	gsHandle.Param <- Down
	gsHandle.Param <- Over
	 */

	mainLoop()
}

func mainLoop() {
	loop := time.NewTicker(3 * time.Second)
	for {
		select {
		case <- loop.C:
			fmt.Println("tick...")
		//	case
		//  fmt.Println("game server finish...")
		}
	}
}
