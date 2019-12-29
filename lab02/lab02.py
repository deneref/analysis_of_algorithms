import numpy as np
from random import randint
from time import time
import logging

def std(mtr1, mtr2):
    if len(mtr2) != len(mtr1[0]):
        print("Wrong size of matrix")
        return

    row1 = len(mtr1); col1 = len(mtr1[0])
    col2 =  len(mtr2[0])

    res = [[0 for i in range(col2)] for j in range(row1)]
    for i in range(row1):
        for j in range(col1):
            for k in range(col2):
                res[i][k] += mtr1[i][j] * mtr2[j][k]
    return res

def imp_std(mtr1, mtr2):
    if len(mtr2) != len(mtr1[0]):
        print("Wrong size of matrix")
        return
    return

def winograd(mtr1, mtr2):
    row1 = len(mtr1)
    row2 = len(mtr2)
    col2 = len(mtr2[0])

    if row2 != len(mtr1[0]):
        print("Different dimension of the matrics")
        return
    d = row2 // 2
    row_factor = [0 for i in range(row1)]
    col_factor = [0 for i in range(col2)]

    for i in range(row1):
        for j in range(d):
            row_factor[i] += mtr1[i][2 * j] * mtr1[i][2 * j + 1]
            
    for i in range(col2):
        for j in range(d):
            col_factor[i] += mtr2[2 * j][i] * mtr2[2 * j + 1][i]

    answer = [[0 for i in range(col2)] for j in range(row1)]
    for i in range(row1):
        for j in range(col2):
            answer[i][j] = - row_factor[i] - col_factor[j]
            for k in range(d):
                answer[i][j] += ((mtr1[i][2 * k] + mtr2[2 * k + 1][j]) *\
                                 (mtr1[i][2 * k + 1] + mtr2[2 * k][j]))

    if row2 % 2:
        for i in range(row1):
            for j in range(col2):
                answer[i][j] += mtr1[i][row2 - 1] * mtr2[row2 - 1][j]

    return answer

def imp_winograd(mtr1, mtr2):
    row1 = len(mtr1)
    row2 = len(mtr2)
    col2 = len(mtr2[0])

    if row2 != len(mtr1[0]):
        print("Different dimension of the matrics")
        return

    d = row2 // 2

    row_factor = [0 for i in range(row1)]
    col_factor = [0 for i in range(col2)]

    for i in range(row1):
        row_factor[i] = sum(mtr1[i][2 * j] * mtr1[i][2 * j + 1] for j in range(d))

    for i in range(col2):
        col_factor[i] = sum(mtr2[2 * j][i] * mtr2[2 * j + 1][i] for j in range(d))

    answer = [[0 for i in range(col2)] for j in range(row1)]
    for i in range(row1):
        for j in range(col2):
            answer[i][j] = sum((mtr1[i][2 * k] + mtr2[2 * k + 1][j]) * (mtr1[i][2 * k + 1] + mtr2[2 * k][j]) for k in range(d))\
                           - row_factor[i] - col_factor[j]

    if row2 % 2:
        for i in range(row1):
            answer[i][j] = sum(mtr1[i][row2 - 1] * mtr2[row2 - 1][j] for j in range(col2))

    return answer
def get_mtr(n, m):
    return [[randint(-100,100) for i in range(n)] for j in range(m)]

def get_time(size, function):
##    mtr1 = get_mtr(size, size)
##    mtr2 = get_mtr(size, size)
    t1 = time()
    function(get_mtr(size, size), get_mtr(size, size))
    t2 = time()
    return (t2-t1)

def estimate_time():
    for i in range(100, 800, 100):
        print('{} std = {}; win = {}; imp_win = {}'.format(i, round(get_time(i, std),5),\
                                                           round(get_time(i, winograd),5),\
                                                           round(get_time(i, imp_winograd),5)))

def main():
##    A = [[1, 2, 3, 4], [2, 3, 4, 5], [3, 4, 5, 6]]
##    B = [[1, 2, 3, 4, 5], [2, 3, 4, 5, 6], [3, 4, 5, 6, 7], [4, 5, 6, 7, 8]]
##    print('A = {}, \n B = {}\n'.format(A, B))
##    print('Стандартный алгоритм умножения A*B {}\n '.format(std(A, B)))
##    print('Алгоритм Винограда A*B {}\n '.format(winograd(A, B)))
##    print('Отпимизированный алгоритм Винограда A*B {}\n '.format(imp_winograd(A, B)))
    estimate_time()
    
if __name__ == '__main__':
    main()
