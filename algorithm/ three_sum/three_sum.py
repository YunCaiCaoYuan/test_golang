def threeSum(nums):
    if len(nums) < 3:
        return []

    nums.sort()

    res = set()

    for i, v in enumerate(nums[:-2]):
        if i >= 1 and v == nums[i - 1]:
            continue
        d = {}
        for x in nums[i + 1:]:
            if x not in d:
                d[-v - x] = 1
            else:
                res.add((v, -v - x, x))
    # print(res)
    return map(list, res)


if __name__ == '__main__':
    nums = [-1, 0, 1, 2, -1, -4]
    res = threeSum(nums)
    print(list(res))

# [[-1, 0, 1], [-1, -1, 2]]
