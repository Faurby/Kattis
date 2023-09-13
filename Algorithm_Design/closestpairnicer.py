n = int(input())

coordinates = []

for i in range(n):
    x, y = map(float, input().split())
    coordinates.append((x, y))

def dist(p1, p2):
    return ((p1[0]-p2[0])**2 + (p1[1]-p2[1])**2)**0.5

coordinates.sort(key=lambda x: x[0])

def closestPairOnLine(coordinates):
    BestDistance = dist(coordinates[0], coordinates[1])
    BestPair = (coordinates[0], coordinates[1])
    
    for i in range(n-1):
        currentDistance = dist(coordinates[i], coordinates[i+1])
        if currentDistance < BestDistance:
            BestDistance = currentDistance
            BestPair = (coordinates[i], coordinates[i+1])
    return (BestPair)

def closestPairNaive(coordinates):
    BestDistance = dist(coordinates[0], coordinates[1])
    BestPair = (coordinates[0], coordinates[1])
    
    for i in range(n-1):
        for j in range(i+1, n):
            currentDistance = dist(coordinates[i], coordinates[j])
            if currentDistance < BestDistance:
                BestDistance = currentDistance
                BestPair = (coordinates[i], coordinates[j])
    return (BestPair)

closestPair = closestPairNaive(coordinates)

print(f'{closestPair[0][0]} {closestPair[0][1]}\n{closestPair[1][0]} {closestPair[1][1]}')