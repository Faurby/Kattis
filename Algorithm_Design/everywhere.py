iterations = int(input())
for i in range(iterations):
    cities = int(input())
    city_set = set()
    for j in range(cities):
        city_set.add(input())
    print(len(city_set))
