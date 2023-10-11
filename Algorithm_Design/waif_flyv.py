childrenCount, toyCount, toyCat = map(int, input().split())

childPreferences = [[] for _ in range(childrenCount)]
for i in range(childrenCount):
    perferences = list(map(int, input().split()))[1:]
    childPreferences[i] = perferences

categories = [0 for _ in range(toyCat)]
category_lookup = {}
for i in range(toyCat):
    line = list(map(int, input().split()))[1:]
    r =  line[-1];
    categories[i] = r
    for j in range(len(line)-1):
        toyId = line[j]
        category_lookup[toyId] = i

def getCatId(toyId, lookup):
    if toyId in lookup:
        return (lookup[toyId], categories[lookup[toyId]])
    return (None, 1)

def buildGraph(childPreferences, category_lookup):
    graph = {} # [edges = (vertex, weight)]
    startEdges = {}
    for child in range(len(childPreferences)):
        perferences = childPreferences[child]
        childkey = "c" + str(child)
        childGraphPerferences = {}
        alreadyUsedCats = []
        for toyId in perferences:
            (cat, r) = getCatId(toyId,category_lookup)
            toyKey = "t" + str(toyId)
            childGraphPerferences[toyKey] = 1
            if cat is None:
                graph[toyKey] = {"T": 1}
            else:
                catKey = "g" + str(cat)
                if catKey in alreadyUsedCats: continue
                if catKey not in graph:
                    graph[catKey] = {"T": r}
                graph[toyKey] = {catKey: 1}
        graph[childkey] = childGraphPerferences
        startEdges[childkey] = 1
    graph["S"] = startEdges
    graph["T"] = {}
    return graph


def bfs(g, start, end, flow):
    queue = [(start,[start], flow)]
    explored = [start]
    while len(queue) > 0:
        (v, path, cap) = queue.pop(0)
        if v == end:
            return (path, cap)
        w = g[v]
        for (adjacentV, w) in w.items():
            if adjacentV not in explored:
                explored.append(adjacentV)
                localPath = path.copy()
                localPath.append(adjacentV)
                minCap = min(cap, w)
                queue.append((adjacentV, localPath, minCap))

def augmentCur(vEdges, nextV,flow):
    if nextV in vEdges.keys():
        weight = vEdges[nextV]
        if weight > flow:
            #Decrease
            vEdges[nextV] = weight-flow
        else:
            vEdges.pop(nextV)
    

def augmentNext(vEdges, prevV,flow):
    if prevV in vEdges.keys():
        weight = vEdges[prevV]
        vEdges[prevV] = weight+flow
    else:
        vEdges[prevV] = flow

def augment(g, path, flow):
    i = 0
    for v in path:
        if i == len(path)-1:
            break
        nextV = path[i+1]
        vEdges = g[v]
        augmentCur(vEdges, nextV, flow)
        nextVEdges = g[nextV]
        augmentNext(nextVEdges, v, flow)
        i += 1
    return g
g = buildGraph(childPreferences, category_lookup)
r = bfs(g, "S", "T", float('inf'))
flow = 0
while r is not None:
    (path, cap) = r
    flow += cap
    g = augment(g, path, cap)
    r = bfs(g, "S", "T", cap)

print(flow)