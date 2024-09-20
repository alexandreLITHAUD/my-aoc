MARKER_SIZE=14

def check_tab(arr):
    if len(arr) < MARKER_SIZE:
        return False
    return len(set(arr)) == MARKER_SIZE

def add_element(arr, char):
    arr.insert(0, char)
    if len(arr) > MARKER_SIZE:
        arr.pop()

res = 0
arr = []
file = open('input.txt', 'r')
while 1:
    
    # read by character
    char = file.read(1)          
    if not char: 
        break
    
    res += 1

    add_element(arr, char)
    if check_tab(arr):
        print("result: ", res)
        break

file.close()