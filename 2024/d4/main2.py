import sys
import os

def how_many_xmas(grid, x, y):
    count = 0

    if x == 0 or x == len(grid)-1:
        return 0

    # CROSS M M
    if x-1 >= 0 and x+1 < len(grid) and y-1 >= 0 and y+1 < len(grid[0]) and grid[x-1][y-1] == 'M' and grid[x+1][y-1] == 'M' and grid[x-1][y+1] == 'S' and grid[x+1][y+1] == 'S':
        count += 1        
    # CROSS M S
    if x-1 >= 0 and x+1 < len(grid) and y-1 >= 0 and y+1 < len(grid[0]) and grid[x-1][y-1] == 'M' and grid[x+1][y-1] == 'S' and grid[x-1][y+1] == 'M' and grid[x+1][y+1] == 'S':
        count += 1
    # CROSS S S
    if x-1 >= 0 and x+1 < len(grid) and y-1 >= 0 and y+1 < len(grid[0]) and grid[x-1][y-1] == 'S' and grid[x+1][y-1] == 'S' and grid[x-1][y+1] == 'M' and grid[x+1][y+1] == 'M':
        count += 1
    # CROSS S M
    if x-1 >= 0 and x+1 < len(grid) and y-1 >= 0 and y+1 < len(grid[0]) and grid[x-1][y-1] == 'S' and grid[x+1][y-1] == 'M' and grid[x-1][y+1] == 'S' and grid[x+1][y+1] == 'M':
        count += 1

    return count

def search_grid(grid):    
    count = 0

    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == 'A':
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