import itertools
INPUT = "example.txt"
from tqdm import tqdm

with open(INPUT, "r") as f:
    inp = f.read()


def parse_input(inp: str) -> list[tuple[int, list[int]]]:
    """Parse input into list of tuples containing (id, list of numbers)"""
    result = []
    for line in inp.splitlines():
        id_str, nums_str = line.split(": ")
        id_num = int(id_str)
        nums = [int(x) for x in nums_str.split()]
        result.append((id_num, nums))
    return result

def eval_expr(ints: list[int], operators: list[str]) -> int:
    # evaluate the expression from left to right
    result = ints[0]
    for i in range(1, len(ints)):
        if operators[i - 1] == "+":
            result += ints[i]
        elif operators[i - 1] == "*":
            result *= ints[i]
        elif operators[i - 1] == "||":
            # concatenate into new number
            result = int(str(result) + str(ints[i]))
    return result

def get_results(ints: list[int], operators: list[str]) -> list[int]:
    # gets all possible results from evaluating
    # ints with all combination of operators from left to right

    results = []
    op_count = len(ints) - 1
    # loop over all possible combinations of operators
    for op_comb in itertools.product(operators, repeat=op_count):
        print(op_comb)
        results.append(eval_expr(ints, op_comb))
    return results
    

def solve_a(inp: str) -> int:
    parsed = parse_input(inp)

    OPERATORS = ["+", "*"]

    s = 0
    for res, ints in tqdm(parsed):
        results = get_results(ints, OPERATORS)
        if res in results:
            s += res
    return s

def solve_b(inp: str) -> int:
    parsed = parse_input(inp)

    OPERATORS = ["+", "*", "||"]

    s = 0
    for res, ints in tqdm(parsed):
        results = get_results(ints, OPERATORS)
        if res in results:
            s += res
    return s

if __name__ == "__main__":
    print(solve_a(inp))
    # print(solve_b(inp))
