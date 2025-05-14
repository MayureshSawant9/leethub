class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        
        map = dict()
        for index, num in enumerate(nums):
            if target - num in map:
                return [index, map[target-num]]
            else:
                map[num] = index        