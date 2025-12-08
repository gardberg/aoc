import os
from pprint import pprint
from typing import Any
from functools import cache

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()

    
    
def solve_a(inp: str) -> Any:
    diagram = []
    for line in inp.split("\n"):
        diagram.append([v for v in line])
        
    # x v, y >
    startx = 0
    for i, s in enumerate(diagram[0]):
        if s == "S":
            starty = i
            break
            
    height = len(diagram)
    
    # count the new post split position tuples in a set?
    splits = set()
    visited = set()
            
    def get_next_pos(pos: tuple[int]):
        if pos in visited: return
        else: visited.add(pos)
        # walk downwards until we either find a split, or the end
        x, y = pos[0], pos[1]
        dx = 0
        # pprint(diagram)
        while dx + x < height - 1:
            dx += 1
            if diagram[x + dx][y] == "^":
                # split left and right
                    
                splits.add((x+dx, y))
                
                l = (x + dx - 1 , y - 1)
                r = (x + dx - 1, y + 1)
                get_next_pos(l)
                get_next_pos(r)
                return
            else:
                diagram[x + dx][y] = "|"
            
        
        # here we are at bottom
        return
        
    get_next_pos((startx, starty))
    
            
    return len(splits)
    

def solve_b(inp: str) -> int:
    diagram = []
    for line in inp.split("\n"):
        diagram.append([v for v in line])
        
    # x v, y >
    startx = 0
    for i, s in enumerate(diagram[0]):
        if s == "S":
            starty = i
            break
            
    height = len(diagram)
    
    # count the new post split position tuples in a set?
    splits = set()
            
    @cache
    def get_next_pos(pos: tuple[int]):
        # walk downwards until we either find a split, or the end
        x, y = pos[0], pos[1]
        dx = 0
        # pprint(diagram)
        while dx + x < height - 1:
            dx += 1
            if diagram[x + dx][y] == "^":
                # split left and right
                    
                splits.add((x+dx, y))
                
                l = (x + dx - 1 , y - 1)
                r = (x + dx - 1, y + 1)
                return get_next_pos(l) + get_next_pos(r)
            else:
                diagram[x + dx][y] = "|"
        
        # here we are at bottom
        return 1
        
    return get_next_pos((startx, starty))
    


if __name__ == "__main__":
    # pprint(solve_a(inp))
    print(solve_b(inp))
