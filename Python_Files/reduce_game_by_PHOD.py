from GameClass import Game
from itertools import permutations
from dominance import is_dominant

g21 = Game(
		"Game 2.1",
		"Nina",
		"Juan",
		[
			#A B C D E
			[1,1,1,2,2],#A
			[2,1,1,1,2],#B
			[2,2,1,1,1],#C
			[2,2,2,1,0]#D
		]
	)

gif = Game(
		"SlideShare Game",
		"Manetha",
		"Tontin",
		[
			[35,65,25,5],
			[30,20,15,0],
			[40,50,0,10],
			[55,60,10,15]
		] 
	)

g35a = Game(
    "Game Cap 3, exercise 5, a",
    "Rose",
    "Colin",
    [
        [3, 0, 1],
        [-1, 2, 2],
        [1, 0, -1]
    ]
)

g35c = Game(
    "Game Cap 3,ex 5, c",
    "Rose",
    "Colin",
    [
        [4,-3,2,-4],
        [4,-4,4,-2],
        [0,1,-3,2],
        [-5,2,-7,2],
        [3,-2,2,-2]
    ]
)

tg34 = Game(
	"Juego 3.4 truncado ver ex-6",
	"Rose",
	"Colin",
	[
		[-4,2,3],
		[-1,0,4],
		[2,-1,-3]
	]
)

g34 = Game(
	"Game cap 3.4",
	"Rose",
	"Colin",
	[
		[4,-4,3,2,-3,3],
		[-1,-1,-2,0,0,4],
		[-1,2,1,-1,2,-3]
	]
)

g37 = Game(
	"Morra Game",
	"Rose",
	"Colin",
 	[
		[0,2,-3,0],
		[-2,0,0,3],
		[3,0,0,-4],
		[0,-3,4,0]
	]
 
)

g38 = Game(
	"Cuasi Simetric Game",
	"Rose",
	"Colin",
	[
		[1,2,2,2],
		[2,1,2,2],
		[2,2,1,2],
		[2,2,2,0]
	]
)

g6_vs_p7 = Game(
	"6 Guerrillas vs 7 Policias (Ej-1, Cap-5)",
	"Guerrillas",
	"Policias",
	[
		[0.5, 0.5, 1, 1],
		[1, 0.5, 0.5, 1],
		[1, 1, 0.5, 0.5],
		[1, 1, 1, 0]
	]
)

w2 = Game(
    "Missile Game Cap-5, Ex-2",
    "Blue Country",
    "Red Country",
    [
        [1,1,0,0],
        [0,1,1,0],
        [0,0,1,1],
        [0,0,0,1],
    ]   
)

#Esta funcion encuentra todas las dominancias existentes entre las estrategias de ambos jugadores.
def find_dominances(game):
    #Pimero hay que generar todas las posibles permutaciones de dos en dos tanto de filas como de columnas:
	row_permutations = list(permutations(range(len(game.matrix)), 2))
	col_permutations = list(permutations(range(len(game.matrix[0])), 2))
    
	dominances_dict = {
		"dominated_rows": [],
		"dominated_cols": []
	}
    
	for r1, r2 in row_permutations:
		#Esto evita tener que llamar la funcion is_dominat mas de una vez para cada indice de fila.
		if not r2 in dominances_dict["dominated_rows"]:
			if is_dominant(game, "row", r1, r2):
				#print(f"Rose {r1} domina a Rose {r2}.")
				dominances_dict["dominated_rows"].append(r2)
	
	for c1, c2 in col_permutations:
		if not c2 in dominances_dict["dominated_rows"]: 	
			if is_dominant(game, "column", c1, c2):
				#print(f"Colin {c1} domina a Colin {c2}.")
				dominances_dict["dominated_cols"].append(c2)
  
	print(dominances_dict)      
	return dominances_dict


#Esta funcion reduce un juego siguiendo el Principle of Higher Order of Dominance ver pag xx del libro 
def reduce(game):
	dominances = find_dominances(game)
	#Si no hay estrategias dominadas retornamos el mismo juego
	if len(dominances["dominated_rows"]) == 0 and len(dominances["dominated_cols"]) == 0:
		game.name += " reduced by PHOD" 
		return game
    
   #De lo contrario creamos uno nuevo que no tendra las estrategias dominadas    
	reduced_game = Game(
		game.name,
		game.row_name,
		game.column_name,
		[]
	)
	
	game_matrix = game.matrix
	
 	#Aqui se filtran las estrategias (filas y columnas) dominadas:
	for r in range(len(game_matrix)):
		if not(r in dominances["dominated_rows"]):
			row_to_add = []
			for c in range(len(game_matrix[r])):
				if not(c in dominances["dominated_cols"]):
					row_to_add.append(game_matrix[r][c])
			reduced_game.matrix.append(row_to_add)
	
	print(reduced_game)
	#se repite el proceso en forma recursiva:
	return reduce(reduced_game)

#print(reduce(g21))
#print(reduce(gif))
#print(reduce(g35a))
#print(reduce(g35c))
#print(reduce(tg34))
#print(reduce(g34))
#print(reduce(g37))
#print(reduce(g38))
#print(reduce(g6_vs_p7))
print(reduce(w2))
