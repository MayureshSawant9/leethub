class Solution:
    def maxEvents(self, events: List[List[int]]) -> int:

        events.sort()
        days = 0
        
        for event in events:
            days = max(days, event[1])

        i = 0
        n = len(events)
        min_heap = []
        result = 0

        for day in range(1, days + 1):

            while i < n and events[i][0] <= day:
                heapq.heappush(min_heap, events[i][1])
                i += 1

            while min_heap and min_heap[0] < day:
                heapq.heappop(min_heap)

            if min_heap:
                heapq.heappop(min_heap)
                result += 1

        
        return result
            

        
        