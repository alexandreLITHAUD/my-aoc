import sys
import os

def parse_map(filename):
    return ""

def main():
    if len(sys.argv) != 2:
        print("Usage: " + sys.argv[0] + " <file>")
        exit(1)

    filename = sys.argv[1]
    if not os.path.exists(filename):
        print("File not found: " + filename)
        exit(1)

    map = parse_map(filename)
    print(map)

if __name__ == "__main__":
    main()   