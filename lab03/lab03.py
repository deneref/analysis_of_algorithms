from time import time
import sys
import numpy.random as rand
def bubbleSort(a):
    for i in range(len(a)):
        for j in range(len(a)-1, 0, -1):
            if a[j-1] > a[j]:
                a[j-1], a[j] = a[j], a[j-1]

    return a

def insertSort(a):
    for i in range(1, len(a)):
        #print(i, ' ')
        key = a[i]
        j = i-1
        while (j >= 0 and a[j] > key):
            a[j+1] = a[j]
            j -= 1
        a[j+1] = key
    return a

def partition(a,low, high):
    i = low - 1
    pivot = a[high]
    for j in range(low, high):
        if a[j] <= pivot:
            i += 1
            a[i], a[j] = a[j], a[i]
    a[i+1],a[high] = a[high], a[i+1]
    return (i+1)

def quickSort(a, low, high):
    if low < high:
        pi = partition(a, low, high)
        quickSort(a, low,pi-1)
        quickSort(a, pi+1, high)
    return a
    
def get_time(a, function, quick = False):
    if quick:
        t1 = time()
        function(a, 0, len(a)-1)
        t2 = time()
    else:
        t1 = time()
        function(a)
        t2 = time()
        
    return (t2 - t1)
def estimateTime():
    size = int(input("Введите размер массива\n"))
    t = []
    a = [i for i in range(size)]
    b = a.copy(); c = a.copy()
    t.append(get_time(a, bubbleSort))
    t.append(get_time(b, insertSort))
    t.append(get_time(c, quickSort, True))
    print('Отсортированный массив: \n',t)
    a.reverse(); b.reverse(); c.reverse(); t = []

    t.append(get_time(a, bubbleSort))
    t.append(get_time(b, insertSort))
    t.append(get_time(c, quickSort, True))
    print('Отсортированный в отбр. порядке массив: \n',t)

    a = rand.randint(0,100,size); b = rand.randint(0,100,size);
    c = rand.randint(0,100,size); t = []
    t.append(get_time(a, bubbleSort))
    t.append(get_time(b, insertSort))
    t.append(get_time(c, quickSort, True))
    print('Случайный массив: \n',t)

def input_and_run(random = False):
    size = int(input("Введите размер массива\n"))
    a = []
    for i in range(size):
        a.append(int(input("a[{}] = ".format(i))))
    b = a; c = a
    print("Input array {}\n".format(a))
    print("Сортировка пузырьком: {}\n".format(bubbleSort(a)))
    print("Сортировка вставками: {}\n".format(insertSort(b)))
    print("Быстрая сортировка: {}\n".format(quickSort(c, 0, size-2)))
    

def main():
    #print(sys.getrecursionlimit())
    flag = True
    while(flag):
        case = input("Меню:\n \
\t1. Сортировка\n \
\t2. Временой анализ\n ")
        if (case == "1"):
            input_and_run()
        elif (case == "2"):
            estimateTime()
        else:
            flagDo = False

if __name__ == "__main__":
    main()
