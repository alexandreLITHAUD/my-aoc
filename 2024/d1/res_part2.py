import sys
import os

def calculate_similarity_score(arr1,arr2):
    res = 0
    count = 0
    for i in arr1:
        for j in arr2:
            if i == j:
                count += 1
        res += i*count
        count = 0
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

   res = calculate_similarity_score(arr1, arr2)
   print("Similarity score : " + str(res))
   return 0

if __name__ == "__main__":
    main()