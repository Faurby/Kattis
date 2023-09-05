n, m = map(int, input().split())

# Stable marriage problem
# insert preferences into dictionary
people = list(input().split())

preferences = {}
matches = {}
for i in range(n):
    preferences[people[0]] = people[1:]
    matches[people[0]] = ""

print("Preferences")
print(preferences)

proposers = people[0:n//2]

