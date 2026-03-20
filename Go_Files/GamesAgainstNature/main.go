package main

import "fmt"
import "math"

const maxFloat = math.MaxFloat32
const minFloat = -maxFloat

type game struct {
	rowName string
	colName string
	rowStrategies []string
	matrix  [][]int
}

type result struct {
	val      float64
	strategy string
}

type methodsResults struct {
	laplace result
	wald    result
	hurwicz result
	savage  result
}


func main() {
	//ver pagina 59
	gameOneCap10 := game{
		rowName: "You",
		colName: "Nature",
		rowStrategies: []string{"A", "B", "C", "D"},
		matrix: [][]int{
			{2,2,0,1},
			{1,1,1,1},
			{0,4,0,0},
			{1,3,0,0},
		},
	}
	
	//ver pagina 61
	gameOfExerciseOneCap10 := game{
		rowName: "Manager",
		colName: "Economy",
		rowStrategies: []string{"Hold steady", "Expand Slightly", "Expnad greatly", "Diversify"},
		matrix: [][]int{
			{3,2,2,0},
			{4,2,0,0},
			{6,2,0,-2},
			{1,1,2,2},
		},
	}

	fmt.Println(gameOneCap10.computeRecomendations())
	fmt.Println(gameOfExerciseOneCap10.computeRecomendations())

}

//Este metodo calcula la estrategia recomendada al jugar un juego en contra de la naturaleza
//ver cap 10
func (g *game) computeRecomendations() methodsResults {
	//Laplace: choose the row with the higest average entry or the higest sum
	laplaceResult := result{val: minFloat, strategy: ""}
	for rowIndex := range g.matrix {
		curr := g.rowAvg(rowIndex)
		if curr > laplaceResult.val {
			laplaceResult.val = curr
			laplaceResult.strategy = g.rowStrategies[rowIndex]
		}
	}

	//Wald: Write down the minimum entry in each row. choose the row wiht the largest minumun
	waldResult := result{val: minFloat, strategy: ""}
	for rowIndex := range g.matrix {
		curr := g.rowMin(rowIndex)
		if curr > waldResult.val {
			waldResult.val = curr
			waldResult.strategy = g.rowStrategies[rowIndex]
		}
	}

	//Hurwics: choose a "coefficient of optimism" alpha between 0 and 1. For each row, compute:
	//let max = max entry in row, let min = aja
	// alpha * max + (1 - alpha) min
	//then choose the row wiht the higest result
	hurwicsResult := result{val: minFloat, strategy: ""}
	alpha := 0.75//Esto es subjetivo, pero uso el 3/4 del ejemplo.
	for rowIndex := range g.matrix {
		curr := g.rowWeightedAverage(alpha, rowIndex)
		if curr > hurwicsResult.val {
			hurwicsResult.val = curr
			hurwicsResult.strategy = g.rowStrategies[rowIndex]
		}
	}

	//Savage: For the regret matrix, writte down the largest entry in each row,
	//then choose the row for which this largets entry is smaller
	//La idea de este metodo es hallar el minimo de los maximos arrepentimientos
	savageResutl := result{val: maxFloat, strategy: ""}
	regretMatix := g.computeRegretMatrix()
	placeholdeGame := game{rowName: "Foo", colName: "Bar", matrix: regretMatix}

	for rowIndex := range placeholdeGame.matrix {
		curr := placeholdeGame.rowMax(rowIndex)
		if curr < savageResutl.val {
			savageResutl.val = curr
			savageResutl.strategy = g.rowStrategies[rowIndex]
		}
	}	

	return methodsResults{
		laplace: laplaceResult,
		wald:    waldResult,
		hurwicz: hurwicsResult,
		savage:  savageResutl,
	}

}

func (g *game) rowAvg(rowIndex int) float64 {
	sum := float64(0.0)
	for _, val := range g.matrix[rowIndex] {
		sum = sum + float64(val)
	}
	return sum / float64(len(g.matrix[rowIndex]))
}

func (g *game) rowMin(rowIndex int) float64 {
	min := maxFloat
	for _, val := range g.matrix[rowIndex] {
		floatVal := float64(val)
		if floatVal < min {
			min = floatVal
		}
	}
	return min
}

func (g *game) rowMax(rowIndex int) float64 {
	max := minFloat
	for _, val := range g.matrix[rowIndex] {
		floatVal := float64(val)
		if floatVal > max {
			max = floatVal
		}
	}
	return max
}

func (g *game) computeRegretMatrix() [][]int {

	rows := len( g.matrix)
	cols := len(g.matrix[0])

	regretMatrix := make([][]int, rows)
	
	for i := range regretMatrix {
    	regretMatrix[i] = make([]int, cols)	
	}

	//copy each entry of the original matrix in the regret matirx:
	for i, row := range g.matrix {
		copy(regretMatrix[i], row)
	}

	//compute colums maxs:
	//Aqui colsMax[i] = val, significa que val es el major valor de la columna con indice i 
	colsMax := make([]int, rows)
	for i := range colsMax {
		colsMax[i] = math.MinInt
	}

	for _, row := range regretMatrix {
		for j, entry := range row {
			if entry > colsMax[j] {
				colsMax[j] = entry
			}
		}
	}

	//In the regret matrix every entry is equal to MaxValInCoulm - entry see page 58
	for i, row := range regretMatrix {
		for j, entry := range row {
			regretMatrix[i][j] = colsMax[j] - entry  
		}
	}

	return regretMatrix

}

func (g *game) rowWeightedAverage(optimismCoefficient float64, rowIndex int) float64 {
	rowMax := g.rowMax(rowIndex)
	rowMin := g.rowMin(rowIndex)
	return optimismCoefficient * rowMax + (1 - optimismCoefficient) * rowMin
}

