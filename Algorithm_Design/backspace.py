stack = []

for c in input():
    if c == '<':
        stack.pop()
    else:
        stack.append(c)

print(''.join(stack))