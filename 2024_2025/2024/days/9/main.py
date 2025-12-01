INPUT = "input.txt"
from tqdm import tqdm

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(inp: str) -> list[int]:
    return [int(x) for x in inp]
    
def expand_input(inp: list[int]) -> list[int]:
    out = []
    id = 0
    for idx, i in enumerate(inp):
        if idx % 2 == 0:
            # file
            for _ in range(i):
                out.append(id)

            id += 1
        else:
            # space
            for _ in range(i):
                out.append(None)
    return out

def leftmost_space(inp: list[int]) -> int:
    for i in inp:
        if i is None:
            return inp.index(i)
    return len(inp)

def rightmost_file(inp: list[int]) -> int:
    i = len(inp) - 1
    while inp[i] is None:
        i -= 1
    return i

def swap_idx(inp: list[int], i: int, j: int) -> list[int]:
    inp[i], inp[j] = inp[j], inp[i]
    return inp

def calc_checksum(inp: list[int]) -> int:
    s = 0
    for i, j in enumerate(inp):
        if j is not None:
            s += i * j
    return s

def solve_a(inp: str) -> int:
    inp = parse_input(inp)
    inp = expand_input(inp)

    nbr_nones = inp.count(None)
    # print(f"nbr_nones: {nbr_nones}")

    leftmost = leftmost_space(inp)
    rightmost = rightmost_file(inp)

    i = 0
    while leftmost < rightmost:
        swap_idx(inp, leftmost, rightmost)
        leftmost = leftmost_space(inp)
        rightmost = rightmost_file(inp)
        i += 1
        print(f"i: {i}/{nbr_nones}")

    return calc_checksum(inp)

def parse_input_b(inp: str) -> list[int]:
    files = []
    spaces = []

    starting_idx = 0
    file_id = 0

    for i, j in enumerate(inp):
        if i % 2 == 0:
            files.append((starting_idx, int(j), file_id))
            starting_idx += int(j)
            file_id += 1
        else:
            if int(j) == 0:
                continue
            spaces.append((starting_idx, int(j)))
            starting_idx += int(j)

    return files, spaces

def find_next_swap(files: list[tuple[int, int, int]], spaces: list[tuple[int, int]], 
                   attempted_files: set[int]) -> tuple[tuple[int, int, int], tuple[int, int]]:
    # files are (starting_idx, length, id)
    # spaces are (starting_idx, length)
    # Process files in descending order of ID
    sorted_files = sorted(files, key=lambda x: x[2], reverse=True)
    
    for file in sorted_files:
        # Skip if we've already attempted to move this file
        if file[2] in attempted_files:
            continue
            
        # Mark this file as attempted, regardless of whether we can move it or not
        attempted_files.add(file[2])
            
        for space in spaces:
            # Check if file fits in space
            if file[1] <= space[1]:
                # Only move if space is to the left of file
                if space[0] < file[0]:
                    return file, space
    return None, None

def update_files_and_spaces(files: list[tuple[int, int, int]], spaces: list[tuple[int, int]], 
                          file: tuple[int, int, int], space: tuple[int, int]) -> tuple[list[tuple[int, int, int]], list[tuple[int, int]]]:
    # Remove file from current position
    files.remove(file)
    
    # Find insertion point for file at space position
    insert_idx = 0
    while insert_idx < len(files) and files[insert_idx][0] < space[0]:
        insert_idx += 1
    
    # Insert file at new position
    files.insert(insert_idx, (space[0], file[1], file[2]))
    
    # Update space - either remove it or shrink it
    spaces.remove(space)
    if space[1] > file[1]:
        spaces.append((space[0] + file[1], space[1] - file[1]))
    
    # Add new space where file was
    spaces.append((file[0], file[1]))
    
    # Sort spaces and merge adjacent ones
    spaces.sort(key=lambda x: x[0])
    spaces = merge_spaces(spaces)
    
    return files, spaces

def merge_spaces(spaces: list[tuple[int, int]]) -> list[tuple[int, int]]:
    # looks through the spaces list and checks if there are any adjacent spaces
    # two spaces in the form (start_1, length_1), (start_2, length_2) are adjacent
    # if start_2 == start_1 + length_1
    # if they are, merge them into one space
    # and remove the second space
    # we have to be careful since there can be more than two spaces adjacent to eachother
    
    i = 0
    while i < len(spaces) - 1:
        if spaces[i][0] + spaces[i][1] == spaces[i + 1][0]:
            spaces[i] = (spaces[i][0], spaces[i][1] + spaces[i + 1][1])
            spaces.pop(i + 1)
        else:
            i += 1
    return spaces

def lists_to_seq(files: list[tuple[int, int]], spaces: list[tuple[int, int]]) -> list[int]:
    seq = []
    # convert each range of file : (start_idx, length, id)
    # and space : (start_idx, length)
    # to a sequence of ids, with None where there is a space
    
    # for example: file=[(0, 3, 1), (5, 1, 2)], space=[(3, 2)] -> [1, 1, 1, None, None, 2] 
    
    files_copy = files.copy()
    spaces_copy = spaces.copy()
    
    while files_copy or spaces_copy:
        # pop the element with the lowest starting idx
        if files_copy and (not spaces_copy or files_copy[0][0] < spaces_copy[0][0]):
            # add 'length' file id ints 
            file = files_copy.pop(0)
            to_add = [file[2]] * file[1]
            seq.extend(to_add)
        else:
            space = spaces_copy.pop(0)
            to_add = [None] * space[1]
            seq.extend(to_add)

    return seq

def solve_b(inp: str) -> int:
    files, spaces = parse_input_b(inp)
    attempted_files = set()  # Track all files we've tried to move (success or fail)
    file, space = find_next_swap(files, spaces, attempted_files)
    total_files = len(files)

    while file and space:
        # print(f"files: {files}")
        # print(f"spaces: {spaces}")
        # print(lists_to_seq(files, spaces))
        # print()
        
        print(f"{len(attempted_files)}/{total_files}")
        files, spaces = update_files_and_spaces(files, spaces, file, space)
        file, space = find_next_swap(files, spaces, attempted_files)

    # print(lists_to_seq(files, spaces))
    s = calc_checksum(lists_to_seq(files, spaces))
    return s

if __name__ == "__main__":
    # print(solve_a(inp))
    print(solve_b(inp))

