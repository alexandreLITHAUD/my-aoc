def parse_stacks(filename):
    with open(filename, 'r') as file:
        lines = file.readlines()

    # Find the line with stack numbers
    for i, line in enumerate(lines):
        if '1' in line:
            num_stacks = len(line.split())
            stack_base = i
            break

    # Initialize stacks
    stacks = [[] for _ in range(num_stacks)]

    # Parse the stacks from bottom to top
    for line in reversed(lines[:stack_base]):
        for i, crate in enumerate(line[1::4]):
            if crate != ' ':
                stacks[i].append(crate)

    return stacks

def parse_moves(filename):
    with open(filename, 'r') as file:
        lines = file.readlines()

    moves = []
    for line in lines:
        if "move" not in line:
            continue
        _, count, _, source, _, dest = line.strip().split()
        moves.append((int(count), int(source) - 1, int(dest) - 1))

    return moves

def exec_move(move, stacks):
    count, source, dest = move
    for _ in range(count):
        stacks[dest].append(stacks[source].pop())

# Example usage
stacks = parse_stacks('input.txt')
moves = parse_moves('input.txt')

res = ""

# Execute the moves
for move in moves:
    exec_move(move, stacks)

for stack in stacks:
    res += stack.pop()

print(res)

# # Print the parsed stacks
# for i, stack in enumerate(stacks, 1):
#     print(f"Stack {i}: {stack}")


