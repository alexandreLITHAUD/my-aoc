import sys

LIMIT_RED=12
LIMIT_GREEN=13
LIMIT_BLUE=14

def check_game(line: str) -> int:
    game_part = line.split(':',1)
    gameid = game_part[0].split(' ',1)[1]

    gamerounds = game_part[1].split('; ')

    for round in gamerounds:
        cubes_picked = round.split(', ')
        
        for cubes in cubes_picked:
            cubes_information = cubes.split(' ',1)
            match cubes_information[1]:
                case "red":
                    if int(cubes_information[0]) > LIMIT_RED:
                        return 0
                case "green":
                    if int(cubes_information[0]) > LIMIT_GREEN:
                        return 0
                case "blue":
                    if int(cubes_information[0]) > LIMIT_BLUE:
                        return 0
    return int(gameid)

def read_file(filename: str) -> int:
    
    res = 0
    try:
        with open(filename) as file:
            for line in file:
                res += check_game(line)
        return res                
    except FileNotFoundError:
        sys.exit(1)

def main():
    if len(sys.argv) != 2:
        print("Usage: python app.py <filename>")
        sys.exit(1)

    print(read_file(sys.argv[1]))
    

if __name__ == "__main__":
    main()