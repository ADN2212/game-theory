from GameClass import Game

g23 = Game(
			"Game 2.2",
			"Rose",
			"Colin",
			[
				#C D E
				[1,2,2],#A
				[1,1,2],#B
				[1,1,1],#C
				[2,1,0],#D
			]
	)

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

# def is_stricly_better(
# 	one_stricly_better,
# 	game,
# 	strategy_index_1,
# 	strategy_index_2,
# 	player_type):

# 	player_name = game.row_name if player_type == "row" else game.column_name

# 	if not(one_stricly_better):
# 		print(f'La estrategia {player_name}[{strategy_index_1}] NO domina a la estrategia {player_name}[{strategy_index_2}] en {game.name}.')
# 	else:
# 		print(f'La estrategia {player_name}[{strategy_index_1}] DOMINA a la estrategia {player_name}[{strategy_index_2}] en {game.name}.')


#Esta funcion determina si una estrategia (strategy_1) de un jugador es dominante sobre otra (strategy_2)
#las estrategias se espesifican por los indices en la matriz del juego
def is_dominant(game, player_type, strategy_index_1, strategy_index_2):

	one_stricly_better = False
	
	if player_type == "row":

		strategy_1 = game.matrix[strategy_index_1]
		strategy_2 = game.matrix[strategy_index_2]

		for i in range(len(strategy_1)):
			if not(strategy_1[i] >= strategy_2[i]):
				#print(f'La estrategia {game.row_name}[{strategy_index_1}] NO domina a la estrategia {game.row_name}[{strategy_index_2}] en {game.name}.')
				return False
			
			if strategy_1[i] > strategy_2[i]:
				one_stricly_better = True

		# is_stricly_better(
		# 	one_stricly_better,
		# 	game,
		# 	strategy_index_1,
		# 	strategy_index_2,
		# 	player_type,
		# )

		return one_stricly_better

	if player_type == "column":

		strategy_1 = []
		strategy_2 = []

		for row in game.matrix:
			strategy_1.append(row[strategy_index_1])
			strategy_2.append(row[strategy_index_2])

		for i in range(len(strategy_1)):
			if not(strategy_1[i] <= strategy_2[i]):
				#print(f'La estrategia {game.column_name}[{strategy_index_1}] NO domina a la estrategia {game.column_name}[{strategy_index_2}] en {game.name}.')
				return False

			if strategy_1[i] < strategy_2[i]:
				one_stricly_better = True


		# is_stricly_better(
		# 	one_stricly_better,
		# 	game,
		# 	strategy_index_1,
		# 	strategy_index_2,
		# 	player_type,
		# )

		return one_stricly_better

# print(is_dominant(g23, "row", 0, 3))
# print(is_dominant(g21, "column", 2, 0))
# print(is_dominant(g21, "column", 2, 1))
