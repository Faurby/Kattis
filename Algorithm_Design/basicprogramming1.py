N, t = map(int, input().split())

array = []

arrayInput = input().split()

for i in range(N):
    array.append(int(arrayInput[i]))

if t == 1: # Print 7
        print(7)

elif t == 2: # Print Bigger, Equal, or Smaller
    if array[0] > array[1]:
        print("Bigger")
    elif array[0] == array[1]:
        print("Equal")
    else:
        print("Smaller")

elif t == 3: # Print median of first 3 elements
    subarray = array[:3]
    subarray.sort()
    print(subarray[1])

elif t == 4: # Print sum of array
    print(sum(array))

elif t == 5: # Print sum of even numbers
    evenNumbers = [x for x in array if x % 2 == 0]
    print(sum(evenNumbers))

elif t == 6: # Apply modulo 26 to each element and print corresponding letter
    for i in array:
        print(chr(i % 26 + 97), end="") # Thanks to stackoverflow for teaching me chr() and ord().

elif t == 7: # jump between indices until you reach the end or a loop
    index = 0
    oldIndex = 0
    while True:
        # jump to next index
        index = array[index]

        # if index is out of bounds, print Out and break
        if index >= len(array) or index < 0:
            print("Out")
            break

        # if index is len(array) - 1, print Done and break
        if index == len(array) - 1:
            print("Done")
            break

        # check if index is a loop
        if array[index] == oldIndex or array[index] == index:
            print("Cyclic")
            break

        oldIndex = index
