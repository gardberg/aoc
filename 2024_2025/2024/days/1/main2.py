from collections import defaultdict

with open("input.txt", "r") as f:
    inp = f.read()

left = []
right = []

for line in inp.split("\n"):
    l, r = line.split("   ")
    left.append(l)
    right.append(r)


# count occurences in right

occurences = defaultdict(int)
for r in right:
    occurences[r] += 1

score = 0

for l in left:
    score += int(l) * occurences[l]

print(score)
