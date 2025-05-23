class Solution:
    def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
        delta = []

        for num in nums:
            delta.append((num ^ k) - num)

        delta.sort(reverse=True)
        ans = sum(nums)

        for i in range(0, len(delta), 2):
            if i + 1 >= len(delta):
                break
            num1, num2 = delta[i], delta[i+1]
            if num1 + num2 > 0:
                ans += num1 + num2

        return ans