INPUT = "input.txt"

with open(INPUT, "r") as f:
    inp = f.read()

def parse_into_matrix(inp: str) -> list[list[str]]:
    return [list(line) for line in inp.split("\n")]

def check_pattern(line: str, pattern: str = "XMAS") -> int:
    count = line.count(pattern) + line.count(pattern[::-1])
    return count

def get_all_lines(matrix: list[list[str]]) -> list[str]:
    height = len(matrix)
    width = len(matrix[0])
    lines = []
    
    # horizontal
    lines.extend([''.join(row) for row in matrix])
    
    # vertical
    for col in range(width):
        lines.append(''.join(matrix[row][col] for row in range(height)))
    
    # diagonal
    for start_col in range(width):
        diagonal = []
        r, c = 0, start_col
        while r < height and c < width:
            diagonal.append(matrix[r][c])
            r += 1
            c += 1
        if len(diagonal) >= 4:
            lines.append(''.join(diagonal))
    
    for start_row in range(1, height):
        diagonal = []
        r, c = start_row, 0
        while r < height and c < width:
            diagonal.append(matrix[r][c])
            r += 1
            c += 1
        if len(diagonal) >= 4:
            lines.append(''.join(diagonal))
    
    for start_col in range(width-1, -1, -1):
        diagonal = []
        r, c = 0, start_col
        while r < height and c >= 0:
            diagonal.append(matrix[r][c])
            r += 1
            c -= 1
        if len(diagonal) >= 4:
            lines.append(''.join(diagonal))
    
    for start_row in range(1, height):
        diagonal = []
        r, c = start_row, width-1
        while r < height and c >= 0:
            diagonal.append(matrix[r][c])
            r += 1
            c -= 1
        if len(diagonal) >= 4:
            lines.append(''.join(diagonal))
    
    return lines

def solve_a(inp: str) -> int:
    matrix = parse_into_matrix(inp)
    lines = get_all_lines(matrix)
    
    total_count = sum(check_pattern(line) for line in lines)
    return total_count

def check_x_pattern(matrix: list[list[str]], row: int, col: int) -> bool:
    if row == 0 or row >= len(matrix) - 1 or col == 0 or col >= len(matrix[0]) - 1:
        return False
    
    if matrix[row][col] != 'A':
        return False
    
    tl = matrix[row-1][col-1]
    tr = matrix[row-1][col+1]
    bl = matrix[row+1][col-1]
    br = matrix[row+1][col+1]
    
    valid_patterns = [
        (tl == 'M' and br == 'S' and bl == 'S' and tr == 'M'),
        (tl == 'M' and tr == 'S' and bl == 'M' and br == 'S'),
        (tl == 'S' and br == 'M' and bl == 'M' and tr == 'S'),
        (tl == 'S' and tr == 'M' and bl == 'S' and br == 'M')
    ]
    
    return any(valid_patterns)

def solve_b(inp: str) -> int:
    matrix = parse_into_matrix(inp)
    count = 0
    
    for row in range(1, len(matrix) - 1):
        for col in range(1, len(matrix[0]) - 1):
            if check_x_pattern(matrix, row, col):
                count += 1
    
    return count

if __name__ == "__main__":
    print(solve_a(inp))
    print(solve_b(inp))
