import os
from typing import Any
from pprint import pprint

import numpy as np

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()


def solve_a(inp: str) -> Any:
    points = []
    for line in inp.split("\n"):
        pos = line.split(",")
        points.append([int(a) for a in pos])
        
    # (n_points, 3)
    points = np.array(points)
    
    # (n_points, n_points)
    distances = np.linalg.norm(points[:, None, :] - points[None, :, :], axis=2)
    
    dist_triu = np.triu(distances, k=1)
    
    distmap = []
    
    for x in range(dist_triu.shape[0]):
        for y in range(dist_triu.shape[1]):
            if dist_triu[x, y] == 0: continue
            
            distmap.append(((x,y), float(dist_triu[x, y])))
            
    distmap = sorted(distmap, key=lambda x : x[1])
    
    connected_sets = []
    
    # connect together 1000 pairs of points which are closest together
    n = 0
    N = 1000
    for p, _dist in distmap:
        point1, point2 = p[0], p[1]
        print(f"connecting {point1}:{points[point1]}, {point2}:{points[point2]}")
        
        matches = set()
        for i, conset in enumerate(connected_sets):
            if point1 in conset or point2 in conset:
                matches.add(i)
                
        if len(matches) == 0:
            print(f"found no matches, adding new")
            # add as a new set
            connected_sets.append({point1, point2})
            
        if len(matches) == 1:
            print(f"found 1 match: {matches}")
            # found one connection, add both to this set
            set_idx = matches.pop()
            connected_sets[set_idx].add(point1)
            connected_sets[set_idx].add(point2)
            
        if len(matches) > 1:
            # merge all of these sets into one
            print(f"multiple matches: {matches}")
            
            first_set_idx = matches.pop()
            first_set = connected_sets[first_set_idx]
            
            for j in matches:
                for point in connected_sets[j]:
                    first_set.add(point)
                    
            connected_sets[first_set_idx] = first_set
            
            # remove the sets that were connected
            for j in matches:
                connected_sets.pop(j)
                    
        print(f"updated connections: {connected_sets}\n")
        n += 1
        if n >= N:
            break
            
    sizes = sorted([len(c) for c in connected_sets], reverse=True)
    
    p = 1
    for s in sizes[:3]:
        p *= s
            
    return p


def solve_b(inp: str) -> int:
    points = []
    for line in inp.split("\n"):
        pos = line.split(",")
        points.append([int(a) for a in pos])
        
    # (n_points, 3)
    points = np.array(points)
    
    N_POINTS = len(points)
    
    # (n_points, n_points)
    distances = np.linalg.norm(points[:, None, :] - points[None, :, :], axis=2)
    
    dist_triu = np.triu(distances, k=1)
    
    distmap = []
    
    for x in range(dist_triu.shape[0]):
        for y in range(dist_triu.shape[1]):
            if dist_triu[x, y] == 0: continue
            
            distmap.append(((x,y), float(dist_triu[x, y])))
            
    distmap = sorted(distmap, key=lambda x : x[1])
    
    connected_sets = []
    
    for p, _dist in distmap:
        point1, point2 = p[0], p[1]
        print(f"connecting {point1}:{points[point1]}, {point2}:{points[point2]}")
        
        matches = set()
        for i, conset in enumerate(connected_sets):
            if point1 in conset or point2 in conset:
                matches.add(i)
                
        if len(matches) == 0:
            print(f"found no matches, adding new")
            # add as a new set
            connected_sets.append({point1, point2})
            
        if len(matches) == 1:
            print(f"found 1 match: {matches}")
            # found one connection, add both to this set
            set_idx = matches.pop()
            connected_sets[set_idx].add(point1)
            connected_sets[set_idx].add(point2)
            
        if len(matches) > 1:
            # merge all of these sets into one
            print(f"multiple matches: {matches}")
            
            first_set_idx = matches.pop()
            first_set = connected_sets[first_set_idx]
            
            for j in matches:
                for point in connected_sets[j]:
                    first_set.add(point)
                    
            connected_sets[first_set_idx] = first_set
            
            # remove the sets that were connected
            for j in matches:
                connected_sets.pop(j)
                    
        print(f"updated connections: {connected_sets}\n")
        
        if len(connected_sets) == 1 and len(connected_sets[0]) == N_POINTS:
            s = points[point1][0] * points[point2][0]
            break
        
    return s


if __name__ == "__main__":
    # pprint(solve_a(inp))
    print(solve_b(inp))
