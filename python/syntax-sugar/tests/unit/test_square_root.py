from unittest import TestCase

from leetcode.square_root import Solution


class TestSolution(TestCase):
    def test_my_sqrt(self):
        s = Solution()
        self.assertEqual(2, s.mySqrt(4))
        self.assertEqual(2, s.mySqrt(8))
