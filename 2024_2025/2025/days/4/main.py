import os

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()


def get_adj(pos:tuple[int], grid: list[list[str]]):
    # x v, y >
    x = pos[0]
    y = pos[1]
    
    width = len(grid[0])
    depth = len(grid)
    
    poss = [
        (-1, 0), # top
        (-1, 1), # top right
        (0, 1), # right
        (1, 1), # bottom right
        (1, 0), # bottom
        (1, -1), # bottom left
        (0, -1), # left
        (-1, -1) # top left
    ]
    
    # top right, right, bottom right, bottom, bottom left, left, top left
    # 
    adjs = []
    for dx, dy in poss:
        nx = x + dx
        ny = y + dy
        newpos = (nx, ny)
        
        if (nx >= 0 and nx < depth) and (ny >= 0 and ny < width):
            adjs.append((newpos, grid[newpos[0]][newpos[1]])) # ((x, y), @)
            
    return adjs
   

def get_n_rolls(pos: tuple[int], grid: list[list[int]]) -> int:
    adjs = get_adj(pos, grid)
    
    n_rolls = 0
    for poss, val in adjs:
        if val == "@": n_rolls += 1
        
    # print(f"{pos}, {grid[pos[0]][pos[1]]}: {n_rolls=}")
    return n_rolls
    
def get_all_rolls(grid):
    rollspos = []
    
    for x, row in enumerate(grid):
           for y, el in enumerate(row):
               if el == "@":
                   n_rolls = get_n_rolls((x, y), grid)
                   
                   if n_rolls < 4:
                       rollspos.append((x, y))
                       
    return rollspos

def solve_a(inp: str) -> int:
    grid = []
    
    for row in inp.split("\n"):
        grid.append([el for el in row])
        
    s = 0
        
    for x, row in enumerate(grid):
        for y, el in enumerate(row):
            if el == "@":
                n_rolls = get_n_rolls((x, y), grid)
                # print(f"{(x, y)}: {el}, {n_rolls=}")
                if n_rolls < 4: s += 1
            
    return s

def solve_b(inp: str) -> int:
    grid = []
    
    for row in inp.split("\n"):
        grid.append([el for el in row])
        
    s = 0
    
    i = 0
    while new_rollspos := get_all_rolls(grid):
        n = len(new_rollspos)
        # print(f"found {n=} at {i=}")
        s += n
        grid = update_grid(new_rollspos, grid)
        i += 1
    
    return s
    
def update_grid(new_rollspos, grid):
    for x, y in new_rollspos:
        grid[x][y] = "x"
        
    return grid

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
