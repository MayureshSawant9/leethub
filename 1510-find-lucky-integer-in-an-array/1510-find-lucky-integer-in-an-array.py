class Solution:
    def findLucky(self, arr: List[int]) -> int:

        count = Counter(arr)

        ans = -1

        for key, value in count.items():
            if key == value and key > ans:
                ans = key
        
        return ans