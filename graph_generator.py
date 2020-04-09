from networkx.generators.random_graphs import erdos_renyi_graph
import random
import sys

n = int(sys.argv[1])
p = 8.0/n
g = erdos_renyi_graph(n, p)

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