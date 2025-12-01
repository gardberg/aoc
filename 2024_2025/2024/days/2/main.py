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
    print(increasing)
    decreasing = all_decreasing(nums)
    print(decreasing)
    md = max_diff(nums)
    print(md)
    return (increasing or decreasing) and md >= 1 and md <= 3

for nums in all_nums:
    if is_safe(nums):
        print(nums, "safe")
        n_safe += 1
    else:
        print(nums, "unsafe")
       
print(n_safe)