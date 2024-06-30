package main

import "fmt"
import "game/game"
import "gonum.org/v1/gonum/stat/combin"
import "math"


type point struct {
	x float32
	y float32
}


func main () {

	//Positivo infinito como punto de partida para hallar un valor minimo
	//inf := math.Inf(1)

	g2 := game.Game{
		GameName: "2xn Example",
		RowPlayer: "Nina",
		ColumnPlayer: "Juan",
		Matrix: []game.Rose{
			[]int{-4, 2, 0, 3, -2},
			[]int{4, -1, 0, -3, 1},
		},
	}

	// g3 := game.Game{
	// 	GameName: "nx2 Example",
	// 	RowPlayer: "Manetha",
	// 	ColumnPlayer: "Henoch",
	// 	Matrix: []game.Rose{
	// 		[]int{-3, 5},
	// 		[]int{-1, 3},
	// 		[]int{2, -2},
	// 		[]int{3, -6},
	// 	},
	// }

	//g2.Show()
	//fmt.Printf("here.\n")

	gv, err := solve2xnOrnx2Game(g2)

	if err == nil {
		fmt.Printf("El valor del juego está en %v \n", gv)
	}
}

//Resuleve un jego de la forma 2xn o nx2 usando el metodo de las graficas explicado en la pagina 16.
func solve2xnOrnx2Game(g game.Game) (point, error) {

	gameValuePoint := point{
		x:0.00,
		y: float32(math.Inf(1)),
	}

	//En este caso se calculan los varores esperados desde la perspectiva de Rose,
	//y los potenciales valores del juego son parte del Upper Evelop.
	if isNx2Game(g) {
		//Primero generar todas las combinaciones entre las estrategias de Rose.
		rowCombinations := combin.Combinations(len(g.Matrix), 2)
		var currentRoseAIndex int
		var currentRoseBIndex int

		//En este map la llaves son el valor de la coordenada x del punto de interseccion, 
		//y los valores son el indice de uande las estrategias de Rose
		//que contienen dicho punto
		intersectionPointsMap := make(map[float32]int)
		var currentErr error = nil
		var currentIntPoint float32

		//Luego caluclar la intersecciones entre las ecuaciones de los valores esperados de rose:
		for pairIndex := range rowCombinations {
			currentRoseAIndex = rowCombinations[pairIndex][0]
			currentRoseBIndex = rowCombinations[pairIndex][1]
			currentIntPoint, currentErr = findExpectedValIntersection(g.Matrix[currentRoseAIndex], g.Matrix[currentRoseBIndex])
			if currentErr == nil {
				intersectionPointsMap[currentIntPoint] = currentRoseAIndex
			}
		}

		//Despues filtramos por E(Rose J) <= E(Rose A) para toda estrategia J que se intersecte con A		
		//En este slice isInUpperEvelopPint[x] = true si el punto es parte del Upper Evelop
		isInUpperEvelop := make(map[float32]bool)
		var cuurentRoseExpectedval float32;
		var thisRoseExpectedVal float32;

		for x, roseI := range intersectionPointsMap {			
			cuurentRoseExpectedval = computeRoseExpectedVal(g.Matrix[roseI], x)
			isIn := true
			for _, roseJ := range intersectionPointsMap {
				thisRoseExpectedVal	= computeRoseExpectedVal(g.Matrix[roseJ], x)
				if !(thisRoseExpectedVal <= cuurentRoseExpectedval) {
					isIn = false
					break
				}
			}
			isInUpperEvelop[x] = isIn
		}
		//Finalmente tomar el punto con menor coordenada y (Valor esperado) que sea parte del upper evelop:
		for x, isInUpperEL := range isInUpperEvelop {
			if isInUpperEL {
				roseForThisXPoint := g.Matrix[intersectionPointsMap[x]]
				expectedVal := computeRoseExpectedVal(roseForThisXPoint, x)
				if expectedVal < gameValuePoint.y {
					gameValuePoint.x = x
					gameValuePoint.y = expectedVal
				}
			}
		}
		return gameValuePoint, nil
	}

	//En este caso se buscan los puntos que esten en el lower evelop
	if is2xNGame(g) {
		//fmt.Println("... Working ...")
		gameValuePoint.y = float32(math.Inf(-1))
		columCombination := combin.Combinations(len(g.Matrix[0]), 2)		
		//fmt.Println(columCombination)
		//Siguiendo la misma logica que con Rose
		intersectionPointsMap := make(map[float32]int)
		var currentErr error = nil
		var currentXPoint float32

		for _, pairIndex := range columCombination {
			currentXPoint, currentErr = findExpectedValIntersectionForColin(g, pairIndex[0], pairIndex[1])
			//fmt.Println(currentXPoint, currentErr)
			if currentErr == nil {
				intersectionPointsMap[currentXPoint] = pairIndex[0]
			}
		}

		//fmt.Println(intersectionPointsMap)
		//Despues filtramos por E(Colin A) <= E(Colin J) para toda estrategia J que se intersecte con A		
		//pointsInLowerEvelop := []point{}		
		
		//var isInLowerEvelop bool
		//var currentColinExpectedVal float32
		//var thisColinExpectedVal float32


		// for x, colinAIndex := range intersectionPointsMap {

		// 	isInLowerEvelop = true
		// 	currentColinExpectedVal = computeColinExpectedVal(g, colinAIndex, x)

		// 	for _, colinBIndex := range intersectionPointsMap {
		// 		thisRoseExpectedVal := computeColinExpectedVal(g, colinBIndex, x)
		// 		if !(currentColinExpectedVal <= thisRoseExpectedVal) {
		// 			isInLowerEvelop = false
		// 			break
		// 		}
		// 	}

		// 	if isInLowerEvelop {
		// 		pointsInLowerEvelop = append(pointsInLowerEvelop, point{x: x, y: currentColinExpectedVal})
		// 	}

		colinDOn05 := computeColinExpectedVal(g, 3, 0.5) 
		colinEOn05 := computeColinExpectedVal(g, 4, 0.5)

		fmt.Println(colinDOn05, colinEOn05)


		
		//}

		//fmt.Println(pointsInLowerEvelop)
		//Finalmente buscar el punto con al coordenada y mas alta entre los que pertenecen al lower eveloop.
		// for _, p := range pointsInLowerEvelop {
		// 	if p.y > gameValuePoint.y {
		// 		gameValuePoint = p
		// 	}
		// }

		return gameValuePoint, nil
	}
	return gameValuePoint, fmt.Errorf(" Este juego no es de la forma '2xn' ni 'nx2'.")
}

//Calcula el valor espera de una estrategia de colin en un punto de interseccion.
func computeColinExpectedVal(g game.Game, colinIndex int, x float32) float32 {
	a := float32(g.Matrix[0][colinIndex])
	b := float32(g.Matrix[1][colinIndex])
	return a * x + (1.00 - x) * b
}

//Calcula el valor esperado de una estrategia en un punto de interseccion.
func computeRoseExpectedVal(rose []int, interSectionPoint float32) float32 {
	return float32(rose[0]) * (interSectionPoint) + float32(rose[1]) * (1.00 - interSectionPoint)
}

//Calcula el punto de interseccion entre dos estrategis de colin
func findExpectedValIntersectionForColin(g game.Game, colinIndexA int, colinIndexB int) (float32, error) {
	
	var a1 int = g.Matrix[0][colinIndexA]
	var b1 int = g.Matrix[1][colinIndexA]

	var a2 int = g.Matrix[0][colinIndexB]
	var b2 int = g.Matrix[1][colinIndexB]

	numerator := float32(b2 - b1)
	denominator := float32((a1 - a2) + (b2 - b1))

	if denominator == 0 {
		return 0.00, fmt.Errorf("Denominador igual a cero.")
	}
	
	//fmt.Printf("Para colinIndexA = %v y colinIndexB = %v. \n", colinIndexA, colinIndexB)
	//fmt.Println(numerator, denominator)


	return numerator / denominator, nil

}

//Esta funcion encuentra el punto de interseccion entre las ecuaciones de los valores esperados
//de dos estrategias de Rose de longitud dos.
func findExpectedValIntersection(roseA []int, roseB []int) (float32, error) {
	n := float32(roseB[1] - roseA[1])
	d := float32(roseA[0] - roseB[0] + roseB[1] - roseA[1])
	//Esto significa que no hay interseccion entre las ecuaciones de los valores esperados?
	if d == 0 {
		return 0.00, fmt.Errorf("Denominador igual a cero.")
	}
	return float32(n / d), nil
}

//Comprueba si la matriz de un juego es de la forma nx2
func isNx2Game(g game.Game) bool {
	for rowIndex := range g.Matrix {
		if len(g.Matrix[rowIndex]) != 2 {
			return false
		}
	}
	return true
}

//Comprueba si la matriz de un juego es de la forma 2xn
func is2xNGame(g game.Game) bool {
	var firtsRowLen int = len(g.Matrix[0])
	matrizLen := len(g.Matrix) 
	for i := 1; i < matrizLen; i++ {
		if len(g.Matrix[i]) != firtsRowLen {
			return false
		}
	}
	return true
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
