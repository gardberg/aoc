from pprint import pprint
from collections import defaultdict
INPUT = "input.txt"

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(inp: str) -> list[list[int]]:
    # parse into 2d array
    return [[int(x) for x in line] for line in inp.splitlines()]

def find_trailheads(inp: list[list[int]]) -> list[tuple[int, int]]:
    # a trailhead starts at 0
    trailheads = defaultdict(int)
    for i, row in enumerate(inp):
        for j, cell in enumerate(row):
            if cell == 0:
                trailheads[(i, j)] = 0
    return trailheads

def find_trails(inp: list[list[int]], trailhead: tuple[int, int]) -> int:
    # does dfs on the height map inp from the trailhead starting pos
    # and returns the number of trails, which is the number of unique 9-height positions
    # reachable from the trailhead
    # returns the number of trails
    # we can only go up down, left right
    # the path must increase by exactly one at each step

    h, w = len(inp), len(inp[0])
    
    visited = set()
    endpoints = set()
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    def dfs(i: int, j: int, height: int):
        # base cases
        # check if in bounds
        if i < 0 or i >= h or j < 0 or j >= w:
            return
        # check if visited
        if (i, j) in visited:
            return
        # check if height is valid
        if inp[i][j] != height + 1:
            return
        # add to visited
        visited.add((i, j))

        # check if we've reached the end
        if inp[i][j] == 9:
            endpoints.add((i, j))
            return

        # recurse
        for di, dj in directions:
            dfs(i + di, j + dj, height + 1)

        visited.remove((i, j))

    # chech each starting direction
    for di, dj in directions:
        dfs(trailhead[0] + di, trailhead[1] + dj, 0)

    return len(visited)

def find_unique_trails(inp: list[list[int]], trailhead: tuple[int, int]) -> int:
    h, w = len(inp), len(inp[0])
    
    unique_count = 0
    visited = set()
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    def dfs(i: int, j: int, height: int):
        # base cases
        # check if in bounds
        if i < 0 or i >= h or j < 0 or j >= w:
            return
        # check if visited
        if (i, j) in visited:
            return
        # check if height is valid
        if inp[i][j] != height + 1:
            return
        # add to visited
        visited.add((i, j))

        # check if we've reached the end
        if inp[i][j] == 9:
            nonlocal unique_count
            unique_count += 1
            visited.remove((i, j))
            return

        # recurse
        for di, dj in directions:
            dfs(i + di, j + dj, height + 1)

        visited.remove((i, j))

    # chech each starting direction
    for di, dj in directions:
        dfs(trailhead[0] + di, trailhead[1] + dj, 0)

    return unique_count

def solve_a(inp: str) -> int:
    parsed = parse_input(inp)
    # pprint(parsed)
    trailhead_scores = find_trailheads(parsed) # (i, j) -> score
    # print(trailhead_scores)
    s = 0
    for trailhead in trailhead_scores:
        s += find_trails(parsed, trailhead)
    return s

def solve_b(inp: str) -> int:
    parsed = parse_input(inp)
    trailhead_scores = find_trailheads(parsed) # (i, j) -> score
    # print(trailhead_scores)
    s = 0
    for trailhead in trailhead_scores:
        s += find_unique_trails(parsed, trailhead)
    return s

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
