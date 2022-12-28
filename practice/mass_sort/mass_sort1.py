#!/usr/bin/python2.7

import time


def readline_by_yield(bfile):
    with open(bfile, 'r') as rf:
        for line in rf:
            yield line


def quick_sort(lst):
    if len(lst) < 2:
        return lst
    pivot = lst[0]
    left = [ele for ele in lst[1:] if ele < pivot]
    right = [ele for ele in lst[1:] if ele >= pivot]
    return quick_sort(left) + [pivot, ] + quick_sort(right)


def split_bfile(bfile):
    count = 0
    nums = []
    for line in readline_by_yield(bfile):
        num = int(line)
        if num not in nums:
            nums.append(num)
        if 10000 == len(nums):
            nums = quick_sort(nums)
            with open('subfile/subfile{}.txt'.format(count + 1), 'w') as wf:
                wf.write(''.join([str(i) for i in nums]))
            nums[:] = []
            count += 1
            print(count)


now = time.time()
split_bfile('total.txt')
run_t = time.time() - now
print('Runtime : {}'.format(run_t))
