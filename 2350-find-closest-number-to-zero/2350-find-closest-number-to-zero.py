class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        lowest_dist = float('inf')
        closest_num = None
        for num in nums:
            if abs(num) < lowest_dist:
                closest_num = num
                lowest_dist = abs(num)

            elif abs(num) == lowest_dist:
                closest_num = max(closest_num, num)

        return closest_num

        