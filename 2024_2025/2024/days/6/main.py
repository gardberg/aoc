INPUT = "input.txt"

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(inp: str) -> list[list[str]]:
    out = []
    starting_pos = None
    for i, line in enumerate(inp.split("\n")):
        out.append([])
        for j, c in enumerate(line):
            if c == "^":
                starting_pos = (i, j)
            out[i].append(c)
    return out, starting_pos

def get_next_pos(pos: tuple[int, int], matrix: list[list[str]], dir: str) -> tuple[int, int]:
    # if next item in matrix from pos in dir is ".", go there
    # if it is "#", turn right
    # if it is out of matrix bounds, return None
    
    if dir == 0:
        next_pos = (pos[0] - 1, pos[1])
    elif dir == 1:
        next_pos = (pos[0], pos[1] + 1)
    elif dir == 2:
        next_pos = (pos[0] + 1, pos[1])
    elif dir == 3:
        next_pos = (pos[0], pos[1] - 1)


    if next_pos[0] < 0 or next_pos[1] < 0 or next_pos[0] >= len(matrix) or next_pos[1] >= len(matrix[0]):
        return None, None # out of bounds

    if matrix[next_pos[0]][next_pos[1]] == "#":
        # turn right
        dir = (dir + 1) % 4
        return pos, dir
    else:
        return next_pos, dir

def solve_a(inp: str) -> int:
    # parse into matrix
    matrix, starting_pos = parse_input(inp)
    dir = 0 # 0 up, 1 right, 2 down, 3 left
    # create set of visited areas
    visited = set()
    visited.add((starting_pos, dir))
    
    # step through the map until we are out of bounds
    pos = starting_pos
    while True:
        pos, dir = get_next_pos(pos, matrix, dir)
        if pos is None:
            break
        visited.add((pos, dir))
    return len(visited)

def creates_loop(starting_pos: tuple[int, int], matrix: list[list[str]]) -> bool:
    pos = starting_pos
    dir = 0
    visited = set()
    visited.add((pos, dir))
    
    while True:
        pos, dir = get_next_pos(pos, matrix, dir)
        if pos is None:
            return False
        if (pos, dir) in visited:
            return True
        visited.add((pos, dir))

def solve_b(inp: str) -> int:
    matrix, starting_pos = parse_input(inp)
    loop_count = 0

    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == "#" or (i, j) == starting_pos:
                continue
            
            # Create a deep copy of the matrix
            test_matrix = [row[:] for row in matrix]
            test_matrix[i][j] = "#"
            
            if creates_loop(starting_pos, test_matrix):
                loop_count += 1
    
    return loop_count

if __name__ == "__main__":
    print(solve_a(inp))
    print(solve_b(inp))
