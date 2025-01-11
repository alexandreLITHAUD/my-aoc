import sys
import os

def position(map):
    for i in range(len(map)):
        for j in range(len(map[i])):
            if map[i][j] == "^":
                return i,j
    return None

def travel_map_simulation_infinite(map, obstacle_position_to_add, guard_base_position):
    direction = "UP"
    visited = set()  # Keep track of visited positions
    map_copy = [row[:] for row in map]
    map_copy[obstacle_position_to_add[0]][obstacle_position_to_add[1]] = "#"
    guard_position = guard_base_position
    
    while True:
        if direction == "UP":
            potential_guard_position = (guard_position[0]-1, guard_position[1])
            if potential_guard_position[0] < 0:
                break
            elif map_copy[potential_guard_position[0]][potential_guard_position[1]] == "#":
                if (guard_position[0],guard_position[1],direction) in visited:
                    return True
                else:
                    visited.add((guard_position[0],guard_position[1],direction))
                direction = "RIGHT"
            else:
                guard_position = potential_guard_position
                
        elif direction == "DOWN":
            potential_guard_position = (guard_position[0]+1, guard_position[1])
            if potential_guard_position[0] >= len(map_copy):
                break
            elif map_copy[potential_guard_position[0]][potential_guard_position[1]] == "#":
                if (guard_position[0],guard_position[1],direction) in visited:
                    return True
                else:
                    visited.add((guard_position[0],guard_position[1],direction))
                direction = "LEFT"
            else:
                guard_position = potential_guard_position
                
        elif direction == "LEFT":
            potential_guard_position = (guard_position[0], guard_position[1]-1)
            if potential_guard_position[1] < 0:
                break
            elif map_copy[potential_guard_position[0]][potential_guard_position[1]] == "#":
                if (guard_position[0],guard_position[1],direction) in visited:
                    return True
                else:
                    visited.add((guard_position[0],guard_position[1],direction))
                direction = "UP"
            else:
                guard_position = potential_guard_position
                
        elif direction == "RIGHT":
            potential_guard_position = (guard_position[0], guard_position[1]+1)
            if potential_guard_position[1] >= len(map_copy[0]):
                break
            elif map_copy[potential_guard_position[0]][potential_guard_position[1]] == "#":
                if (guard_position[0],guard_position[1],direction) in visited:
                    return True
                else:
                    visited.add((guard_position[0],guard_position[1],direction))
                direction = "DOWN"
            else:
                guard_position = potential_guard_position
    
    return False

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

    guard_base_position = position(map)
    if guard_base_position == None:
        print("Guard not found !")
        exit(1)

    res = 0

    for i in range(len(map)):
        for j in range(len(map[i])):
            if map[i][j] == ".":
                if travel_map_simulation_infinite(map, (i,j), guard_base_position):
                    res += 1
    print(res)

if __name__ == "__main__":
    main()   