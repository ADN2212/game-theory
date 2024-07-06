
game = [
	[17.3, 11.5],
	[-4.4, 20.6],
	[5.2, 17.0]
]

inside_expected_val = 0
out_side_expected_val = 0
in_out_expected_val = 0

#Iterar sobre todas las parejas de porcentajes posibles 
#y ver en cuales la estrategia InOut genera un valor esperado superior a las otrad dos.
for x in range(101):
	inside_expected_val = x * (1 / 100) * game[0][0] + (1 - x / 100) * game[0][1]
	print(f"E(Inside)({x}) = {inside_expected_val}")

	out_side_expected_val = x * (1 / 100) * game[1][0] + (1 - x / 100) * game[1][1]
	print(f"E(OutSide)({x}) = {out_side_expected_val}")

	in_out_expected_val = x * (1 / 100) * game[2][0] + (1 - x / 100) * game[2][1]
	print(f"E(InOut)({x}) = {in_out_expected_val}")
	print(" ")
	
	if (in_out_expected_val > inside_expected_val) and (in_out_expected_val > out_side_expected_val):
		#El ejercicio pregunta por una, pero esto sucede en mas de una ocacion, por lo que hay varias respuestas correctas:
		print(f"---------> La mejor estrategia es {x}% Run, {100 - x}% Not Run, para E(InOut) = {in_out_expected_val} <--------------")
		print(" ")
		#break
