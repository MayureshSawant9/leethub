class Solution:
    def trap(self, height: List[int]) -> int:

        n = len(height)
        l_to_r_processed = [0] * n
        r_to_l_processed = [0] * n
        
        current_max = 0
        for i in range(1, n):
            current_max = max(current_max, height[i-1])
            l_to_r_processed[i] = current_max

        current_max = 0
        for i in range(n-2, 0, -1):
            current_max = max(current_max, height[i+1])
            r_to_l_processed[i] = current_max

        total_water = 0
        for i in range(n):
            water = min(l_to_r_processed[i], r_to_l_processed[i]) - height[i]
            print(f"min({l_to_r_processed[i]}, {r_to_l_processed[i]}) - {height[i]}")
            if water > 0:
                total_water += water

        return total_water
        