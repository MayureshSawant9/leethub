class Solution:
    def romanToInt(self, s: str) -> int:
        result = 0
        value = {
            'I':1,
            'V':5,
            'X':10,
            'L':50,
            'C':100,
            'D':500,
            'M':1000
        }
        
        i = 0
        while i < len(s):
            if i < len(s) - 1 and value[s[i]] < value[s[i+1]]:
                result += value[s[i+1]] - value[s[i]]
                i += 1

            else:
                result += value[s[i]]
                
            i += 1

        return result
        