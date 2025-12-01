from pprint import pprint
from collections import defaultdict
import itertools
INPUT = "input.txt"

with open(INPUT, "r") as f:
    inp = f.read()

def parse_into_matrix(inp: str) -> list[list[str]]:
    return [list(line) for line in inp.splitlines()]

def get_locations(matrix: list[list[str]]) -> dict[str, list[tuple[int, int]]]:
    locations = defaultdict(list)
    for i, row in enumerate(matrix):
        for j, val in enumerate(row):
            if val != ".":
                locations[val].append((i, j))
    return locations

def get_antinodes(point_a: tuple[int, int], point_b: tuple[int, int]) -> list[tuple[int, int]]:
    # an antinode occurs at any point that is perfectly in line with two points
    # but only when one of the points is twice as far away as the other

    a_to_b = (point_b[0] - point_a[0], point_b[1] - point_a[1])

    b_antinode = (point_b[0] + a_to_b[0], point_b[1] + a_to_b[1])
    a_antinode = (point_a[0] - a_to_b[0], point_a[1] - a_to_b[1])

    return [b_antinode, a_antinode]

def get_antinodes_b(point_a: tuple[int, int], point_b: tuple[int, int], l: int, w: int) -> list[tuple[int, int]]:
    # instead of just returning the first two nodes in line with point a and b at a distance from them,
    # return all nodes in line with point a and b at a distance from them

    a_to_b = (point_b[0] - point_a[0], point_b[1] - point_a[1])

    b_nodes = [point_b]

    while True:
        next_b = (b_nodes[-1][0] + a_to_b[0], b_nodes[-1][1] + a_to_b[1])
        if point_in_bounds(next_b, l, w):
            b_nodes.append(next_b)
        else:
            break

    a_nodes = [point_a]

    while True:
        next_a = (a_nodes[-1][0] - a_to_b[0], a_nodes[-1][1] - a_to_b[1])
        if point_in_bounds(next_a, l, w):
            a_nodes.append(next_a)
        else:
            break

    return b_nodes + a_nodes

def point_in_bounds(point: tuple[int, int], l: int, w: int) -> bool:
    return 0 <= point[0] < l and 0 <= point[1] < w

def solve_a(inp: str) -> int:
    matrix = parse_into_matrix(inp)
    l, w = len(matrix), len(matrix[0])
    # pprint(matrix)
    locations = get_locations(matrix)

    unique_antinodes = set()

    for freq, points in locations.items():
        for point_a, point_b in itertools.combinations(points, 2):
            antinodes = get_antinodes(point_a, point_b)
            for antinode in antinodes:
                if point_in_bounds(antinode, l, w):
                    matrix[antinode[0]][antinode[1]] = "#"
                    unique_antinodes.add(antinode)
    # pprint(matrix)

    return len(unique_antinodes)

def solve_b(inp: str) -> int:
    matrix = parse_into_matrix(inp)
    # pprint(matrix)
    l, w = len(matrix), len(matrix[0])
    locations = get_locations(matrix)

    unique_antinodes = set()

    for freq, points in locations.items():
        for point_a, point_b in itertools.combinations(points, 2):
            antinodes = get_antinodes_b(point_a, point_b, l, w)
            for antinode in antinodes:
                if point_in_bounds(antinode, l, w):
                    matrix[antinode[0]][antinode[1]] = "#"
                    unique_antinodes.add(antinode)
    # pprint(matrix)

    return len(unique_antinodes)

if __name__ == "__main__":
    print(solve_a(inp))
    print(solve_b(inp))
