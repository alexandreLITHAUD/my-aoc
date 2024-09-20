def check_tab(arr):
    if len(arr) < 4:
        return False
    return len(set(arr)) == 4

def add_element(arr, char):
    arr.insert(0, char)
    if len(arr) > 4:
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