n, m = map(int, input().split())

preferences = {}
matches = {}
people = []
for i in range(n):
    personAndTheirPreferences = list(input().split())
    preferences[personAndTheirPreferences[0]] = (0, personAndTheirPreferences[1:])
    matches[personAndTheirPreferences[0]] = ""
    people.append(personAndTheirPreferences[0])

proposers = people[0:n//2]

def matchPairs(p, r):
    matches[p] = r
    matches[r] = p

def rPrefersP(p, r, rPreferences):
    return rPreferences.index(p) < rPreferences.index(matches[r])

while proposers: # With great help from the legendary Anton Bertelsen
    p = proposers.pop()
    pIndex, pPreferences = preferences[p]

    # check if p is free
    if matches[p] != "":
        continue
    else:
        r = pPreferences[pIndex]
        rIndex, rPreferences = preferences[r]

        # check if r is free
        if matches[r] == "":
            matchPairs(p, r)
        else:
            # check if r prefers p to their current match
            if rPrefersP(p, r, rPreferences):
                proposers.append(matches[r])
                previousMatch = matches[r]
                matches[previousMatch] = ""

                matchPairs(p, r)
            else: # r prefers their current match
                preferences[p] = (pIndex + 1, pPreferences)
                proposers.append(p)

for i in range(n//2):
    print(people[i] + " " + matches[people[i]])