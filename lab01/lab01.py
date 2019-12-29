import string
import random
from time import time
from memory_profiler import *

def LevRecursion(str1, str2, output = False):
    if str1 == '' or str2 == '':
        return abs(len(str1) - len(str2))
    if (str1[-1] == str2[-1]):
        penalty = 0
    else:
        penalty = 1
    return min(LevRecursion(str1, str2[:-1]) + 1,
               LevRecursion(str1[:-1], str2) + 1,
               LevRecursion(str1[:-1], str2[:-1]) + penalty)

#@profile(precision = 4)
def LevMatr(str1, str2, output = False):
    len_i = len(str1) + 1
    len_j = len(str2) + 1
    mtr = [[i + j for j in range(len_j)] for i in range(len_i)]

    for i in range(1, len_i):
        for j in range(1, len_j):
            if (str1[i-1] == str2[j-1]):
                penalty = 0
            else:
                penalty = 1
            mtr[i][j] = min(mtr[i-1][j] + 1,
                              mtr[i][j-1] + 1,
                              mtr[i-1][j-1] + penalty)
    if output:
        PrintMtr(mtr, str1, str2)
    return(mtr[-1][-1])

def DamLevRecursion(str1, str2, output = False):
    if str1 == "" or str2 == "":
        return abs(len(str1) - len(str2))
    if str1[-1] == str2[-1]:
        penalty = 0
    else:
        penalty = 1
    result = min(DamLevRecursion(str1, str2[:-1])+1,
                 DamLevRecursion(str1[:-1], str2)+1,
                 DamLevRecursion(str1[:-1], str2[:-1]) + penalty)
    if (len(str1) >= 2 and len(str2) >= 2 and str1[-1] == str2[-2]\
        and str1[-2] == str2[-1]):
        result = min(result, DamLevRecursion(str1[:-2], str2[:-2]) + 1)
    return result

def DamLevMtr(str1, str2, output = False):
    len_i = len(str1) + 1
    len_j = len(str2) + 1
    mtr = [[i+j for j in range (len_j)] for i in range (len_i)]

    for i in range(1, len_i):
        for j in range(1, len_j):
            if (str1[i-1] == str2[j-1]):
                penalty = 0
            else:
                penalty = 1
            mtr[i][j] = min(mtr[i-1][j] + 1,
                            mtr[i][j-1] + 1,
                            mtr[i-1][j-1] + penalty)
            if (i > 1 and j > 1) and \
               str1[i-1] == str2[j-2] and str1[i-2] == str2[j-1]:
                mtr[i][j] = min(mtr[i][j], mtr[i-2][j-2]+1)
    if output:
        PrintMtr(mtr, str1, str2)
    return mtr[-1][-1]

def GetStrAndRun(function, output = False):
    str1 = input("Первая строка: ")
    str2 = input("Вторая строка: ")
    res = function(str1, str2, output)
    print("Distance =", res)

def PrintMtr(table, str1, str2):
    print("\n   ", end = " ")
    for i in range( len(str2)):
        print(str2[i], end = " ")

    for i in range(len(table)):
        if i:
            print("\n" + str1[i-1], end = " ")
        else:
            print("\n ", end = " ")
        for j in range(len(table[i])):
            print(table[i][j], end = " ")
    print("\n")

def TimeAnalysis(function, num_iter, str_len = 10):
    t1 = time()
    for i in range(num_iter):
        str1 = gain_str(str_len)
        str2 = gain_str(str_len)
        function(str1, str2, False)
    t2 = time()
    return (t2 - t1) 

def gain_str(str_len = 10):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters)for i in range(str_len))

def main():
    flag = True
    while(flag):
        case = input("Меню:\n \
\t1. Расстояние Левенштейна рекурсия\n \
\t2. Расстояние Левенштейна матрица\n \
\t3. Расстояние Дамерау-Левенштейна рекурсия\n \
\t4. Расстояние Дамерау-Левенштейна рекурсия\n \
\t5. Сравнить все\n \
\t6. Временной анализ\n ")
        if (case == "1"):
            GetStrAndRun(LevRecursion, True)
        elif (case == "2"):
            GetStrAndRun(LevMatr, True)
        elif (case == "3"):
            GetStrAndRun(DamLevRecursion, True)
        elif (case == "4"):
            GetStrAndRun(DamLevTable, True)
        elif (case == "5"):
            output = True
            str1 = input("Первая строка: ")
            str2 = input("Вторая строка: ")
            print("Расстояние Левенштейна (рекурсия) = ", LevRecursion(str1, str2, output))
            print("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
            print("Расстояние Левенштейна (матрица) = ", LevMatr(str1, str2, output))
            print("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
            print("Расстояние Дамерау-Левенштейна (рекурсия)= ", DamLevRecursion(str1, str2, output))
            print("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
            print("Расстояние Дамерау-Левенштейна (матрица)= ", DamLevMtr(str1, str2, output))
        elif (case == "6"):
            nIter = 500
            for i in range(1, 8):
                print("Strlen: ", i)
                print("   Lev recursion   : ", "{0:.15f}".format(TimeAnalysis(LevRecursion, nIter, i)))
                print("   Lev table       : ", "{0:.15f}".format(TimeAnalysis(LevMatr, nIter, i)))
                print("   DamLev recursion: ", "{0:.15f}".format(TimeAnalysis(DamLevRecursion, nIter, i)))
                print("   DamLev table    : ", "{0:.15f}".format(TimeAnalysis(DamLevMtr, nIter, i)))

            
        else:
            flagDo = False

if __name__ == "__main__":
    LevRecursion(gain_str(10), gain_str(10))
