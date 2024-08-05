#Ver problema 1 Cap-5
strategias_guerrillas = [(6, 0), (5, 1), (4, 2), (3, 3)]
strategias_policias = [(7, 0), (6, 1), (5, 2), (4, 3)]

game_matriz = []

def play(g, p):
	a = g[0]
	b = g[1]
	c = p[0]
	d = p[1]

	play_1 = 1 if a > c or b > d else 0
	play_2 = 1 if a > d or b > c else 0

	return 0.5 * (play_1 + play_2)


for guerr_split in strategias_guerrillas:
	row = []
	for police_split in strategias_policias:
		row.append(play(guerr_split, police_split))
	game_matriz.append(row)

for row in game_matriz:
	print(row)
