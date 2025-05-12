class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        available = collections.Counter(magazine)

        for char in ransomNote:
            if char not in available:
                return False

            elif available[char] == 1:
                del available[char]
            
            else:
                available[char] -= 1

        return True
        