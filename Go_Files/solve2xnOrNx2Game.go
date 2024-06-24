package main

import "fmt"
import "game/game"
import "gonum.org/v1/gonum/stat/combin"

func main () {

	// g2 := game.Game{
	// 	GameName: "2xn Example",
	// 	RowPlayer: "Nina",
	// 	ColumnPlayer: "Juan",
	// 	Matrix: []game.Rose{
	// 		[]int{3, 4, 5, 10},
	// 		[]int{7, 4, 8, 12},
	// 		//[]int{7, 4, 8, 12},	
	// 	},
	// }


	g3 := game.Game{
		GameName: "nx2 Example",
		RowPlayer: "Manetha",
		ColumnPlayer: "Henoch",
		Matrix: []game.Rose{
			[]int{-3, 5},
			[]int{-1, 3},
			[]int{2, -2},
			[]int{3, -6},
		},
	}

	//g2.Show()
	//fmt.Printf("here.\n")


	gv, err := solve2xnOrnx2Game(g3, "nx2")

	if err == nil {
		fmt.Printf("El valor del juego es = %v \n", gv)
	}

}

//
func findExpectedValIntersection(roseA []int, roseB []int) (float32, error) {
	//fmt.Println(roseA)
	//fmt.Println(roseB)
	n := float32(roseB[1] - roseA[1])
	d := float32(roseA[0] - roseB[0] + roseB[1] - roseA[1])
	//Esto significa que no hay interseccion entre las ecuaciones de los valores esperados?
	if d == 0 {
		return 0.00, fmt.Errorf("Denominador igual a cero.")
	}
	return float32(n / d), nil
}

func solve2xnOrnx2Game(g game.Game, t string) (float32, error) {

		_, err := is2xnOrnx2Game(g, t)

	if err != nil {
		fmt.Println(err)
		return 0.0, err
	}

	//gameMatrix := g.Matrix

	var gameValue float32 = 0.0

	//En este caso se calculan los varores esperados desde la perspectiva de Rose.
	if t == "nx2" {
		
		//Primero generar todas las combinaciones entre las estrategias de Rose.
		rowCombinations := combin.Combinations(len(g.Matrix), 2)
		//fmt.Println(rowPermutations)
		
		var currentRoseAIndex int
		var currentRoseBIndex int
		intersectionPoints := []float32{}
		var currentErr error = nil
		var currentIntPoint float32

		for pairIndex := range rowCombinations {
			currentRoseAIndex = rowCombinations[pairIndex][0]
			currentRoseBIndex = rowCombinations[pairIndex][1]
			// fmt.Println(findExpectedValIntersection(
			// 	g.Matrix[currentRoseAIndex], 
			// 	g.Matrix[currentRoseBIndex])
			// )
			currentIntPoint, currentErr = findExpectedValIntersection(g.Matrix[currentRoseAIndex], g.Matrix[currentRoseBIndex])
			if currentErr == nil {
				intersectionPoints = append(intersectionPoints, currentIntPoint)
			}
		}
		fmt.Println(intersectionPoints)
	}
	return gameValue, nil
}


//Esta funcion comprueba si un juego es de la forma 2xn o nx2:
func is2xnOrnx2Game(g game.Game, t string) (bool, error) {

	if t != "2xn" && t != "nx2" {
		return false, fmt.Errorf("Elija entre '2xn' o 'nx2' para validar la forma del juego")
	}

	var firtsRowLen int = len(g.Matrix[0])
	matrizLen := len(g.Matrix) 

	for i := 1; i < matrizLen; i++ {
		if len(g.Matrix[i]) != firtsRowLen {
			return false, fmt.Errorf("En este juego hay filas con longitudes diferentes.")
		}
	} 

	if t == "2xn" {
		return matrizLen == 2, nil
	}

	if t == "nx2" {
		for rowIndex := range g.Matrix {
			if len(g.Matrix[rowIndex]) != 2 {
				return false, nil
			}
		}
	}

	//Originalmente este return iba dentro de la condicional para "nx2" games,
	//Pero el compilador no permite que la funcion no tenga return fuera de la raiz,
	//Sin embargo, que este aqui es logicamente equivalente.
	return true, nil

}







