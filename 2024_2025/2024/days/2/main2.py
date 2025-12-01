from pprint import pprint

with open("input.txt", "r") as f:
    inp = f.read()

# increasing OR decreasing
# diff by at least one, at most three

all_nums = []
for line in inp.split("\n"):
    nums = list(map(int, line.split()))
    all_nums.append(nums)

pprint(all_nums)
# check list of ints
n_safe = 0

def all_increasing(nums: list[int]) -> bool:
    return all(nums[i] < nums[i+1] for i in range(len(nums) - 1))

def all_decreasing(nums: list[int]) -> bool:
    return all(nums[i] > nums[i+1] for i in range(len(nums) - 1))

def max_diff(nums: list[int]) -> int:
    return max(abs(nums[i] - nums[i+1]) for i in range(len(nums) - 1))

def is_safe(nums: list[int]) -> bool:
    increasing = all_increasing(nums)
    decreasing = all_decreasing(nums)
    md = max_diff(nums)
    return (increasing or decreasing) and md >= 1 and md <= 3

# for each list of integers in all_nums, check if the list can be safe if
# a single integer is removed from the list

for nums in all_nums:
    if is_safe(nums):
        n_safe += 1
    else:
        # check if any single integer can be removed to make the list safe
        for i in range(len(nums)):
            if is_safe(nums[:i] + nums[i+1:]):
                n_safe += 1
                break
       
print(n_safe)