import sys
import re

def calculate_result(mul):
    res = 0
    for i in mul:
        pattern = r'mul\((\d+),(\d+)\)'
        match = re.match(pattern, i)
        num1 = int(match.group(1))
        num2 = int(match.group(2))
        res += num1 * num2
    return res

def read_file(filename):

    final_res = 0

    with open(filename, 'r') as f:
        lines = f.readlines()
        for line in lines:
            final_res += calculate_result(re.findall(r'mul\(\d+,\d+\)', line))
    
    return final_res

def main():
    if len(sys.argv) != 2:
        print("Usage: python app.py <filename>")
        sys.exit(1)
    
    result = read_file(sys.argv[1])
    print(result)

if __name__ == "__main__":
    main()