import math

numberOfTests = int(input())

for _ in range(numberOfTests):
    encodedMessage = input()
    rowLength = int(math.sqrt(len(encodedMessage)))
    rows = []
    for rowIndex in range(rowLength):
        row = ""
        for columnIndex in range(rowLength):
            index = rowLength * rowIndex + columnIndex
            letter = encodedMessage[index]
            row += letter
        rows.append(row[::-1])

    transposedRows = [*zip(*rows)]
    message = ""
    for row in transposedRows:
        for letter in row:
            message += letter
    print(message)

# ----- The code below somehow results in runtime error on Kattis, while the tests are fine >:( -----
# from io import StringIO
# from math import isqrt

# n = int(input())
# result = StringIO()
# for i in range(n):
#     sentence = input()
#     modu = isqrt(len(sentence))
#     result.write("")
#     for j in range(modu-1, -1, -1):
#         for k in range(modu):
#             result.write(sentence[j + (k*modu)])
#     result.write("\n")

# print(result.getvalue())