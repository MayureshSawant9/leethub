class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        # j_set = set(jewels)
        # count = 0
        # for stone in stones:
        #     if stone in j_set:
        #         count += 1
        
        # return count

        return sum(1 for stone in stones if stone in set(jewels))