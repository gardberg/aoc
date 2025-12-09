import os
from pprint import pprint
from itertools import product
from tqdm import tqdm
from functools import partial
from tqdm import tqdm

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

def solve_b(inp: str):
    red_points = []
    for line in inp.split("\n"):
        x,y = line.split(",")
        red_points.append((int(x), int(y)))
        
    # show_grid((set(red_points), "#"))
    
    from shapely import points, Polygon, plotting, box, contains
    
    red_points_poly = points(red_points)
        
    body = Polygon(red_points_poly)
    
    # get all possible rectangles
    all_pairs = []
    print("creating all pairs...")
    for (x1, y1), (x2, y2) in tqdm(product(red_points, red_points)):
        # if (x1, y1) == (x2, y2): continue
        # if ((x1, y1), (x2, y2)) in all_pairs: continue
        all_pairs.append(((x1, y1), (x2, y2)))
        
    print("done")
         
    all_recs = []
    print("creating all rectangles...")
    for p1, p2 in tqdm(all_pairs):
        all_recs.append(box(p1[0], p1[1], p2[0], p2[1]))
        
    print("done")
        
    # check which boxes are inside body
    
    sizes = []
    print("testing each rectangle...")
    for r, (p1, p2) in tqdm(zip(all_recs, all_pairs)):
        # if contains(body, r):
        if body.covers(r):
            height = abs(p2[1] - p1[1]) + 1
            width = abs(p2[0] - p1[0]) + 1
            sizes.append(height * width)
        
    print("done")
    # print(sizes)
    return(max(sizes))

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))