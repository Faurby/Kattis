
def main():
    numChildren, numToys, numCategories = map(int, input().split())
    allToys = [i for i in range(1, numToys+1)]
    childPrefferedToys = {} # {child: [toys]}
    toyCategories = {} # {category: ([toys], numOfToysThatCanBeGiven)}

    # Get children's preffered toys
    for i in range(numChildren):
        line = input().split()
        childPrefferedToys[f"child {i}"] = []
        numToysForChild = int(line[0])

        for j in range(numToysForChild):
            childPrefferedToys[f"child {i}"].append(f"toy {int(line[j+1])}")

    # Get toy categories
    for i in range(1, numCategories+1):
        line = input().split()
        numToysInCategory = int(line[0])
        toys = []

        for j in range(numToysInCategory):
            toys.append(f"toy {int(line[j+1])}")

        toyCategories[f"cat {i}"] = toys, int(line[-1]) # (toys, numOfToysThatCanBeGiven)

    leftOverToys = getLeftOverToys(allToys, toyCategories)

    graph = buildGraph(childPrefferedToys, toyCategories, leftOverToys)

    max = max_flow(graph, "s", "t")
    print(max)

def prettyPrintChildren(children):
    for child in children:
        print (f"{child} prefers toys {children[child]}")

def getLeftOverToys(allToys, toyCategories):    
    toysLeftOver = []

    for toy in allToys:
        for category in toyCategories:
            if toy not in toyCategories[category][0]:
                toysLeftOver.append(f"toy {toy}")
                break

    return toysLeftOver

def buildGraph(children, toyCategories, leftOverToys):
    graph = {}
    graph["s"] = {}

    for child in children:
        graph["s"][child] = 1

        if child not in graph:
            graph[child] = {}

        for toy in children[child]:
                if toy not in graph:
                    graph[toy] = {}

                graph[child][toy] = 1

                if toy in leftOverToys:
                    graph[toy]["t"] = 1

        for category in toyCategories:
            for toy in toyCategories[category][0]:
                if category not in graph:
                    graph[toy][category] =  1
                    graph[category]["t"] = toyCategories[category][1]


    return graph

def max_flow(graph, source, sink):
    def bfs(u, hi):
        Seen = {}

        if u == sink:
            return hi
        if u in Seen:
            return 0
        
        Seen[u] = True

        for v, cap in graph[u]:
            if cap > 0 and v not in Seen:
                Seen[v] = True
                delta = bfs(v, min(hi, cap))
                if delta:
                    graph[u][v] -= delta
                    graph[v][u] += delta

                    return delta
        return 0
    
    flow = 0

    while True:

        delta = bfs(source, float('inf'))

        if not delta:
            break
        flow += delta
        
    return flow

if __name__ == "__main__":
    main()