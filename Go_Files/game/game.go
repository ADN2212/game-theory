package game;

import "fmt"

type Rose []int
//Un juego hecho desde la perspectiva de Rose:
type Game struct {
	GameName string
	RowPlayer string
	ColumnPlayer string
	Matrix []Rose
}

// func ShowGame(gamePointer *Game) {
// 	fmt.Printf("Mostrando juego %v\n", (*gamePointer).GameName)
// 	fmt.Printf("Jugador en filas: %v, Jugador en columnas: %v\n", (*gamePointer).RowPlayer, (*gamePointer).ColumnPlayer)	
// 	fmt.Print("Matriz del juego:\n")
// 	for _, row := range (*gamePointer).Matrix {
// 		fmt.Printf("%v\n", row)
// 	}
// }

//Como en procipio los juegos no seran tan grandes este metodo puede funcionar con una copia del juego original:
func (game Game) Show() {
	fmt.Printf("Mostrando juego %v\n", game.GameName)
	fmt.Printf("Jugador en filas: %v, Jugador en columnas: %v\n", game.RowPlayer, game.ColumnPlayer)	
	fmt.Print("Matriz del juego:\n")
	for _, row := range game.Matrix {
		fmt.Printf("%v\n", row)
	}
}
