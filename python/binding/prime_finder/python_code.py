def is_prime(num: int):
    #Check if a number is prime.
    if num < 2:
        return False
    for i in range(2, num):
        if num % i == 0:
            return False
    return True


def find_nth_prime(n: int):
    #Find the nth prime number.
    count = 0
    num = 1
    while count < n:
        num += 1
        if is_prime(num):
            count += 1
    return num
