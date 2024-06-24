#Esta es la clase que crea los objetos con los que
#estare trabajando en todo mi estudio de Game Theory
#esta puede ser extendida a futuro pero el plan es pasar estos objetos como argumentos a funciones
#ya que cada algoritmo tendra sus propias nesecidades. 
#Los valores de la matriz son dados desde la perspectiva del row player
class Game():
	def __init__(self, name, row_name, column_name, matrix):
		self.name = name
		self.row_name = row_name
		self.column_name = column_name
		self.matrix = matrix

	def __str__(self):
		rep_str = f'{self.name}:\n'
		rep_str += f'Row Player: {self.row_name}, Column Player: {self.column_name} and matrix:\n'
		row_len = len(self.matrix[0])
		for row in self.matrix:
			rep_str += "["
			for i in range(row_len):
				rep_str += str(row[i])
				if i < (row_len - 1):
					rep_str += ", "
			rep_str += "]\n"

		return rep_str
			
#Ejemplo:
# g1 = Game("G1", "Rose", "Colin", 
# 	[
# 		[1 ,0, 1],
# 		[2, -10, 4]
# 	]
# )

# print(g1)
