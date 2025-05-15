class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:

        def count_sort(word: str) -> str:
            counts = [0] * 26
            for char in word:
                counts[ord(char) - ord('a')] += 1

            result = []
            for index, count in enumerate(counts):
                result.append(chr(index + ord('a'))*count)
            return ''.join(result)            
        
        hash_map = collections.defaultdict(list)

        for word in strs:
            hash_map[count_sort(word)].append(word)
        
        return list(hash_map.values())
            