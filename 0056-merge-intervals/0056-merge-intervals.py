class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:

        intervals.sort(key=lambda interval: interval[0])
        result = []

        for interval in intervals:
            
            if not result or result[-1][1] < interval[0]:
                result.append([interval[0], interval[1]])
            else:
                result[-1] = [result[-1][0], max(result[-1][1], interval[1])]

        return result