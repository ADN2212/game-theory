package main

import "fmt"

// En un non-zero-sum game cada casilla de la matiz del juego es un par ordenado con el outcome de cada jugador
type OutcomePair struct {
	Rose  int
	Colin int
}

// again, una strategia no es mas que un arreglog de outcomes
type Strategy []OutcomePair

// y un juego no es mas que un arreglo de
type Game []Strategy

type Node struct {
	pair     OutcomePair
	outEdges int
}

type Posotion struct {
	row int
	col int
}

func (p *Posotion) isValid(maxRow, maxCol int) bool {
	return (p.row >= 0 && p.row <= maxRow) && (p.col >= 0 && p.col <= maxCol)
}


//Este algoritmo se basa en la idea de los diagramas de flujo que se muestan en el cap 11
//Como se puede ver en esto ningun outcome que sea un Nash Equilibrium tiene flechas que apunten hacia afuera,
//por lo que, si se crea un grafo en base a la matriz del juego, los nodos que no tengan aristas salientes seran los N.E
func (g *Game) searchPureNashEquilibria() []OutcomePair {

	//Primero creamos la grilla de nodos
	//Esta grilla en un grafo dirigido en el cual cada nodo respresanta un outcome de la matriz del juego
	var grid [][]Node
	for _, strategy := range *g {
		row := []Node{}
		for _, outcome := range strategy {
			row = append(row, Node{
				pair:     outcome,
				outEdges: 0,
			})
		}
		grid = append(grid, row)
	}

	s := *g
	maxI := len(*g) - 1
	maxJ := len(s[0]) - 1

	var nashEquilibria []OutcomePair
	for r, row := range grid {
		for c := range row {
			node := &grid[r][c] //Si no uso un puntero aqui no podre modificar el nodo original.
			north := Posotion{row: r - 1, col: c}
			if north.isValid(maxI, maxJ) {
				northNode := grid[north.row][north.col]
				//Sean (a, b) y (c, d) dos nodos contigues en la grilla
				//existe la arista vertical (a, b) -> (c, d) si a <= c
				if node.pair.Rose <= northNode.pair.Rose {
					node.outEdges += 1
				}
			}
			east := Posotion{row: r, col: c + 1}
			if east.isValid(maxI, maxJ) {
				esatNode := grid[east.row][east.col]
				//Sean (a, b) y (c, d) dos nodos contigues en la grilla
				//existe la arista horizontal (a, b) -> (c, d) si b <= d
				if node.pair.Colin <= esatNode.pair.Colin {
					node.outEdges += 1
				}
			}
			south := Posotion{row:r + 1, col: c}
			if south.isValid(maxI, maxJ) {
				southNode := grid[south.row][south.col]
				if node.pair.Rose <= southNode.pair.Rose {
					node.outEdges += 1
				}
			}
			west := Posotion{row: r, col: c - 1}
			if west.isValid(maxI, maxJ) {
				westNode := grid[west.row][west.col]
				if node.pair.Colin <= westNode.pair.Colin {
					node.outEdges += 1
				}
			}
			//Finalmete buscamos los nodos que no tengan aristas de salida, estos son los Equilibrios de Nash
			//ver cap 11
			if node.outEdges == 0 {
				nashEquilibria = append(nashEquilibria, node.pair)
			}
		}
	}
	return nashEquilibria
}

func main() {
	fmt.Println("Searching for Nash Equilibria: ")

	gamme11Dot1 := Game{
		{{2, 3}, {3, 2}},
		{{1, 0}, {0, 1}},
	}
	fmt.Println("Equilibrios de Nash en el juego 11.1: ")
	fmt.Println(gamme11Dot1.searchPureNashEquilibria())

	gamme11Dot2 := Game{
		{{2, 4}, {1, 0}},
		{{3, 1}, {0, 4}},
	}
	fmt.Println("Equilibrios de Nash en el juego 11.2: ")
	fmt.Println(gamme11Dot2.searchPureNashEquilibria())

	gamme11Dot3 := Game{
		{{1, 1}, {2, 5}},
		{{5, 2}, {-1, -1}},
	}
	fmt.Println("Equilibrios de Nash en el juego 11.3: ")
	fmt.Println(gamme11Dot3.searchPureNashEquilibria())

	gamme11Dot4 := Game{
		{{3, 3}, {-1, 5}},
		{{5, -1}, {0, 0}},
	}
	fmt.Println("Equilibrios de Nash en el juego 11.4: ")
	fmt.Println(gamme11Dot4.searchPureNashEquilibria())

	gamme11Dot5 := Game{
		{{0, -1}, {0, 2}, {2, 3}},
		{{0, 0}, {2, 1}, {1, -1}},
		{{2, 2}, {1, 4}, {1, -1}},
	}

	fmt.Println("Equilibrios de Nash en el juego 11.5: ")
	fmt.Println(gamme11Dot5.searchPureNashEquilibria())

	lastGameOfCap11 := Game{
		{{0, 1}, {0, 1}, {2, 4}},
		{{5, 1}, {4, 2}, {1, 0}},
		{{4, 3}, {1, 4}, {1, 0}},
	}

	fmt.Println("Equilibrios de Nash en el juego del ejercicio 4: ")
	fmt.Println(lastGameOfCap11.searchPureNashEquilibria())

}
