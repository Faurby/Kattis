numberOfWeights = int(input())

weights = []

for i in range(numberOfWeights):
    weights.append(int(input()))

def Subset_Sum(n, W):
    M = [[0 for _ in range(n+1)] for _ in range(W+1)]

    for i in range(n+1):
        for w in range(W+1):
            if w < weights[i-1]:
                M[i][w] = M[i-1][w]
                continue

            drop = M[i][w]
            take = M[i][w-weights[i-1]] + weights[i-1]

            M[i][w] = max(drop, take)
    
    return M[n][W]

print(Subset_Sum(numberOfWeights, sum(weights)//2)*2)