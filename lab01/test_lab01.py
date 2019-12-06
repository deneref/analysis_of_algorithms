import unittest

from lab01 import LevRecursion, LevMatr, DamLevRecursion, DamLevMtr
from lab01 import gain_str

class TestDistanse(unittest.TestCase):
        
    def testEmpty(self):
        self.assertEqual(self.function("", ""), 0)

    def testSame(self):
        self.assertEqual(self.function("abc", "abc"), 0)
        self.assertEqual(self.function("0", "0"), 0)

    def testDifferent(self):
        self.assertEqual(self.function("a", ""), 1)
        self.assertEqual(self.function("", "1"), 1)
        self.assertEqual(self.function("b", "c"), 1)
        self.assertEqual(self.function("bc", "b"), 1)
        self.assertEqual(self.function("bc", "c"), 1)
        self.assertEqual(self.function("ab", "cd"), 2)


class TestLevDistanse(TestDistanse):
    def setUp(self):
        self.function = LevMatr
    def testTypo(self):
        self.assertEqual(self.function("ac", "ca"), 2)
        self.assertEqual(self.function("abc", "cba"), 2)

                         
class TestDamLevDistanse(TestDistanse):
    def setUp(self):
        self.function = DamLevMtr
    def testTypo(self):
        self.assertEqual(self.function("ac", "ca"), 1)
        self.assertEqual(self.function("abc", "cba"), 2)
                         

class TestTwoFunctions(unittest.TestCase):

    n = 15
    def testCompareSameLen(self):
        for i in range(TestTwoFunctions.n):
            str1 = gain_str(5)
            str2 = gain_str(5)
            self.assertEqual(self.f1(str1, str2), self.f2(str1, str2))

    def testCompareDifLen(self):
        for i in range(TestTwoFunctions.n):
            str1 = gain_str(3)
            str2 = gain_str(5)
            self.assertEqual(self.f1(str1, str2), self.f2(str1, str2))

    def testCompareEmpty(self):
        for i in range(TestTwoFunctions.n):
            str1 = gain_str(4)
            str2 = gain_str(0)
            self.assertEqual(self.f1(str1, str2), self.f2(str1, str2))
        

class TestLev(TestTwoFunctions):
    def setUp(self):
        self.f1 = LevRecursion
        self.f2 = LevMatr

        
class TestDamLev(TestTwoFunctions):
    def setUp(self):
        self.f1 = DamLevRecursion
        self.f2 = DamLevMtr
    

if __name__ == '__main__':
    suite = unittest.TestLoader().loadTestsFromTestCase(TestLev)
    suite.addTests(unittest.TestLoader().loadTestsFromTestCase(TestDamLev))
    suite.addTests(unittest.TestLoader().loadTestsFromTestCase(TestDamLevDistanse))
    suite.addTests(unittest.TestLoader().loadTestsFromTestCase(TestLevDistanse))
    unittest.TextTestRunner().run(suite)
    #unittest.main()
