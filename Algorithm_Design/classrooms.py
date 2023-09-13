activities, classrooms = map(int, input().split())

intervals = []

for i in range(activities):
    (start, end) = map(int, input().split())
    intervals.append((start, end))

# Sort by start time
intervals.sort(key=lambda x: x[0])

usedClassrooms = [0] * classrooms
activityCounter = 0

while len(intervals) > 0:
    currentActivity = intervals.pop(0)
    for i in range(classrooms):
        if usedClassrooms[i] <= currentActivity[0]:
            usedClassrooms[i] = currentActivity[1]
            activityCounter += 1
            break

print(activityCounter)