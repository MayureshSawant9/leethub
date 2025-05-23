class Solution:
    def isPalindrome(self, s: str) -> bool:
        pattern = '[^A-Za-z0-9]'
        phrase = re.sub(pattern, '', s).lower()
        
        l, r = 0, len(phrase) - 1
        while l <= r:
            if phrase[l] == phrase[r]:
                l += 1
                r -= 1
            else:
                return False

        return True
        