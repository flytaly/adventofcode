import sys
import networkx as nx
import matplotlib.pyplot as plt

filename = sys.argv[1] if len(sys.argv) > 1 else "input.txt"
lines = open(filename, "r").read().splitlines()

nodes = [l.replace(":", "").split(" ") for l in lines]
G = nx.Graph()
for node, *connected in nodes:
    for n in connected:
        G.add_edge(node, n)

DRAW = True

if DRAW:
    nx.draw(G, with_labels=True)
    plt.show()

edges = nx.minimum_edge_cut(G)
print(edges)

G.remove_edges_from(edges)

count = 1
for g in nx.connected_components(G):
    count *= len(g)

print("=>", count)

