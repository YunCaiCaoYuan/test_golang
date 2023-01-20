def binary_search(array, target):
    left, right = 0, len(array) - 1
    while left <= right:
        mid = int(left + (right - left) / 2)
        if array[mid] == target:
            return mid
        elif array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1


if __name__ == '__main__':
    array = [3, 1, 4, 15, 9, 2, 6, 53]
    target = 53
    ret = binary_search(array, target)
    print(ret)
