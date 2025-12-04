import os

INPUT = "example.txt"
# INPUT = "input.txt"

input_file = os.path.join(os.path.dirname(__file__), INPUT)
with open(input_file, "r") as f:
    inp = f.read()


def solve_a(inp: str) -> int:
    return 0


def solve_b(inp: str) -> int:
    return 0


if __name__ == "__main__":
    print(solve_a(inp))
    print(solve_b(inp))
