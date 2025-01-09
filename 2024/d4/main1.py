import sys
import os

def how_many_xmas(grid, x, y):
    count = 0


    return count

def search_grid(grid):    
    res = 0

    for i in range(len(grid)):


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