import ast
import re

# 1 option. Calculate profit for each day
# Complexity - M: O(n), T: O(n)
# class Solution:
def maxProfit(prices: list[int]) -> int:
    
    min_num = prices[0]
    max_num, max_sell = 0, 0

    for i, elem in enumerate(prices):
        min_num = min(min_num, elem)
        profit = elem - min_num
        max_sell = max(max_sell, profit)


    return max_sell




input_str = input()
input_str = input_str.replace("Input:", "").strip()

prices_str = re.search(r"prices\s*=\s*(\[[^\]]*\])", input_str).group(1)
prices = ast.literal_eval(prices_str)

print(maxProfit(prices))