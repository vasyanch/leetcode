import ast
import re

# 1 option. Calculate profit for each day
# Complexity - M: O(n), T: O(n)
# class Solution:
def productExceptSelf(nums: list[int]) -> bool:

    left_side = [0 for i in nums]
    right_side = [0 for i in nums]
    
    product_except_self = [0 for i in nums]

    product = 1
    left_side[0] = 1

    for i in range(len(nums) - 1):
        product *= nums[i]
        left_side[i + 1] = product
    
    product = 1
    right_side[-1] = 1

    for i in range(len(nums)):
        right_side[-i - 1] = product
        product *= nums[-i - 1]

    for i in range(len(nums)):
        product_except_self[i] = left_side[i] * right_side[i]

    return product_except_self




input_str = input()
input_str = input_str.replace("Input:", "").strip()

nums_str = re.search(r"nums \s*=\s*(\[[^\]]*\])", input_str).group(1)
nums = ast.literal_eval(nums_str)

print(productExceptSelf(nums))