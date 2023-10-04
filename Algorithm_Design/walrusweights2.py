#Made by Lasse Faurby and Philip Flyvholm
import sys
sys.setrecursionlimit(100_000)

n = int(input())
weights = [0]*n

for i in range(0, n):
    weights[i] = int(input())
weights.sort()

def prettyPrint(opt):
    #for o in opt[0]:
        i = 0
        for j in opt[n-1]:
            if(j):
                print(str(i) + ":1", end=", ")
            else:
                print(str(i) + ":0", end=", ")
            i += 1
        print()

def findOptInOpt(opt, lo=0, hi=2000, curBest = None):
    if(lo >= hi):
        return curBest
    mid = (lo + hi) // 2
    if opt[n-1][mid]:
        curBest = mid
    left = findOptInOpt(opt, lo, mid, curBest)
    right = findOptInOpt(opt, mid+1, hi, curBest)
    if(right != None and left != None):
        leftDiff = abs(1000-left)
        rightDiff = abs(1000-right)
        #print(left, leftDiff, right, rightDiff)
        if(rightDiff <= leftDiff):
            return right
        else:
            return left
    elif(right != None):
        return right
    else: 
        return left
    
def findOpt(opt, targetWeight):
    for i in range(0, targetWeight):
        if (opt[n-1][targetWeight + i]):
            return targetWeight + i
        elif (opt[n-1][targetWeight - i]):
            return targetWeight - i

def solve():
    opt = [[False for _ in range(0, 2000)] for _ in range(n)]
    for i, weight in enumerate(weights):
        for j in range(0,2000):
            if(j == 0 or j == weight):
                opt[i][j] = True
            elif(i != 0 and opt[i-1][j]):
                opt[i][j] = opt[i-1][j]
            else: # j > weight
                if (i != 0 and j > weight):
                    rest = j-weight
                    opt[i][j] = opt[i-1][rest]
    #prettyPrint(opt)
    # print(findOptInOpt(opt))
    print(findOpt(opt, 1000))

solve()