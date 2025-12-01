import os

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()

def rotate(start: int, dir: str, steps: int) -> int:
    if dir == "L":
        
        for _ in range(steps):
            start -= 1
            if start == -1: start = 99
            
        return start
    
    if dir == "R":
        for _ in range(steps):
            start += 1
            if start == 100: start = 0
            
        return start
        
def rotate_and_count(start: int, dir: str, steps: int) -> int:
    zeros = 0
    
    if dir == "L":
        
        for _ in range(steps):
            start -= 1
            if start == 0: zeros += 1
            if start == -1: start = 99
            
    
    if dir == "R":
        for _ in range(steps):
            start += 1
            if start == 100: start = 0
            if start == 0: zeros += 1
            
    return start, zeros
        
    
def get_rots(inp):
    return [(x[0], int(x[1:])) for x in inp.split("\n")]

def solve_a(inp: str) -> int:
    rots = get_rots(inp)
    
    res = [50]
    
    for dir, steps in rots:
        next_val = rotate(res[-1], dir, steps)
        res.append(next_val)
    
    return sum([r == 0 for r in res])

def solve_b(inp: str) -> int:
    zeros = 0
    rots = get_rots(inp)
    
    res = [50]
    
    for dir, steps in rots:
        print(f"Rotating from {res[-1]}, {dir}:{steps}")
        next_val, zs = rotate_and_count(res[-1], dir, steps)
        print(f"{next_val=}, {zs=}")
        res.append(next_val)
        zeros += zs
        
    print(f"Passed zeros: {zeros}")
    print(res)
    return zeros

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
