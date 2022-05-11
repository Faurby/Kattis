a, b = input(), input()

combinations = 1
for i in range(4):
    if a[i] != b[i]:
        combinations *= 2
print(combinations)