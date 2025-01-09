import sys
import re

def calculate_result(i, enable):

    if not enable:
        return 0

    res = 0

    pattern = r'mul\((\d+),(\d+)\)'
    match = re.match(pattern, i)
    num1 = int(match.group(1))
    num2 = int(match.group(2))
    return num1 * num2

def read_file(filename):

    final_res = 0
    enable = True

    with open(filename, 'r') as f:
        lines = f.readlines()
        for line in lines:
            matches = re.findall(r'mul\(\d+,\d+\)|do\(\)|don\'t\(\)', line)

            for i in matches:
                if i == 'do()':
                    enable = True
                elif i == 'don\'t()':
                    enable = False
                else:
                    final_res += calculate_result(i, enable)
    
    return final_res

def main():
    if len(sys.argv) != 2:
        print("Usage: python app.py <filename>")
        sys.exit(1)
    
    result = read_file(sys.argv[1])
    print(result)

if __name__ == "__main__":
    main()