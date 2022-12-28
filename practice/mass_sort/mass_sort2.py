#!/usr/bin/python2.7
# -*- coding: utf-8 -*-

import os
import time

testdir = '/ssd/subfile'

now = time.time()

# Step 1 : 获取全部文件描述符
fds = []
for f in os.listdir(testdir):
    ff = os.path.join(testdir,f)
    fds.append(open(ff,'r'))

# Step 2 : 每个文件获取第一行，即当前文件最小值
nums = []
tmp_nums = []
for fd in fds:
    num = int(fd.readline())
    tmp_nums.append(num)

# Step 3 : 获取当前最小值放入暂存区，并读取对应文件的下一行；循环遍历。
count = 0
while 1:
    val = min(tmp_nums)
    nums.append(val)
    idx = tmp_nums.index(val)
    next = fds[idx].readline()
    # 文件读完了
    if not next:
        del fds[idx]
        del tmp_nums[idx]
    else:
        tmp_nums[idx] = int(next)
    # 暂存区保存1000个数，一次性写入硬盘，然后清空继续读。
    if 1000 == len(nums):
        with open('final_sorted.txt','a') as wf:
            wf.write(''.join([ str(i) for i in nums ]) + '')
            nums[:] = []
            if 499999999 == count:
            break
    count += 1

with open('runtime.txt','w') as wf:
    wf.write('Runtime : {}'.format(time.time()-now))