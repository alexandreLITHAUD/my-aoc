def is_safe_sequence(levels):
    """Check if a sequence of levels is safe (all increasing or all decreasing by 1-3)"""
    if len(levels) <= 1:
        return True
    
    # Check if increasing
    is_increasing = True
    for i in range(len(levels) - 1):
        diff = levels[i + 1] - levels[i]
        if diff <= 0 or diff > 3:
            is_increasing = False
            break
            
    # Check if decreasing
    is_decreasing = True
    for i in range(len(levels) - 1):
        diff = levels[i] - levels[i + 1]
        if diff <= 0 or diff > 3:
            is_decreasing = False
            break
            
    return is_increasing or is_decreasing

def is_safe_with_dampener(report):
    """
    Check if a report is safe either:
    1. As is, or
    2. By removing exactly one number
    """
    # First check if it's safe without removing anything
    levels = [int(x) for x in report.split()]
    if is_safe_sequence(levels):
        return True
        
    # Try removing each number one at a time
    for i in range(len(levels)):
        test_levels = levels[:i] + levels[i+1:]
        if is_safe_sequence(test_levels):
            return True
            
    return False

def count_safe_reports(input_data):
    """Count how many reports are safe with the Problem Dampener"""
    reports = input_data.strip().split('\n')
    safe_count = 0
    
    for report in reports:
        if is_safe_with_dampener(report):
            safe_count += 1
            
    return safe_count

# To use with the actual puzzle input:
with open('input.txt', 'r') as f:
    puzzle_input = f.read()
print(f"Puzzle result: {count_safe_reports(puzzle_input)}")