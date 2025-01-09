import sys
import os

def how_many_xmas(grid, x, y):
    count = 0

    # RIGHT
    if x+3 < len(grid) and grid[x+1][y] == 'M' and grid[x+2][y] == 'A' and grid[x+3][y] == 'S':
        count += 1
    # LEFT
    if x-3 >= 0 and grid[x-1][y] == 'M' and grid[x-2][y] == 'A' and grid[x-3][y] == 'S':
        count += 1
    # UP
    if y-3 >= 0 and grid[x][y-1] == 'M' and grid[x][y-2] == 'A' and grid[x][y-3] == 'S':
        count += 1
    # DOWN
    if y+3 < len(grid[0]) and grid[x][y+1] == 'M' and grid[x][y+2] == 'A' and grid[x][y+3] == 'S':
        count += 1
    # UP-RIGHT
    if x+3 < len(grid) and y-3 >= 0 and grid[x+1][y-1] == 'M' and grid[x+2][y-2] == 'A' and grid[x+3][y-3] == 'S':
        count += 1
    # UP-LEFT
    if x-3 >= 0 and y-3 >= 0 and grid[x-1][y-1] == 'M' and grid[x-2][y-2] == 'A' and grid[x-3][y-3] == 'S':
        count += 1
    # DOWN-RIGHT
    if x+3 < len(grid) and y+3 < len(grid[0]) and grid[x+1][y+1] == 'M' and grid[x+2][y+2] == 'A' and grid[x+3][y+3] == 'S':
        count += 1
    # DOWN-LEFT
    if x-3 >= 0 and y+3 < len(grid[0]) and grid[x-1][y+1] == 'M' and grid[x-2][y+2] == 'A' and grid[x-3][y+3] == 'S':
        count += 1

    return count

def search_grid(grid):    
    count = 0

    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == 'X':
                count += how_many_xmas(grid, i, j)
    return count

def parse_grid(filename):
    grid = []
    with open(filename, 'r') as f:
        for line in f:
            # Convert each line into a list of characters and remove whitespace
            row = list(line.strip())
            grid.append(row)
    return grid

def main():
    if len(sys.argv) != 2:
        print("Usage: " + sys.argv[0] + " <file>")
        exit(1)

    filename = sys.argv[1]
    if not os.path.exists(filename):
        print("File not found: " + filename)
        exit(1)

    grid = parse_grid(filename)
    res = search_grid(grid)
    print(res)

if __name__ == "__main__":
    main()   