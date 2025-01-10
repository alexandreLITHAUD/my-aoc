import sys
import os

def position(map):
    for i in range(len(map)):
        for j in range(len(map[i])):
            if map[i][j] == "^":
                return i,j
    return None

def travel_map(map):
    direction = "UP"
    visited = set()  # Keep track of visited positions
    guard_position = position(map)
    if guard_position == None:
        print("Guard not found !")
        return -1    
    
    # Add initial position to visited
    visited.add(guard_position)
    tiles_traveled = 1
    
    while True:
        if direction == "UP":
            potential_guard_position = (guard_position[0]-1, guard_position[1])
            if potential_guard_position[0] < 0:
                break
            elif map[potential_guard_position[0]][potential_guard_position[1]] == "#":
                direction = "RIGHT"
            else:
                if potential_guard_position not in visited:  # Only increment if not visited
                    tiles_traveled += 1
                    visited.add(potential_guard_position)
                guard_position = potential_guard_position
                
        elif direction == "DOWN":
            potential_guard_position = (guard_position[0]+1, guard_position[1])
            if potential_guard_position[0] >= len(map):
                break
            elif map[potential_guard_position[0]][potential_guard_position[1]] == "#":
                direction = "LEFT"
            else:
                if potential_guard_position not in visited:  # Only increment if not visited
                    tiles_traveled += 1
                    visited.add(potential_guard_position)
                guard_position = potential_guard_position
                
        elif direction == "LEFT":
            potential_guard_position = (guard_position[0], guard_position[1]-1)
            if potential_guard_position[1] < 0:
                break
            elif map[potential_guard_position[0]][potential_guard_position[1]] == "#":
                direction = "UP"
            else:
                if potential_guard_position not in visited:  # Only increment if not visited
                    tiles_traveled += 1
                    visited.add(potential_guard_position)
                guard_position = potential_guard_position
                
        elif direction == "RIGHT":
            potential_guard_position = (guard_position[0], guard_position[1]+1)
            if potential_guard_position[1] >= len(map[0]):
                break
            elif map[potential_guard_position[0]][potential_guard_position[1]] == "#":
                direction = "DOWN"
            else:
                if potential_guard_position not in visited:  # Only increment if not visited
                    tiles_traveled += 1
                    visited.add(potential_guard_position)
                guard_position = potential_guard_position
    
    return tiles_traveled

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

    map = parse_map(filename)
    print(map)
    res = travel_map(map)
    print(res)

if __name__ == "__main__":
    main()   