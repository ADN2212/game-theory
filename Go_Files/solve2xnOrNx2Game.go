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

	// 	[]int{-4, 2, 0, 3, -2},
	// []int{4, -1, 0, -3, 1},

	//fmt.Println(from2xNToNx2(g2))


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

	//Transformar el juego si es de la forma 2xn	
	var is2xn bool = is2xNGame(g)

	if is2xn {
		//Existe una relacion entre el valor de un juego (V) de la forma 2xn y su version trasformada por la funcion from2xNToNx2,
		//V(2xn) = -V(nx2), de esta manera me ahorre tener que reacer el algoritmo para juegos de la dorma 2xn. 
		g = from2xNToNx2(g)
	} else if !isNx2Game(g) {
			return gameValuePoint, fmt.Errorf(" Este juego no es de la forma '2xn' ni 'nx2'.")		
	}

	//En este caso se calculan los varores esperados desde la perspectiva de Rose,
	//y los potenciales valores del juego son parte del Upper Evelop.
	//Primero generar todas las combinaciones entre las estrategias de Rose.
	rowCombinations := combin.Combinations(len(g.Matrix), 2)
	var currentRoseAIndex int
	var currentRoseBIndex int

	//En este map las llaves son el valor de la coordenada x del punto de interseccion, 
	//y los valores son el indice de una de las estrategias de Rose
	//que contienen dicho punto:
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
	var matrizLen int = len(g.Matrix) 

	for x, roseI := range intersectionPointsMap {
		//fmt.Printf("Para Rose %v, con x = %v \n", roseI, x)
		cuurentRoseExpectedval = computeRoseExpectedVal(g.Matrix[roseI], x)
		isIn := true
		//Iterar sobre todas las estrategias para comprobar si cumple con la condicion:
		for roseJ := 0; roseJ < matrizLen; roseJ++ {			
			thisRoseExpectedVal	= computeRoseExpectedVal(g.Matrix[roseJ], x)
			//fmt.Printf("Provando interseccion entre Rose %v y Rose %v.`\n", roseI, roseJ)
			if !(thisRoseExpectedVal <= cuurentRoseExpectedval) {
				isIn = false
				break
			}
		}
		// if isIn {
		// 	fmt.Printf("%v esta en el Upper envelop.\n", x)
		// }
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
	
	if is2xn {
		gameValuePoint.y = gameValuePoint.y * (-1)
		return gameValuePoint, nil
	}

	return gameValuePoint, nil

}

//Esta funcion transforma un juego de la forma 2xN a uno de la forma Nx2
//Ademas de que cambia los signos de los valores de la matriz.  
//Asume que su inout es de la forma 2xN
func from2xNToNx2(g game.Game) (game.Game) {
	
	resGame := game.Game{
		GameName: "Converted Game",
		RowPlayer: "Rose",
		ColumnPlayer: "Colin",
		Matrix: []game.Rose{},
	}

	var firtsRowLen = len(g.Matrix[0])

	for col := 0; col < firtsRowLen; col++ {
		resGame.Matrix = append(
			resGame.Matrix, []int {
				-1 * g.Matrix[0][col],
				-1 * g.Matrix[1][col],
			},
		)
	}
	return resGame
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

	matrizLen := len(g.Matrix)

	if matrizLen != 2 {
		return false
	}

	var firtsRowLen int = len(g.Matrix[0])
	
	for i := 1; i < matrizLen; i++ {
		if len(g.Matrix[i]) != firtsRowLen {
			return false
		}
	}
	
	return true
}
