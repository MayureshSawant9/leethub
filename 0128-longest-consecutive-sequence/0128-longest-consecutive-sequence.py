class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        hash_set = set(nums)
        maximum = 0
        for num in hash_set:
            if not num-1 in hash_set:
                longest = 1
                current = num
                while current + 1 in hash_set:
                    longest += 1
                    current += 1
                maximum = max(longest, maximum)

        return maximum