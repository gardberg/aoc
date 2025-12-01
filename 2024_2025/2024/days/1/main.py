
with open("input.txt", "r") as f:
    inp = f.read()

left = []
right = []

for line in inp.split("\n"):
    l, r = line.split("   ")
    left.append(l)
    right.append(r)

left.sort()
right.sort()

abs_diff = [abs(int(l) - int(r)) for l, r in zip(left, right)]

print(sum(abs_diff))
