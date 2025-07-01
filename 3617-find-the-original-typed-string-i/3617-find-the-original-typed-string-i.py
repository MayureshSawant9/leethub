class Solution:
    def possibleStringCount(self, word: str) -> int:
        # unique = set(word)

        # return len(word) - len(unique) + 1

        prev = word[0]
        mistakes = 0
        count = 1
        for i in range(1, len(word)):
            if word[i] == prev:
                mistakes += 1
            else:
                count += 1

            prev = word[i]

        return mistakes + 1

        
        