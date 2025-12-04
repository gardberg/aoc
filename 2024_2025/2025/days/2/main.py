import os

# INPUT = "example.txt"
INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__),INPUT)
with open(input_file, "r") as f:
    inp = f.read()

def is_invalid_id(id: str):
    l = len(id)
    if l % 2 != 0: return False
    
    left = id[:int(l/2)]
    right = id[int(l/2):]
    
    if left == right: return True
    
    return False
    
def is_invalid_id2(id: str):
    # if we have a string repeated at least twice
    l = len(id)
    
    subset_lens = range(1, int(l/2)+1)
    
    for le in subset_lens:
        # om det inte går jämnt ut
        if l % le != 0:
            # print(f"skipped {le} for {id}:{l}")
            continue
            
            
        # print(f"{int(l / le)=}")
        subsets = [id[i*le:i*le+le] for i in range(0, int(l / le))]
        
        # print(f"{id}: {subsets}")
        
        if len(set(subsets)) == 1: return True
        
    return False
    

def solve_a(inp: str) -> int:
    ranges = inp.split(",") 
    ids = []
    for r in ranges:
       start, stop = r.split("-")
       ids.append((start, stop))
       
    invalids = []
   
    for start, stop in ids:
       
       for id in range(int(start), int(stop) + 1):
           if is_invalid_id(str(id)):
               print("found invalid id")
               print(start, stop, id)
               invalids.append(id)
           
    return sum(invalids)


def solve_b(inp: str) -> int:
    ranges = inp.split(",") 
    ids = []
    for r in ranges:
       start, stop = r.split("-")
       ids.append((start, stop))
       
    invalids = []
   
    for start, stop in ids:
       
       for id in range(int(start), int(stop) + 1):
           if is_invalid_id2(str(id)):
               # print("found invalid id")
               # print(start, stop, id)
               invalids.append(id)
           
    return sum(invalids)


if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
