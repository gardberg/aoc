INPUT = "input.txt"
from collections import defaultdict

with open(INPUT, "r") as f:
    inp = f.read()

def parse_input(input_text):
    # Split into sections
    lines = input_text.strip().split('\n')
    
    # Parse first section (pairs)
    pairs = []
    i = 0
    while i < len(lines) and '|' in lines[i]:
        a, b = map(int, lines[i].split('|'))
        pairs.append((a, b))
        i += 1
    
    # Skip empty line if present
    i += 1
    
    # Parse second section (sequences)
    sequences = []
    while i < len(lines):
        if lines[i]:  # Skip empty lines
            seq = list(map(int, lines[i].split(',')))
            sequences.append(seq)
        i += 1
        
    return pairs, sequences

def is_valid(seq, adj):
    for i in range(len(seq) - 1):
        if seq[i + 1] not in adj[seq[i]]:
            return False
    return True

def get_incoming_count(sequence, adj):
    incoming_count = {x: 0 for x in sequence}
    for node in sequence:
        for next_node in adj[node]:
            if next_node in sequence:
                incoming_count[next_node] += 1
    return incoming_count

def find_nodes_with_no_incoming(sequence, adj):
    incoming_count = get_incoming_count(sequence, adj)
    return [node for node in sequence if incoming_count[node] == 0]

def solve_a(inp: str) -> int:
    pairs, sequences = parse_input(inp)

    # adjacency list
    adj = defaultdict(list)
    for a, b in pairs:
        adj[a].append(b) # b must come after a

    s = 0
    for seq in sequences:
        if is_valid(seq, adj):
            # get middle element
            m = seq[len(seq) // 2]
            s += m
    return s

def sort_invalid(seq, adj):
    seq_sorted = []

    remaining = set(seq)

    while remaining:

        available = find_nodes_with_no_incoming(remaining, adj)

        node = available.pop()
        seq_sorted.append(node)
        remaining.remove(node)

    return seq_sorted

def solve_b(inp: str) -> int:
    pairs, sequences = parse_input(inp)

    adj = defaultdict(list)
    for a, b in pairs:
        adj[a].append(b) # b must come after a

    s = 0
    for seq in sequences:
        if not is_valid(seq, adj):
            # sort according to rules
            # and get middle element, add to s 
            seq_sorted = sort_invalid(seq, adj)
            m = seq_sorted[len(seq_sorted) // 2]
            s += m

    return s

if __name__ == "__main__":
    print(solve_a(inp))
    print(solve_b(inp))
