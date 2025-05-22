class Solution:
    def maxRemoval(self, nums: List[int], queries: List[List[int]]) -> int:
        n = len(nums)
        queries.sort(key=lambda x: x[0])  # Sort by starting point

        used_query = []  # Min-heap for ending points
        available_query = []  # Max-heap for ending points (simulate with negation)

        query_pos = 0
        applied_count = 0

        for i in range(n):
            # Push all queries starting at i
            while query_pos < len(queries) and queries[query_pos][0] == i:
                heapq.heappush(available_query, -queries[query_pos][1])  # simulate max-heap
                query_pos += 1

            # Adjust for previously used queries overlapping with i
            nums[i] -= len(used_query)

            # Apply queries if needed
            while nums[i] > 0 and available_query and -available_query[0] >= i:
                heapq.heappush(used_query, -heapq.heappop(available_query))
                nums[i] -= 1
                applied_count += 1

            # If we couldn't manage enough queries
            if nums[i] > 0:
                return -1

            # Remove all used queries ending at i
            while used_query and used_query[0] == i:
                heapq.heappop(used_query)

        return len(queries) - applied_count

                
