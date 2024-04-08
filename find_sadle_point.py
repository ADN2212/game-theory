from GameClass import Game

ga = Game(
    "Ga",
    "Nina",
    "Juan",
    [
        [3, 2, 4, 2],
        [2, 1, 3, 0],
        [2, 2, 2, 2]
    ]
)

gb = Game(
    "Gb",
    "Rose",
    "Colin",
    [
        [-2, 0, 4],
        [2, 1, 3],
        [3, -1, 2]
    ]
)

gc = Game(
    "Gc",
    "Manetha",
    "Tontin",
    [
        [4, 3, 8],
        [9, 5, 1],
        [2, 7, 6]
    ]
)

#Esta funcion resive un juego como argumento y se encarga de hallar el valor de sus sadle points y
#sus posiciones en caso de que existan, retorna un python dict con la informacion antes dicha, el cual
#estará vacio en caso de no haber sadle points.
#O(n^(r*c)) time and space (not sure)  
def find_sadle_points(game):
    matrix = game.matrix
    row_len = len(matrix[0])
    col_len = len(matrix)
    
    #Primero hallamos los minimos de las filas:
    #En este array row_mins[i] = al menor valor en la fila i
    row_mins = [None] * col_len
    
    for row_num in range(col_len):
        row_mins[row_num] = min(matrix[row_num])
    
    #Luego hallamos los maximos de las columnas:
    col_maxs = [None] * row_len
    current_col_vals = []
    
    for col_num in range(row_len):
        for row_num in range(col_len):
            current_col_vals.append(matrix[row_num][col_num])        
        col_maxs[col_num] = max(current_col_vals)
        current_col_vals = []
            
    #Buscar el maximo de los minimos de las filas:
    maxmin_rows = max(row_mins)
    #Buscar el minimo de los maximos de las columnas:
    minmax_cols = min(col_maxs)
    
    #Si son distintos no hay sadle points:
    if minmax_cols != maxmin_rows:
        return {}
    
    #Finalmente buscar los que coinciden:    
    sadle_points_positions = []

    for r in range(col_len):
        for c in range(row_len):
            if row_mins[r] == col_maxs[c] and row_mins[r] == maxmin_rows:#Tambien se puede usar el minmax_cols
                sadle_points_positions.append((r, c))
                
    return  {
                'sadle_point_value': matrix[sadle_points_positions[0][0]][sadle_points_positions[0][1]],#el valor del sadle point este en cualquiera de las posiciones.
                'positions': sadle_points_positions
            }
      
# print(find_sadle_points(ga))
# print(find_sadle_points(gb))
# print(find_sadle_points(gc))