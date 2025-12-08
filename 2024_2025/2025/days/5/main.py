import os

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()

   
def merge_intervals(intervals):
    result = []
    (start_candidate, stop_candidate) = intervals[0]
    for (start, stop) in intervals[1:]:
        if start <= stop_candidate:
            stop_candidate = max(stop, stop_candidate)
        else:
            result.append((start_candidate, stop_candidate))
            (start_candidate, stop_candidate) = (start, stop)
    result.append((start_candidate, stop_candidate))
    return result

def solve_a(inp: str) -> int:
    ranges = []
    ids = []
    
    current = ranges
    for line in inp.split("\n"):
        if line == "":
            current = ids
            continue
        current.append(line)
            
    spans = []
    for r in ranges:
        l = int(r.split("-")[0])
        r = int(r.split("-")[1])
        spans.append((l, r))
        
    ids = [int(i) for i in ids]
    ranges = spans
    
    ranges = sorted(ranges, key = lambda x: x[0])
    
    merged_ranges = merge_intervals(ranges)
    
    s = 0
    for id in ids:
        for range in merged_ranges:
            print(f"testing {id=}, {range=}")
            if id >= range[0] and id <= range[1]:
                s += 1
                break
    
    return s
    

def solve_b(inp: str) -> int:
    ranges = []
    ids = []
    
    current = ranges
    for line in inp.split("\n"):
        if line == "":
            current = ids
            continue
        current.append(line)
            
    spans = []
    for r in ranges:
        l = int(r.split("-")[0])
        r = int(r.split("-")[1])
        spans.append((l, r))
        
    ids = [int(i) for i in ids]
    ranges = spans
    
    ranges = sorted(ranges, key = lambda x: x[0])
    
    merged_ranges = merge_intervals(ranges)
    print(merged_ranges)
    
    s = sum([r[1]-r[0]+1 for r in merged_ranges])
   
    return s


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
