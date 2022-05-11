import sys

for i in sys.stdin:
    a = int(i)
    b = 100 - a
    print(1 + b/a)
    print(1 + a/b)