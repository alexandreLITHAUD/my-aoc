import sys
import os
from multiprocessing import Pool

def find_antinodes(args):
    map_data, coordinates = args
    found_antinodes = set()
    
    for freq_coord in coordinates:
        for coord in coordinates:
            if freq_coord != coord:
                # First direction (freq_coord -> coord)
                vector = (coord[0] - freq_coord[0], coord[1] - freq_coord[1])
                
                # REDONCENCY UP
                antinode = (coord[0] + vector[0], coord[1] + vector[1])
                while(0 <= antinode[0] < len(map_data) and 0 <= antinode[1] < len(map_data[0])):
                    found_antinodes.add(antinode)
                    antinode = (antinode[0] + vector[0], antinode[1] + vector[1])
                
                # REDONCENCY DOWN

                # Second direction (coord -> freq_coord)
                antinode = (freq_coord[0] - vector[0], freq_coord[1] - vector[1])
                while(0 <= antinode[0] < len(map_data) and 0 <= antinode[1] < len(map_data[0])):
                    found_antinodes.add(antinode)
                    antinode = (antinode[0] - vector[0], antinode[1] - vector[1])
    
    return found_antinodes

def get_frequency_list(map_data):
    coordinates = {}
    
    for row_idx, row in enumerate(map_data):
        for col_idx, char in enumerate(row):
            if char != '.':  # Skip empty spaces
                if char not in coordinates:
                    coordinates[char] = []
                coordinates[char].append((row_idx, col_idx))
    
    return coordinates

def parse_map(filename):
    map_data = []
    with open(filename, 'r') as file:
        for line in file:
            # Strip whitespace and add each line as a list of characters
            row = list(line.strip())
            if row:  # Only add non-empty rows
                map_data.append(row)
    return map_data

def main():
    if len(sys.argv) != 2:
        print("Usage: " + sys.argv[0] + " <file>")
        exit(1)
        
    filename = sys.argv[1]
    if not os.path.exists(filename):
        print("File not found: " + filename)
        exit(1)
        
    map_data = parse_map(filename)
    frequency_list = get_frequency_list(map_data)
    
    # Create args list with map_data and coordinate lists
    args = [(map_data, coords) for coords in frequency_list.values()]
    
    # Collect all antinodes from all groups
    all_antinodes = set()
    with Pool(processes=16) as pool:
        antinode_sets = pool.map(find_antinodes, args)
        for antinode_set in antinode_sets:
            all_antinodes.update(antinode_set)
    
    print(len(all_antinodes))

if __name__ == "__main__":
    main()