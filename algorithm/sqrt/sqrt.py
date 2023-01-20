# 牛顿迭代法求平方根
# X    = ( X  + n/X  )/2
#  n+1      n      n

def sqrt(x):
    r = x
    while r * r > x:
        print("r:", r)
        r = (r + x / r) / 2
    return r


if __name__ == '__main__':
    # ret = sqrt(4)
    # print(ret)

    ret = sqrt(8)
    print(ret)
