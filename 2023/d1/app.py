import sys

def calculate(line: str) -> int:

    digits = [char for char in line if char.isdigit()]
    # Agregate the two
    return int(digits[0]+digits[-1])

def read_file(filename: str) -> int:

    try:
        res = 0
        with open(filename) as file:
            for line in file:
                res += calculate(line)
            return res
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found")
        sys.exit(1)

def main():
    if len(sys.argv) != 2:
        print("Usage: python app.py <filename>")
        sys.exit(1)
    
    result = read_file(sys.argv[1])
    print(result)

if __name__ == "__main__":
    main()