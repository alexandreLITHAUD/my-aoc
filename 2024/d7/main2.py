import sys
import os
from multiprocessing import Pool
from itertools import product


def find_equation_solution(args):
    target, numbers = args
    operators = list(product(['+','*','|'], repeat=len(numbers)-1))

    for ops in operators:
        expression = numbers[0]
        for i in range(len(ops)):
            if ops[i] == '|':
                expression = int(str(expression) + str(numbers[i+1]))
            elif ops[i] == '*':
                expression = int(int(expression) * int(numbers[i+1]))
            else:
                expression = int(int(expression) + int(numbers[i+1])) 
        if expression == target:
            return expression
    return 0

def parse_file(filename):
    
    file_data = []
    with open(file=filename, mode='r') as file:
        for line in file:
            target = int(line.split(':')[0])
            numbers = [int(x) for x in line.split(': ')[1].split(' ')]
            print(target)
            print(numbers)
            file_data.append((target,numbers))
    return file_data

def main():
    if len(sys.argv) != 2:
        print("Missing parameters !")
        exit(1)
    
    filename = sys.argv[1]
    if not os.path.exists(filename):
        print("File not found : " + filename)
        exit(1)

    args = parse_file(filename)

    with Pool(processes=16) as pool:
        results = pool.map(find_equation_solution, args)

    final_results = sum(x for x in results if x)
    print(final_results)

if __name__ == "__main__":
    main()