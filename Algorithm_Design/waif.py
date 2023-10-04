
def main():
    numChildren, numToys, numCategories = map(int, input().split())
    allToys = [i for i in range(1, numToys+1)]
    childPrefferedToys = {} # {child: [toys]}
    toyCategories = {} # {category: ([toys], numOfToysThatCanBeGiven)}

    for i in range(numChildren):
        line = input().split()
        childPrefferedToys[f"child {i}"] = []
        numToysForChild = int(line[0])

        for j in range(numToysForChild):
            childPrefferedToys[f"child {i}"].append(int(line[j+1]))

    for i in range(1, numCategories+1):
        line = input().split()
        numToysInCategory = int(line[0])
        toys = []

        for j in range(numToysInCategory):
            toys.append(int(line[j+1]))

        toyCategories[i] = toys, int(line[-1])

    leftOverToys = getLeftOverToys(allToys, toyCategories)

    toyCategories[0] = leftOverToys, len(leftOverToys)

    prettyPrintChildren(childPrefferedToys)
    print (toyCategories)

    graph = buildGraph(childPrefferedToys, toyCategories)

def prettyPrintChildren(children):
    for child in children:
        print (f"{child} prefers toys {children[child]}")

def getLeftOverToys(allToys, toyCategories):
    toysLeftOver = []

    for toy in allToys:
        for category in toyCategories:
            if toy not in toyCategories[category][0]:
                toysLeftOver.append(toy)
                break

    return toysLeftOver

def buildGraph(children, toyCategories):
    graph = {}

    for child in children:
        graph[child] = []

        for toy in children[child]:
            for category in toyCategories:
                if toy in toyCategories[category][0]:
                    graph[child].append((category, 1))
                graph[category] = []
                graph[category].append("t", toyCategories[category][1])

    return graph

if __name__ == "__main__":
    main()