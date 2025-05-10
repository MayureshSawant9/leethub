class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:

        def single_print(number: int) -> str:
            return f'{number}'

        def double_print(num1, num2: int) -> str:
            return f'{num1}->{num2}' 

        if len(nums) == 0:
            return []
        
        elif len(nums) == 1:
            return [str(nums[0])]

        elif len(nums) == 2:
            if nums[1] == nums[0] + 1:
                return [f'{nums[0]}->{nums[1]}']
            else:
                return [f'{nums[0]}', f'{nums[1]}']

        result = []

        start, i = 0, 0

        while True:

            if nums[i+1] == nums[i] + 1:
                
                i += 1

            else:

                if i == start:
                    result.append(single_print(nums[start]))
                else:
                    result.append(double_print(nums[start], nums[i]))
                    
                i += 1
                start = i
            
            if i == len(nums) - 1:
                if i == start:
                    result.append(single_print(nums[start]))
                else:
                    result.append(double_print(nums[start], nums[i]))

                return result
                

