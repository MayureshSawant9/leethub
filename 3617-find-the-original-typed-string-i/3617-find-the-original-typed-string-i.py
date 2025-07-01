class Solution:
    def possibleStringCount(self, word: str) -> int:
        
        # 500/700 solution
        # failed for "ere"
        
        # unique = set(word)

        # return len(word) - len(unique) + 1

        mistakes = 0
        for i in range(1, len(word)):
            if word[i] == word[i-1]:
                mistakes += 1

            prev = word[i]

        return mistakes + 1

        
        