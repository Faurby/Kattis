n = int(input())

intervals = []

for i in range(n):
    (start, end) = map(int, input().split())
    intervals.append((start, end))

# Sort by end time
intervals.sort(key=lambda x: x[1])

# Greedy algorithm
result = [intervals[0]]

for i in range(1, n):
    if intervals[i][0] >= result[-1][1]:
        result.append(intervals[i])
    
print(len(result))