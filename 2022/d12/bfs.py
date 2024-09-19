from collections import deque

def parse_input(input_str):
    return [list(line.strip()) for line in input_str.split('\n') if line.strip()]

def find_start_end(grid):
    start, end = None, None
    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if cell == 'S' :
                start = (i, j)
            elif cell == 'E':
                end = (i, j)
    return start, end

def get_elevation(char):
    if char == 'S':
        return ord('a')
    elif char == 'E':
        return ord('z')
    return ord(char)

def bfs(grid, start, end):
    rows, cols = len(grid), len(grid[0])
    queue = deque([(start, 0)])
    visited = set([start])
    
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    
    while queue:
        (x, y), steps = queue.popleft()
        
        if (x, y) == end:
            return steps
        
        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            
            if 0 <= nx < rows and 0 <= ny < cols and (nx, ny) not in visited:
                if get_elevation(grid[nx][ny]) <= get_elevation(grid[x][y]) + 1:
                    visited.add((nx, ny))
                    queue.append(((nx, ny), steps + 1))
    
    return -1 

def solve(input_str):
    grid = parse_input(input_str)
    start, end = find_start_end(grid)
    return bfs(grid, start, end)


with open('input.txt', 'r') as file:
    input_str = file.read()
result = solve(input_str)
print(f"The fewest steps required: {result}")