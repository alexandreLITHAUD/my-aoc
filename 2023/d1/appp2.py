import sys

def check_number(buffer: str) -> tuple[bool, str, int]:
    
    numbers = {
        'one': 1,
        'two': 2,
        'three': 3,
        'four': 4,
        'five': 5,
        'six': 6,
        'seven': 7,
        'eight': 8,
        'nine': 9
    }

    for word, value in numbers.items():
        buffer_index = 0
        word_index = 0

        # Save the positions of the letters that match the word
        matched_indices = []

        # Iterate over the buffer and check if the letters of the word appear in order
        while buffer_index < len(buffer) and word_index < len(word):
            if buffer[buffer_index] == word[word_index]:
                matched_indices.append(buffer_index)
                word_index += 1
            buffer_index += 1

        # If all letters in the word were matched, remove them from the buffer
        if word_index == len(word):
            # Remove the matched letters from the buffer to get the remaining ones
            remaining_buffer = [buffer[i] for i in range(len(buffer)) if i not in matched_indices]
            return (True,remaining_buffer,value)
        
    return (False,buffer, 0)

def calculate(line):
    digits = []
    buffer = ""

    # Read first number
    for char in line:
        if char.isdigit():
            digits.append(char)
        else:
            buffer += char
            res = check_number(buffer)
            if res[0]:
                digits.append(str(res[2]))
                buffer = res[1]

    # Agregate the two
    return int(digits[0]+digits[-1])

def read_file(filename: str) -> int:

    try:
        res = 0
        with open(filename) as file:
            for line in file:
                res += calculate(line)
            return res
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found")
        sys.exit(1)

def main():
    if len(sys.argv) != 2:
        print("Usage: python app.py <filename>")
        sys.exit(1)
    
    result = read_file(sys.argv[1])
    print(result)

if __name__ == "__main__":
    main()