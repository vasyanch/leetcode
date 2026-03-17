import ast
import re


# 1 option. Sum each with each
# Complexity - M: O(n), T: O(n^2)

# class Solution:
#     def twoSum(nums: list[int], target: int) -> list[int]:

#         for i, a in enumerate(nums[:-1]):
#             for j, b in enumerate(nums[i + 1:]):
#                 two_sum = a + b
#                 print(two_sum)
#                 if two_sum == target:
#                     return [i, j + i + 1]


# 2 option. Find target residual, hash find
# Complexity - M: O(n), T: O(n)

# class Solution:

def twoSum(nums: list[int], target: int) -> list[int]:

    nums_dict = {value: i for i, value in enumerate(nums)}

    for i, first_summand in enumerate(nums):
        residual = target - first_summand
        
        if residual in nums_dict and nums_dict[residual] != i:
            return [i, nums_dict[residual]]






input_str = input()
input_str = input_str.replace("Input:", "").strip()

nums_str = re.search(r"nums\s*=\s*(\[[^\]]*\])", input_str).group(1)
target_str = re.search(r"target\s*=\s*([-\d]+)", input_str).group(1)
nums = ast.literal_eval(nums_str)
target = int(target_str)

print(twoSum(nums, target))