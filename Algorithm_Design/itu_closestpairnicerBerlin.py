import math
n =  int(input().strip())


allpoints = []

for i in range(n): 
    Xpoint, Ypoint = map(float, input().split(" "))
    allpoints.append((Xpoint, Ypoint))

allpoints.sort()


def ed(firstX, firstY, secondX, secondY):
    return math.sqrt(math.pow(secondX-firstX, 2) + math.pow(secondY-firstY,2) ) 

def checkdistance(dist, middlevalue, point): 
    if point[0] > middlevalue-dist and point[0] < middlevalue+dist:
        return True
    else: 
        return False

def closestl(list, dist):
    delta = ((),())
    list.sort(key=lambda tup: tup[1])
    for i in range(len(list)):
        ##kunne være en fejl her
        for j in range(i+1, len(list)): 

            if (list[j][1] - list[i][1]) >= dist:
                break
            if ed(list[i][0], list[i][1], list[j][0], list[j][1]) < dist:
                dist = ed(list[i][0],list[i][1],list[j][0], list[j][1])
                delta = ((list[i][0],list[i][1]),(list[j][0],list[j][1]))
    return (delta, dist)


def split(points, optimaldistance):
    if (len(points)<2):
        return (points, optimaldistance)
    elif len(points)==2:
        distance = ed(points[0][0],points[0][1],points[1][0],points[1][1])
        if distance < optimaldistance:
            return (((points[0][0],points[0][1]),(points[1][0],points[1][1])), distance)   
    else: 
        n = math.ceil(len(points)/2)
        lefthalf = points[:n]
        righthalf = points[n:]
        left, lefth = split(lefthalf, optimaldistance)
        right, righth = split(righthalf, optimaldistance)
        closest, minimum = (left, lefth) if lefth < righth else (right, righth)
        findl = (points[n-1][0] + points[n][0])/2
        filteredPoints = list(filter(lambda seq: checkdistance(minimum, findl, seq), points))
        final, finaldist = closestl(filteredPoints, minimum)
        return (final, finaldist) if finaldist < minimum else (closest, minimum)

hej, distt = split(allpoints, float("inf"))

print(f"{hej[0][0]} {hej[0][1]}")
print(f"{hej[1][0]} {hej[1][1]}")