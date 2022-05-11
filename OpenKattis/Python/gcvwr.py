gcwr, truckWeight, _ = map(int, input().split())
totalItemWeight = sum(map(int, input().split()))
print(int((0.90*(gcwr - truckWeight)) - totalItemWeight))