INPUT = "input.txt"

from tqdm import tqdm
from functools import lru_cache

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(inp: str) -> list[int]:
    return list(map(int, inp.split()))

def change_stone(stone: int) -> tuple[int, int | None]:
    if stone == 0:
        return 1, None
    
    # Convert to string only once and store length
    str_stone = str(stone)
    length = len(str_stone)
    
    # if even nbr of digits
    if length % 2 == 0:
        mid = length // 2
        left = int(str_stone[:mid])
        right = int(str_stone[mid:])
        return left, right
    else:
        return stone * 2024, None
    
@lru_cache(maxsize=None)
def get_stone_count_single(stone: int, n: int) -> int:
    if n == 0:
        return 1
    
    left, right = change_stone(stone)
    return (get_stone_count_single(left, n - 1) + 
            get_stone_count_single(right, n - 1) if right is not None 
            else get_stone_count_single(left, n - 1))

def solve_a(inp: str) -> int:
    stones = parse_input(inp)

    for _ in tqdm(range(75)):
        for i, stone in enumerate(stones):
            left, right = change_stone(stone)
            if right is not None:
                # replace stones with the new stone(s)
                stones[i] = (left, right)
            else:
                stones[i] = left
                
        # expand any tuples in the array
        for i, item in enumerate(stones):
            if isinstance(item, tuple):
                stones[i] = item[0]
                stones.insert(i + 1, item[1])

    return len(stones)

def get_stone_count(stones: list[int], n: int) -> int:
    if n == 0:
        return len(stones)

    if n == 1:
        return 

    for _ in range(n):
        for i, stone in enumerate(stones):
            left, right = change_stone(stone)
            if right is not None:
                stones[i] = (left, right)
            else:
                stones[i] = left
    return len(stones)

def solve_b(inp: str) -> int:
    stones = parse_input(inp)
    return sum([get_stone_count_single(stone, 75) for stone in stones])

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))
