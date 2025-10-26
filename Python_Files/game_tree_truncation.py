class Node():
    def __init__(self, player_name, payoff = None, children = []):
        self.player_name = player_name
        self.payoff = payoff
        self.children = children
        
    def __str__(self):
        return f'Node({self.player_name}, {self.payoff}, {len(self.children)})'

#Game de la figura 7.3, pag 41:
#Level 4:
#n5 childs
n11 = Node(player_name='Rose', payoff = 5)
n12 = Node(player_name='Rose', payoff = -1)
n13 = Node(player_name='Rose', payoff = 0)

#n6 childs
n14 = Node(player_name='Rose', payoff = 3)

#n7 childs
n15 = Node(player_name='Rose', payoff = -2)
n16 = Node(player_name='Rose', payoff = 0)

#n8 childs
n17 = Node(player_name='Rose', payoff = 1)

#n9 childs
n18 = Node(player_name='Rose', payoff = -2)
n19 = Node(player_name='Rose', payoff = 1)

#n10 childs
n20 = Node(player_name='Rose', payoff = -4)
n21 = Node(player_name='Rose', payoff = 0)
n22 = Node(player_name='Rose', payoff = 1)
n23 = Node(player_name='Rose', payoff = -3)

#Level 3:
n5 = Node(player_name="Colin", children = [n11, n12, n13])
n6 = Node(player_name="Colin", children = [n14])
n7 = Node(player_name="Colin", children = [n15, n16])
n8 = Node(player_name="Colin", children = [n17])
n9 = Node(player_name="Colin", children = [n18, n19])
n10 = Node(player_name="Colin", children = [n20, n21, n22, n23])

#Level 2:
n2 = Node(player_name="Rose", children = [n5, n6])
n3 = Node(player_name="Rose", children = [n7, n8])
n4 = Node(player_name="Rose", children = [n9, n10])

#Level 1
game7dot3 = Node(player_name="Colin", children = [n2, n3, n4])
#Game fig 7.4 pagina 43:
#Another way to write the tree game:
game7dot4 = Node(
    player_name="Colin",
    children=[
        Node(
            player_name="Rose",
            children=[
                Node(
                    player_name="Colin",
                    children=[
                        Node(player_name="Rose", payoff = -1),
                        Node(player_name="Rose", payoff = 2)
                    ]
                ),
                Node(
                    player_name="Colin",
                    children=[
                        Node(player_name="Rose", payoff = 1),
                        Node(player_name="Rose", payoff = 0)
                    ]
                )   
            ]
        ),
        Node(
            player_name="Rose",
            children=[
                Node(
                    player_name="Colin",
                    children=[
                        Node(player_name="Rose", payoff = -1),
                        Node(player_name="Rose", payoff = 3)
                    ]
                ),
                Node(
                    player_name="Colin",
                    children=[
                        Node(player_name="Rose", payoff = 4),
                        Node(player_name="Rose", payoff = -2)
                    ]
                )   
            ]
        ),
    ]
)

#Esta funcion toma como argumento el root node de un Game Tree
#Aplica el truncation method para calcular el valor del juego.
def truncate(node: Node):
    
    #Si es una hoja solo retorna el payoff
    if node.payoff != None:
        return node.payoff

    #Desde la perspectiva de colin se busca tomar el minimo de los payoffs para Rose.
    if node.player_name == 'Colin':
        return min([truncate(child) for child in node.children])

    #Desde la perspectiva de Rose se busca tomar el maximo de los payoffs para si misma.
    if node.player_name == 'Rose':
        return max([truncate(child) for child in node.children])


print(truncate(game7dot3))
print(truncate(game7dot4))
