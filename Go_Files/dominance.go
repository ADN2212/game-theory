package main

import "fmt"
import "game/game"


func main () {	
	
	g1 := game.Game{
		GameName: "TestGame",
		RowPlayer: "Rose",
		ColumnPlayer: "Colin",
		Matrix: []game.Rose {
			[]int{10, 10, 10},	
			[]int{5, 10, 11},
			[]int{10, 10, 10},
		},		
	}

	//g1.Show()//Ejemplo de llamada al metodo.

	b, err := isDominant(&g1,"col", 0, 1)

	//Este es el equivalente a un try catch en Go:
	if err == nil {
		fmt.Printf("%v\n", b)
	} else {
		fmt.Printf("there was an error: '%v'.\n", err)
	}
}

//Esta funcion toma un juego (su puntero)
//Los indices de dos estrategias y
//El tipo de jugador para determinar si la primera domina la segunda.
func isDominant(
	gamePointer *game.Game, 
	playerType string, 
	strategyIndex1 int, 
	strategyIndex2 int) (bool, error) {

		var oneStriclyBetter = false

		if playerType == "row" {
			
			strategy1 := (*gamePointer).Matrix[strategyIndex1]
			strategy2 := (*gamePointer).Matrix[strategyIndex2]
			
			for i := range strategy1 {
				if !(strategy1[i] >= strategy2[i]) {
					return false, nil
				}
				if strategy1[i] > strategy2[i] {
					oneStriclyBetter = true
				}
			}
			return oneStriclyBetter, nil 
		}


		if playerType == "col" {

			strategy1 := []int{}			
			strategy2 := []int{}

			for _, row := range (*gamePointer).Matrix {
				strategy1 = append(strategy1, row[strategyIndex1])
				strategy2 = append(strategy2, row[strategyIndex2])
			}

			for i := range strategy1 {
				if !(strategy1[i] <= strategy2[i]) {
					return false, nil
				}
				if strategy1[i] < strategy2[i] {
					oneStriclyBetter = true
				}
			}
			return oneStriclyBetter, nil 
		}

		//En go los errores son tratados fuera de las funciones
		return false, fmt.Errorf("Opcion no valida")

}
