import ast
import re

# 1 option. Calculate profit for each day
# Complexity - M: O(n), T: O(n)
# class Solution:
def containsDuplicate(nums: list[int]) -> bool:
    
    set_nums = set()

    for i in nums:
        if i in set_nums:
            return True
        else:
            set_nums.add(i)


    return False




input_str = input()
input_str = input_str.replace("Input:", "").strip()

nums_str = re.search(r"nums \s*=\s*(\[[^\]]*\])", input_str).group(1)
nums = ast.literal_eval(nums_str)

print(containsDuplicate(nums))