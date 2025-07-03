class Solution:
    def kthCharacter(self, k: int) -> str:
        if k == 1:
            return 'a'

        num_ops = int(math.ceil(math.log2(k)))
        current_char = 0  # 'a' + 0 = 'a'

        while num_ops > 0:
            mid = 1 << (num_ops - 1)
            if k > mid:
                current_char += 1  # Right half → increment the character
                k -= mid           # Adjust position in the right half
            # else: k stays the same in the left half, char does not change
            num_ops -= 1

        return chr(ord('a') + (current_char % 26))
