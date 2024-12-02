import os
import sys

def calculate_offset(arr1, arr2):
    res = 0
    for i in range(len(arr1)):
        res += abs(arr1[i] - arr2[i])
    return res


def parse_file(filename):

    arr1 = []
    arr2 = []

    with open(filename, 'r') as f:
        for line in f:
            if line.strip():
                num1, num2 = map(int, line.split())
                arr1.append(num1)
                arr2.append(num2)

    return arr1, arr2
            

def main():
   if len(sys.argv) != 2:
       print("Usage: " + sys.argv[0] + " <file>")
       exit(1)

   filename = sys.argv[1]
   if not os.path.exists(filename):
       print("File not found: " + filename)
       exit(1)

   arr1,arr2 = parse_file(filename)
   sorted_arr1 = sorted(arr1)
   sorted_arr2 = sorted(arr2)

   res = calculate_offset(sorted_arr1, sorted_arr2)
   print(res)
   return 0

if __name__ == "__main__":
    main()