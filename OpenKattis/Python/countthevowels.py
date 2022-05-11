sentence = input().lower()
vowels = "aeiou"
count = 0

for vowel in vowels:
    count += sentence.count(vowel)

print(count)