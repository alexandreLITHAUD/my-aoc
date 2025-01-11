import random
import sys
import os



def generate_map(height,width):

    # Create empty map
    map = [['.' for _ in range(width)] for _ in range(height)]

    # Add random obstacles (about 10% of the map)
    num_obstacles = (width * height) // 10
    for _ in range(num_obstacles):
        x = random.randint(0, width-1)
        y = random.randint(0, height-1)
        map[y][x] = '#'

    # Place guard near center
    randomness_limit_x = random.randint (0,15)
    randomness_limit_y = random.randint (0,15)

    center_x = width // 2 + randomness_limit_x
    center_y = height // 2 + randomness_limit_y
    map[center_y][center_x] = '^'

    # Write to file
    filename = "map" + str(height) + "_" + str(width) + ".txt"
    with open(filename, 'w') as f:
        for row in map:
            f.write(''.join(row) + '\n')

    print("Map generated in " + filename)


def main():

    height = 130
    width = 130

    if len(sys.argv) != 3:
        print("No value defined ! Using base values of 130 130")
    else:
        height = sys.argv[1]
        width = sys.argv[2]
    
    generate_map(int(height),int(width))


if __name__ == "__main__":
    main()