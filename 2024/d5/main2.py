import sys
import os

def calculate_result(llist):
    res = 0
    for lst in llist:
        res += lst[(len(lst)-1)//2]
    return res

def validate_lists(mapping, lists):
    valid_lists = []
    
    for lst in lists:
        is_valid = True
        # Convert list to set for efficient lookup
        numbers_in_list = set(lst)
        
        # Check each rule that applies to the current list
        for before, after_list in mapping.items():
            before = int(before)
            # Skip rules where either number isn't in the current list
            if before not in numbers_in_list:
                continue
                
            # Handle both single value and list of values in mapping
            if isinstance(after_list, str):
                after_list = [after_list]
                
            for after in after_list:
                after = int(after)
                if after in numbers_in_list:
                    # If both numbers are in the list, check their order
                    before_index = lst.index(before)
                    after_index = lst.index(after)
                    if before_index > after_index:
                        is_valid = False
                        break
            
            if not is_valid:
                break
                
        if is_valid:
            valid_lists.append(lst)
            
    return valid_lists

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
    print(reslist)

if __name__ == "__main__":
    main()  