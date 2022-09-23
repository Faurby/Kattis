from collections import defaultdict


N, M, K = input().split()
group = defaultdict(list)

counter = 1

for i in range(0, int(M)):
    students = input()
    for j in range(len(students)):
        group[j].append(students[j])


for i in range(1, int(K)):
    for j in range(i, 0, -1):
        print("Comparing ", group[i], " and ", group[j])
        if (set(group[i]).isdisjoint(set(group[j]))):
            print("plus 1")
            counter = counter+1

print(counter)
