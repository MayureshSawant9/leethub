class Solution:
    def maxValue(self, events: List[List[int]], k: int) -> int:
        events.sort(key=lambda x: x[1])  # sort by end time
        n = len(events)

        # Extract just the end times for binary search
        end_times = [e[1] for e in events]

        # dp[i][j]: max value using first i events, attending at most j
        dp = [[0] * (k + 1) for _ in range(n + 1)]

        for i in range(1, n + 1):
            start, end, value = events[i - 1]

            # Find the last event that ends before start
            prev = bisect_right(end_times, start - 1)

            for j in range(1, k + 1):
                # Option 1: skip
                dp[i][j] = dp[i - 1][j]

                # Option 2: attend
                dp[i][j] = max(dp[i][j], dp[prev][j - 1] + value)

        return dp[n][k]