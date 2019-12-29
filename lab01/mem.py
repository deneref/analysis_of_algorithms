from memory_profiler import *
from lab01 import *

@profile(precision = 10)
def mem_usage(size):
    LevMatr(gain_str(size), gain_str(size))
    LevRecursion(gain_str(size), gain_str(size))
    
if __name__ == "__main__":
    mem_usage(30)
