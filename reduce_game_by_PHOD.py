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

#Esta funcion encuentra todas las dominancias existentes entre las estrategias de ambos jugadores.
def find_dominances(game):
    #Pimero hay que generar todas las posibles permutaciones de dos en dos tanto de filas como de columnas:
	row_permutations = list(permutations(range(len(game.matrix)), 2))
	col_permutations = list(permutations(range(len(game.matrix[0])), 2))
    
	dominances_dict = {
		#Uso sets para evitar repeticion de valores:
		"dominated_rows": [],
		"dominated_cols": []
	}
    
	for r1, r2 in row_permutations:
		if not r2 in dominances_dict["dominated_rows"]:
			if is_dominant(game, "row", r1, r2):
				dominances_dict["dominated_rows"].append(r2)	
    		#dominances_dict["dominated_rows"].add(r2)
	#print("---------------------------------------")
	
	for c1, c2 in col_permutations:
		if not c2 in dominances_dict["dominated_rows"]: 	
			if is_dominant(game, "column", c1, c2):
				dominances_dict["dominated_cols"].append(c2)
    
    #Finalmente vuelvo a transformar los sets en listas:
	dominances_dict["dominated_rows"] = list(dominances_dict["dominated_rows"])
	dominances_dict["dominated_cols"] = list(dominances_dict["dominated_cols"])
    
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
	
	#se repite el proceso en forma recursiva:
	return reduce(reduced_game)

print(reduce(g21))