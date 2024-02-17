import sys
import os
sys.path.insert(0, os.path.join(os.getcwd(), 'ten'))
from ten import contains_duplicate, valid_anagram




if __name__ == "__main__":
    print(contains_duplicate([1, 1]))
    print(valid_anagram("abc", "cba"))