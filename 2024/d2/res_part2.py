import os
import sys

def calculate_input_safeness(input):

    if len(input) < 2:
        return True


    direction = 0
    if input[0] - input[1] >= 0:
        direction = 1
    else:
        direction = -1

    toleration = 1
    previous = input[0]

    for i in range(1,len(input),1):

        if toleration == 0:
            if previous - input[i] >= 0:
                direction = 1
            else:
                direction = -1

        if previous - input[i] >= 0 and direction == -1 or abs(previous - input[i]) > 3:
            if toleration == 0:
                return False
            toleration -= 1
            continue
        elif previous - input[i] <= 0 and direction == 1 or abs(previous - input[i]) > 3:
            if toleration == 0:
                return False
            toleration -= 1
            continue
        previous = input[i]
    return True

def parse_file(filename):

    res = 0
    with open(filename, 'r') as f:
        for line in f:
            if line.strip():  # Check if line is not empty
                # Convert each number in the line to integer directly
                numbers = [int(num) for num in line.split()]
                res += 1 if calculate_input_safeness(numbers) else 0
    return res
            

def main():
    if len(sys.argv) != 2:
       print("Usage: " + sys.argv[0] + " <file>")
       exit(1)

    filename = sys.argv[1]
    if not os.path.exists(filename):
       print("File not found: " + filename)
       exit(1)

    res = parse_file(filename)
    print("The result is : " + str(res))
    return 0

if __name__ == "__main__":
    main()