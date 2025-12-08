import os
from pprint import pprint

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()


def solve_a(inp: str) -> int:
    rows = []
    for line in inp.split("\n"):
        rows.append(line)
        
    all_els = []
    for row in rows:
        els =[r for r in row.split(" ") if r != ""]
        all_els.append(els)
    
    ops = all_els[-1]
    nbrs = [[int(n) for n in ns] for ns in all_els[:-1]]
    
    sums = []
    for i in range(len(ops)):
        op = ops[i]
        s = 0 if op == "+" else 1
        for col in range(len(nbrs)):
            curr_el = nbrs[col][i]
            if op == "+":
                s += curr_el
            elif op == "*":
                s *= curr_el
                
        sums.append(s)
        
    return sum(sums)


def solve_b(inp: str) -> int:
    rows = []
    for line in inp.split("\n"):
        rows.append(line)
        
    ops = [r for r in rows[-1].split(" ") if r != ""]
    ops.reverse()
    
    pprint(rows)
    all_nbrs = []
    for row in rows[:-1]:
        r = []
        for c in row:
            r += [char for char in c]
            
        all_nbrs.append(r)
        
    pprint(all_nbrs)
    
    # Loop over cols backwards
    opi = 0
    nbrs = []
    sums = []
    for col in range(len(all_nbrs[0])-1, -1, -1):
        print(f"doing column {col=}")

        nbr = ""
        for row in range(len(all_nbrs)):
            el = all_nbrs[row][col]
            if el != ' ': nbr += el

        if not nbr:
            # done with this op, move to next
            print(nbrs)
            s = 0 if ops[opi] == "+" else 1
            if ops[opi] == "+":
                for n in nbrs:
                   s += n
            else:
                for n in nbrs:
                    s *= n

            print(sums)
            sums.append(s)

            nbrs = []
            opi += 1
        else:
            nbrs.append(int(nbr))

    # Handle the final problem
    if nbrs:
        s = 0 if ops[opi] == "+" else 1
        if ops[opi] == "+":
            for n in nbrs:
                s += n
        else:
            for n in nbrs:
                s *= n
        sums.append(s)

    print(sums)
    return sum(sums)


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
