testcases = int(input())
for i in range(testcases):
    lengthOfVector = int(input())
    vector1 = list(map(int, input().split()))
    vector2 = list(map(int, input().split()))

    vector1.sort()
    vector2.sort(reverse=True)

    scalarProduct = 0
    for j in range(lengthOfVector):
        scalarProduct += vector1[j] * vector2[j]
    print ("Case #" + str(i+1) + ": " + str(scalarProduct))