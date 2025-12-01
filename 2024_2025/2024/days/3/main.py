
with open("input.txt", "r") as f:
    inp = f.read()

# regex find all occurences of 'mul(x,x)', where x is a whole number with 1 - 3 digits

import re

all_matches = []
for line in inp.split("\n"):
    matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", line)
    all_matches.extend(matches)

def eval_mul(x: str):
    # x: 'mul(x, x)'
    x, y = int(x[0]), int(x[1])
    return x * y

print(sum(eval_mul(x) for x in all_matches))
