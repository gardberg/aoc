import os
from pprint import pprint
from itertools import product
from tqdm import tqdm
from concurrent.futures import ProcessPoolExecutor
from functools import partial

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()


def show_grid(*point_sets):
    """
    Show a 2D grid where empty spots are "." and given point sets are plotted with their characters.

    Args:
        *point_sets: Variable number of tuples of (set_of_points, character)
                    Points are plotted in order, with later ones overwriting earlier ones.

    Example:
        show_grid((green_outer_points, "#"), (edge_points_set, "X"))
    """
    if not point_sets:
        print("No point sets provided")
        return

    # Collect all points to determine bounds
    all_points = set()
    for point_set, _ in point_sets:
        all_points.update(point_set)

    if not all_points:
        print("No points to display")
        return

    # Determine bounds
    xs = [p[0] for p in all_points]
    ys = [p[1] for p in all_points]
    min_x, min_y = 0, 0
    max_x, max_y = max(xs) + 2, max(ys) + 2

    # Create grid
    grid = [['.' for _ in range(max_x - min_x + 1)] for _ in range(max_y - min_y + 1)]

    # Plot points in order (later ones overwrite earlier)
    for point_set, char in point_sets:
        for x, y in point_set:
            grid[y - min_y][x - min_x] = char

    # Print grid
    for row in grid:
        print(''.join(row))


def solve_a(inp: str):
    points = []
    for line in inp.split("\n"):
        x,y = line.split(",")
        points.append((int(x), int(y)))
        
    # x >, y v
    pprint(points)
    
    areas = []
    for (x1, y1), (x2, y2) in product(points, points):
        areas.append(abs(x1 - x2 + 1) * abs(y1 - y2 + 1))
        
    return max(areas)

def get_all_points_in_rect(corner1, corner2):
    maxx = max(corner1[0], corner2[0])
    minx = min(corner1[0], corner2[0])
    maxy = max(corner1[1], corner2[1])
    miny = min(corner1[1], corner2[1])

    return set([(x, y) for x, y in product(range(minx, maxx+1), range(miny, maxy+1))])

def check_candidate(points_pair, body_points):
    (x1, y1), (x2, y2) = points_pair
    minx, maxx = min(x1, x2), max(x1, x2)
    miny, maxy = min(y1, y2), max(y1, y2)

    # Early exit: check each point and return None if any is missing
    for x in range(minx, maxx+1):
        for y in range(miny, maxy+1):
            if (x, y) not in body_points:
                return None
    return points_pair

def solve_b(inp: str):
    red_points = []
    for line in inp.split("\n"):
        x,y = line.split(",")
        red_points.append((int(x), int(y)))
        
    # how do we create a set of all points within
    green_outer_points = []
    print("getting edge points...")
    for (x1, y1), (x2, y2) in zip(red_points, red_points[1:] + [red_points[0]]):
        dir = (x2 - x1, y2 - y1) # from 2nd point to first
        
        if dir[0] != 0:
            # walk in x dir
            sign = -1 if dir[0] < 0 else 1
            
            points_to_add = [(x1 + sign * dx, y1) for dx in range(sign * dir[0])]
            
        else:
            sign = -1 if dir[1] < 0 else 1
            
            points_to_add = [(x1, y1 + sign * dy) for dy in range(sign * dir[1] + 1)]
            
        green_outer_points += points_to_add
        
    # not perfect, misses some of the corners, but if we union we should get all on the edge
    edge_points_set = set(red_points).union(set(green_outer_points))
    print("done")
        
    # show_grid((green_outer_points, "X"), (red_points, "#"))
    # show_grid((edge_points_set, "X"))
    
    miny = min([p[1] for p in edge_points_set])
    maxy = max([p[1] for p in edge_points_set])

    # Build minx_per_y and maxx_per_y
    print("building x-ranges per y...")
    minx_per_y = {}
    maxx_per_y = {}

    for x, y in edge_points_set:
        if y not in minx_per_y:
            minx_per_y[y] = x
            maxx_per_y[y] = x
        else:
            minx_per_y[y] = min(minx_per_y[y], x)
            maxx_per_y[y] = max(maxx_per_y[y], x)

    print("done")

    print("getting body points...")
    body_points = set()
    for y in tqdm(range(miny, maxy+1)):
        inside = False

        # Skip rows with no edge points
        if y not in minx_per_y:
            continue

        # Only loop through the x-range that has edge points
        for x in range(minx_per_y[y], maxx_per_y[y] + 1):
            current = (x, y)

            if current in edge_points_set:
                body_points.add(current)
                if (x+1, y) in edge_points_set:
                    pass
                else:
                    inside = not inside
            else:
                if inside:
                    body_points.add(current)

        inside = False

    print("done")
    
    print("getting candidate points...")
    with ProcessPoolExecutor(max_workers=8) as executor:
        func = partial(check_candidate, body_points=body_points)
        results = executor.map(func, product(red_points, red_points))
        candidate_points = [r for r in tqdm(results, total=len(red_points)**2) if r is not None]
    print("done")
    
    print("getting biggest area...")
    val = max([abs(p1[0] - p2[0] + 1) * abs(p1[1] - p2[1] + 1)] for p1, p2 in candidate_points)
    print("done")
            
    return val


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
