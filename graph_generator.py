from networkx.generators.random_graphs import erdos_renyi_graph
import random
import sys
from bfs import Graph

  


n = int(sys.argv[1])

class Result:
    def __init__(self, nodes, links):
        self.nodes = nodes
        self.edges = links

class Node:
    """

    """
    def __init__(self, id_):
        self.id = id_
        self.inn = []
        self.out = []
        self.max_num = 4
    
    def add_in(self, n_id):
        if n_id == self.id:
            return 0, 0
        if self.num_in() >= 4:
            return 0, 0
        self.inn.append(n_id)
        return 1, n_id
    
    def add_out(self, n_id):
        if n_id == self.id:
            return 0, 0
        if self.num_out() >= 4:
            return 0 ,0
        self.out.append(n_id)
        return 1, n_id
    
    def num_in(self):
        return len(self.inn)
    
    def num_out(self):
        return len(self.out)

    

def gen_shimmer(n: int):
    nodes = [ Node(x ) for x in range(n)] # start with 1

    nodes_dict = {}

    for node in nodes:
        nodes_dict[node.id] = node

    finished_nodes = []

    for x in range(n-10):
        x = x
        node = nodes_dict[x]
        for i in range(4 - node.num_in()):
            added = 0
            nid = 0
            while added != 1:
                added, nid = node.add_in(random.choice(list(nodes_dict.keys())))
            nodes_dict[nid].add_out(node.id)
        for i in range(4 - node.num_out()):
            added = 0
            nid = 0
            while added != 1:
                added, nid = node.add_out(random.choice(list(nodes_dict.keys())))
            nodes_dict[nid].add_in(node.id)

        finished_nodes.append(node)
        del nodes_dict[node.id]

    edges = []
    nodes_ = []
    for node in finished_nodes:
        for in_ in node.inn:
            if not (node.id, in_) in edges:
                edges.append((node.id, in_))
        for out in node.out:
            if not (out, node.id) in edges:
                edges.append((out, node.id))
        nodes_.append(node.id)

    g = Graph()
    for edge in edges:
        g.addEdge(edge[0], edge[1])

    bfs_ordered = []
    g.BFS(0, bfs_ordered)

    result = Result([x.id for x in nodes],edges)

    return result


def gen_erdos(n: int):
    p = 8.0/n
    return erdos_renyi_graph(n, p)
g = gen_shimmer(n)
#print(g.nodes)
print("Number of nodes: " + str(len(g.nodes)))
# [0, 1, 2, 3, 4, 5]

#print(g.edges)
print("Number of edges: " + str(len(g.edges)))
# [(0, 1), (0, 2), (0, 4), (1, 2), (1, 5), (3, 4), (4, 5)]

## Generate a random string of specific characters 
def randString(length=64):
    #put your letters in the following string
    your_letters='abcdefghijklmnopqrstvwyz01234567889'
    return ''.join((random.choice(your_letters) for i in range(length)))

node_ids = []
links = {}


for node in g.nodes:
    node_ids.append(randString())

for edge in g.edges:
    try:
        links[node_ids[edge[0]]].append(node_ids[edge[1]])
    except KeyError:
        links[node_ids[edge[0]]] = [node_ids[edge[1]]]
    else:
        pass

f = open("generated_links.txt", "w")

for src, neighborlist in links.items():
    for neighbor in neighborlist:
        f.write(src+neighbor+"\n")

f.close()

f = open("generated_nodes.txt", "w")

for node in node_ids:
    f.write(node+"\n")

f.close()
print("")