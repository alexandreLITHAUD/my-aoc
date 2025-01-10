import sys
import os

def calculate_result(llist):
    res = 0
    for lst in llist:
        res += lst[(len(lst)-1)//2]
    return res

def validate_lists(mapping, lists):
    llist = []
    
    for lst in lists:
        elem = []
        validated = True
        

        
        if validated:
            llist.append(lst)
    return llist

def print_debug(mapping, lists):
    print("Mapping dictionary:")
    for key, value in mapping.items():
        print(f"{key} -> {value}")
    
    print("\nNumber lists:")
    for lst in lists:
        print(lst)

def parse_file(filename):
    """
    Parse a file containing two sections:
    1. Mapping section with pairs separated by '|'
    2. List section with numbers separated by ','
    
    Returns:
        tuple: (mapping_dict, number_lists)
        - mapping_dict: dictionary of string pairs
        - number_lists: list of integer lists
    """
    mapping_dict = {}
    number_lists = []
    
    # Read file content
    with open(filename, 'r') as file:
        content = file.read().strip().split('\n\n')
    
    # Parse mapping section
    mapping_lines = content[0].strip().split('\n')
    for line in mapping_lines:
        if '|' in line:
            left, right = line.split('|')
            if left in mapping_dict:
                # If key exists, convert to list if needed and append
                if isinstance(mapping_dict[left], list):
                    mapping_dict[left].append(right)
                else:
                    mapping_dict[left] = [mapping_dict[left], right]
            else:
                mapping_dict[left] = [right]
    
    # Parse number lists section
    if len(content) > 1:
        number_lines = content[1].strip().split('\n')
        for line in number_lines:
            numbers = [int(num) for num in line.split(',')]
            number_lists.append(numbers)
    
    return mapping_dict, number_lists


def main():
    if len(sys.argv) != 2:
        print("Usage: " + sys.argv[0] + " <file>")
        exit(1)

    filename = sys.argv[1]
    if not os.path.exists(filename):
        print("File not found: " + filename)
        exit(1)

    mapping, lists = parse_file(filename)

    print_debug(mapping, lists)

    llist = validate_lists(mapping, lists)
    reslist = calculate_result(llist)
    print(llist)
    print(reslist)

if __name__ == "__main__":
    main()  