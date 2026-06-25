package main

import (
	"fmt"
)

// En un non-zero-sum game cada casilla de la matiz del juego es un par ordenado con el outcome de cada jugador
type OutcomePair struct {
	Rose  int
	Colin int
}

// again, una strategia no es mas que un arreglog de outcomes
type Strategy []OutcomePair

// y un juego no es mas que un arreglo de
type Game []Strategy

// This is O(n^2) donde n es igual al numero de outcomes en el juego.
// Tambien puede verse coo O((R * C )^2) siendo r y c la cantidad de estrategias de Rose y Colin respectivamente
// Note that, this is just a burte force algorithm
// Este algoritmo se puede mejorar ????
func (g *Game) findParetoOptimalOutComes() []OutcomePair {
	//En pagina 68 de define un Pareto Optimal outcome como uno tal que,
	//En todo el juego no hay otro outcome que le de a ambos jugadores payoffs mayores o que,
	//le de a uno de los jugadores un paoff igual y al otro uno mayor.
	var allGameOutcomes []OutcomePair
	//Primero tener todos los outcome del juego en un simple array
	for _, s := range *g {
		for _, o := range s {
			allGameOutcomes = append(allGameOutcomes, o)
		}
	}
	var paretoOptimalOutcomes []OutcomePair
	//Luego, por cada outcome revisamos que compla con las codiciones para ser Pareto Optimal
	for _, currentPair := range allGameOutcomes {
		//Asumimos que todo par es pareto optimal desde el pricipio
		currentPairIsParetoOptimal := true
		for _, pair := range allGameOutcomes {
			//See non-Pareto-Optimal definition in page 68
			if (pair.Rose > currentPair.Rose && pair.Colin > currentPair.Colin) || (pair.Rose == currentPair.Rose && pair.Colin > currentPair.Colin || pair.Colin == currentPair.Colin && pair.Rose > currentPair.Rose) {
				currentPairIsParetoOptimal = false
				break //Esto ahorra interasiones, desde que un par rompe la regla no es necesario seguir testeando los demas.
			}
		}
		if currentPairIsParetoOptimal {
			paretoOptimalOutcomes = append(paretoOptimalOutcomes, currentPair)
		}
	}
	return paretoOptimalOutcomes
}

// Esta funcion resive outcome pair y determina si es Pareto Optimal
// Su time complexity es O(n) donde n es la cantidad de outcomes del juego,
// siendo n = r * c,  then O(n) = O(r * c) osea,
// que la time complexity tambien esta determinada por el producto de las filas y columnas del juego
func (g *Game) isParetoOptimal(pair OutcomePair) bool {
	for _, s := range *g {
		for _, gamePair := range s {
			if (gamePair.Rose > pair.Rose && gamePair.Colin > pair.Colin) || (gamePair.Rose == pair.Rose && gamePair.Colin > pair.Colin || gamePair.Colin == pair.Colin && gamePair.Rose > pair.Rose) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Finding Pareto Optimal Pairs: ")
	fmt.Println("For Game 11.4")
	//Notece como en este juego, el Nash Equilibrium (0, 0) no es non-Pareto-Optimal
	//Esto deja demostrado que el hecho de que un outcome sea un Equilibrio de Nash no hace que este sea el mejor para ambos jugadores
	//Es decir, "quedar empate" puede no ser la mejor optcion.
	gamme11Dot4 := Game{
		{{3, 3}, {-1, 5}},
		{{5, -1}, {0, 0}},
	}
	for _, po := range gamme11Dot4.findParetoOptimalOutComes() {
		fmt.Println(po)
	}

	fmt.Println("For Game 11.3")
	//Notece como en este juego los dos Nash Equilibrium son Pareto-Optimal
	gamme11Dot3 := Game{
		{{1, 1}, {2, 5}},
		{{5, 2}, {-1, -1}},
	}

	for _, po := range gamme11Dot3.findParetoOptimalOutComes() {
		fmt.Println(po)
	}

	fmt.Println("For Game 11.5")
	//Notece como en este juego los dos Nash Equilibrium son Pareto-Optimal
	gamme11Dot5 := Game{
		{{0, -1}, {0, 2}, {2, 3}},
		{{0, 0}, {2, 1}, {1, -1}},
		{{2, 2}, {1, 4}, {1, -1}},
	}

	for _, po := range gamme11Dot5.findParetoOptimalOutComes() {
		fmt.Println(po)
	}

	fmt.Println("For Game C in page 72")
	gameConPage72 := Game{
		{{3, 0}, {5, 2}, {0, 5}},
		{{2, 2}, {1, 1}, {3, 3}},
		{{4, 1}, {4, 0}, {1, 0}},
	}

	for _, po := range gameConPage72.findParetoOptimalOutComes() {
		fmt.Println(po)
	}

	fmt.Println("For Last Game in cap 11")
	lastGameOfCap11 := Game{
		{{0, 1}, {0, 1}, {2, 4}},
		{{5, 1}, {4, 2}, {1, 0}},
		{{4, 3}, {1, 4}, {1, 0}},
	}
	for _, po := range lastGameOfCap11.findParetoOptimalOutComes() {
		fmt.Println(po)
	}

	fmt.Println("Provando si un punto es Pareto Optimal dentro de un juego: ")
	test1 := OutcomePair{3,3}
	if gamme11Dot3.isParetoOptimal(test1) {
		fmt.Println(test1, "is pareto optimal in Game 11.3")
	}

	test2 := OutcomePair{2,5}
	if gamme11Dot5.isParetoOptimal(test2) {
		fmt.Println(test2, "is pareto optimal in Game 11.5")
	}

	test3 := OutcomePair{1,1}
	if !gameConPage72.isParetoOptimal(test3) {
		fmt.Println(test3, "is not pareto optimal in Game on page 72")
	}


}
