class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        min_price = prices[0]
        current_profit = 0
        for price in prices:
            if price < min_price:
                min_price = price
            else:
                current_profit = max(price - min_price, current_profit)

        return current_profit


        