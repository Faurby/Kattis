import math

n = int(input())

coordinates = []

for i in range(n):
    x, y = map(float, input().split())
    coordinates.append((x, y))

coordinates.sort(key=lambda x: x[0])

def dist(p1, p2):
    return ((p1[0]-p2[0])**2 + (p1[1]-p2[1])**2)**0.5

def closestPairOnLine(coordinates, minimumDistance): # For subtask 1
    closestPair = ((),())
    coordinates.sort(key=lambda point: point[1])

    for i in range(len(coordinates)):
        for j in range(i+1, len(coordinates)):
            if (coordinates[j][1] - coordinates[i][1]) >= minimumDistance:
                break
            currentDistance = dist(coordinates[i], coordinates[j])
            if currentDistance < minimumDistance:
                minimumDistance = currentDistance
                closestPair = ((coordinates[i][0], coordinates[i][1]), (coordinates[j][0], coordinates[j][1]))

    return (closestPair, minimumDistance)

def checkDistanceToMiddleLine(dist, middlevalue, point):
    return middlevalue-dist < point[0] < middlevalue+dist 

def split(points, optimaldistance): # With great help and guidance from the legendary Kristian Berlin!!
    if len(points) < 2:
        return (points, optimaldistance)
    elif len(points) == 2:
        distance = dist(points[0], points[1])
        if distance < optimaldistance:
            return (((points[0][0], points[0][1]), (points[1][0], points[1][1])), distance)
    else: # Split the points in half
        n = math.ceil(len(points)/2)
        lefthalf = points[:n]
        righthalf = points[n:]

        # Recursively find the closest pair on each side
        leftPoints, leftOptimalDistance = split(lefthalf, optimaldistance)
        rightPoints, rightOptimalDistance = split(righthalf, optimaldistance)

        # Find the closest pair of the two sides
        closestPair, minimumDistance = (leftPoints, leftOptimalDistance) if leftOptimalDistance < rightOptimalDistance else (rightPoints, rightOptimalDistance)
        
        # Find the points that are within the minimum distance of the middle line
        middleLine = (points[n-1][0] + points[n][0])/2
        filteredPoints = list(filter(lambda ppoints: checkDistanceToMiddleLine(minimumDistance, middleLine, ppoints), points))
        closestPairMiddleLine, minimumDistanceMiddleLine = closestPairOnLineQ(filteredPoints, minimumDistance)
        return (closestPairMiddleLine, minimumDistanceMiddleLine) if minimumDistanceMiddleLine < minimumDistance else (closestPair, minimumDistance)

closestPair, distance = split(coordinates, float("inf"))

print(f'{closestPair[0][0]} {closestPair[0][1]}\n{closestPair[1][0]} {closestPair[1][1]}')