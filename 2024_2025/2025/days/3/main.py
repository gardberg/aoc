import os

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__),INPUT)
with open(input_file, "r") as f:
    inp = f.read()

def get_largest(line: str):
    sub_largest = []
    
    for i, l in enumerate(line):
        ssl = []
        
        for r in line[i+1:]:
            ssl.append(int(l+r))
            
        if not ssl: continue
        lr_largest = max(ssl)
        # print(f"found largest {lr_largest} for {l=}, {r=}")
        sub_largest.append(lr_largest)
        
    L = max(sub_largest)
    # print(f"found max {L=} for {line}")
    return L
    

def solve_a(inp: str) -> int:
    # print(inp)
    s = 0
    # print(inp.split("\n"))
    for line in inp.split("\n"):
        # print(f"running for: {line=}")
        s += get_largest(line)
        
    return s
        

def get_largest2(inp: str):
    # 12 nbrs
    print(f"running for {inp}")
    pos = 0
    res = ""
    
    for i in range(12):
        remaining = 12 - i - 1
        end = len(inp) - remaining
        
        subindx, biggest = max(enumerate(inp[pos:end]), key = lambda x : x[1])
        found_idx = pos + subindx
        
        res += biggest
        
        pos = found_idx + 1
    
    return int(res)



def solve_b(inp: str) -> int:
    s = 0
    for line in inp.split("\n"):
        l = get_largest2(line)
        print(f"Found largest {l} for {line}")
        s += l
        
    return s
        


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
