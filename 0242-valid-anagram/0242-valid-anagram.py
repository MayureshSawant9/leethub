class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        count = collections.Counter(s)

        for char in t:
            if char not in count:
                return False
            elif count[char] == 1:
                del count[char]
            else:
                count[char] -= 1
        
        return len(count.keys()) == 0 
        