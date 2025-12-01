with open("input.txt", "r") as f:
    inp = f.read()

# regex find all occurences of 'mul(x,x)', where x is a whole number with 1 - 3 digits

import re

# find all "do()" and "don't()" in the input
do_matches = []
dont_matches = []
all_matches = [] # [[(('num1', 'num2'), (start, end)), ...], ...]  # now a list of lists

for line in inp.split("\n"):
    # Store matches for current line
    line_mul_matches = []
    matches = re.finditer(r"mul\((\d{1,3}),(\d{1,3})\)", line)
    for match in matches:
        line_mul_matches.append((match.groups(), match.span()[0]))
    all_matches.append(line_mul_matches)

    # Store do() matches for current line
    dos = re.finditer(r"do\(\)", line)
    do_matches.append([(True, m.span()[0]) for m in dos])
    
    # Store don't() matches for current line
    donts = re.finditer(r"don't\(\)", line)
    dont_matches.append([(False, m.span()[0]) for m in donts])

merged = []
for ms, dos, donts in zip(all_matches, do_matches, dont_matches):
    merged.append(ms + dos + donts)
    # sort by second element in tuple
    merged[-1].sort(key=lambda x: x[1])

print(merged)


def eval_mul(x):
    # x: (('num1', 'num2'), (start, end))
    return int(x[0]) * int(x[1])

s = 0

# we should only eval_mul on a found expression if a "do()" is the most recent string found before,
# and not to eval_mul if "don't()" if the most recent. If there are none before, it counts as a "do()"

enabled = True
for ms in merged:
    for m, idx in ms:
        if isinstance(m, tuple) and enabled:
            print("adding", m)
            s += eval_mul(m)
        elif isinstance(m, bool):
            enabled = m

print(s)