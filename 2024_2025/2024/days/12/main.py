from pprint import pprint
import itertools


INPUT = "input.txt"
INPUT1 = "example.txt" # 140, 80
INPUT2 = "example2.txt" # 772, 436
INPUT3 = "example3.txt" # 1930, 1206
INPUT4 = "example4.txt" # E: _, 236
INPUT5 = "example5.txt" # E: _, 368

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(inp: str) -> list[list[str]]:
    return [list(row) for row in inp.splitlines()]

def get_contiguous_region(map: list[list[str]], pos: tuple[int, int]) -> set[tuple[int, int]]:
    # starts at a position in the map and returns all positons which are in a connected
    # region of the same letter. There can be several regions containing the same letter, and we only
    # return the region that includes the letter in the starting pos
    
    # all_unq_pos is a set of all unique positions in the map
    # pos is a tuple of the starting position
    # returns a set of all positions in the connected region
    
    # get the letter at the starting position
    letter = map[pos[0]][pos[1]]
    h, w = len(map), len(map[0])

    visited = set()

    def dfs(pos: tuple[int, int]) -> set[tuple[int, int]]:
        x, y = pos
        if x < 0 or x >= h or y < 0 or y >= w:
            return set()
        if map[x][y] != letter:
            return set()
        if pos in visited:
            return set()
        visited.add(pos)
        return set([pos]) | dfs((x+1, y)) | dfs((x-1, y)) | dfs((x, y+1)) | dfs((x, y-1))
    
    return letter, dfs(pos)

def get_area(region: set[tuple[int, int]]) -> int:
    return len(region)

def get_perimeter(region: set[tuple[int, int]]) -> int:

    # get the number of sides of the area spanned by the points in region
    # is point can be seen as a square, on its own with four sides
    # if two points are next to eachother, they share a side, giving 6 sides, etc.
    # for each point, look at if there is any other points next to it. if there are, subtract 1 for each
    perimeter = 0
    # print("calculating perimeter for region", region)

    for pos in region:
        sides = 4
        # print("checking", pos)
        x, y = pos
        for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            # print("checking side", x+dx, y+dy)
            if (x+dx, y+dy) in region:
                sides -= 1
        # print("sides for", pos, sides)
        perimeter += sides
    return perimeter


def get_unique_pos(full_map: list[list[str]]) -> set[tuple[int, int]]:
    h, w = len(full_map), len(full_map[0])
    all_unq_pos = set()
    for i in range(h):
        for j in range(w):
            all_unq_pos.add((i, j))
    return all_unq_pos

def get_regions(all_unq_pos: set[tuple[int, int]], full_map: list[list[str]]) -> list[set[tuple[int, int]]]:
    regions = []

    while all_unq_pos:
        pos = all_unq_pos.pop()
        letter, region = get_contiguous_region(full_map, pos)
        all_unq_pos -= region
        regions.append({letter: region})
    return regions

def calc_price(regions: list[set[tuple[int, int]]]) -> int:
    price = 0
    for region in regions:
        letter, region = list(region.items())[0]
        area = get_area(region)
        perimeter = get_perimeter(region)
        price += area * perimeter
        # print(letter, area, perimeter, region)
    return price

def solve_a(inp: str) -> int:
    full_map = parse_input(inp)
    all_unq_pos = get_unique_pos(full_map)

    # pprint(map)
    # pprint(all_unq_pos)

    regions = get_regions(all_unq_pos, full_map)
    # pprint(regions)

    price = calc_price(regions)
    return price

def get_sides(region: set[tuple[int, int]]) -> list:
    # the number of sides a region has is the number of straight sections
    # for a region with one letter (square), it has 4 sides
    # a region with two letters also has 4 sides (top, botton, left, right)
    # i.e. the number of *faces* a region has
    
    # find all positions that have edges
    edges = {} # {pos: [dir1, dir2, ...]}
    for pos in region:
        x, y = pos
        for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            if (x+dx, y+dy) not in region:
                if (x, y) not in edges:
                    edges[(x, y)] = []
                edges[(x, y)].append((dx, dy))

    visited = {}
    def dfs(edge_pos: tuple[int, int], current_dir: tuple[int, int]) -> list:
        print(f"dfs: {edge_pos} {current_dir}")
        if edge_pos in visited and current_dir in visited[edge_pos]:
            print(f"already visited {edge_pos}")
            return []
        if edge_pos not in edges.keys():
            print(f"no edges at {edge_pos}: {edges.keys()}")
            return []
        
        if current_dir not in edges[edge_pos]:
            print(f"current dir {current_dir} not in edges {edges[edge_pos]}")
            return []

        if edge_pos not in visited:
            visited[edge_pos] = []
        visited[edge_pos].append(current_dir)
        
        # if the edge is to the left or right, we should only traverse up or down
        # if the edge is to the top or bottom, we should only traverse left or right

        if current_dir == (0, 1) or current_dir == (0, -1):
            directions = [(1, 0), (-1, 0)]
        else:
            directions = [(0, 1), (0, -1)]

        next_pos1 = (edge_pos[0] + directions[0][0], edge_pos[1] + directions[0][1])
        next_pos2 = (edge_pos[0] + directions[1][0], edge_pos[1] + directions[1][1])

        return [(edge_pos, current_dir)] + dfs(next_pos1, current_dir) + dfs(next_pos2, current_dir)

    all_faces = []
    while edges:
        print(f"edges: {edges}")
        for edge_pos, edge_dirs in edges.items():
            print(f"edge_pos: {edge_pos}, edge_dirs: {edge_dirs}")
            for edge_dir in edge_dirs:
                print(f"for edge {edge_pos} in direction {edge_dir}")
                face_edges = dfs(edge_pos, edge_dir)
                print(f"done with dfs, face edges: {face_edges}")
                print()
                if face_edges: all_faces.append(face_edges)

                for pos, dir in face_edges:
                    if pos in edges:
                        edges[pos].remove(dir)

            print(f"edges after remove: {edges}")
            print()

        for pos in edges.copy():
            if len(edges[pos]) == 0:
                edges.pop(pos)

    return all_faces


def solve_b(inp: str) -> int:
    full_map = parse_input(inp)
    pprint(full_map)
    all_unq_pos = get_unique_pos(full_map)

    regions = get_regions(all_unq_pos, full_map)
    print(regions)

    price = 0
    print()
    for region in regions:
        letter, positions = list(region.items())[0]
        print(f"region {letter}")
        print(f"positions: {positions}")
        area = len(positions)
        print(f"area: {area}")
        print()
        faces = get_sides(positions)
        print(f"faces: {faces}")
        price += area * len(faces)
        print()

    return price


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
