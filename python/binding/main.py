from timeit import timeit
import sys

from prime_finder import find_nth_prime
from find_prime_rs.find_prime_rs import find_nth_prime_rust

ITERATIONS = 100


def main():
    n = int(sys.argv[1])

    print(f"\nUsing Python => The {n}-th prime number is: {find_nth_prime(n)}")
    print(f"Using Rust => The {n}-th prime number is: {find_nth_prime_rust(n)}")

    python_time_per_iter = timeit(
        lambda: find_nth_prime(n), number=ITERATIONS) / ITERATIONS

    rust_time_per_iter = timeit(
        lambda: find_nth_prime_rust(n), number=ITERATIONS) / ITERATIONS

    print(
        f"\nTime taken using Python Prime finder: {python_time_per_iter* 1000:.2f} milliseconds per iteration.")

    print(
        f"Time taken using Rust Prime finder: {rust_time_per_iter * 1000:.2f} milliseconds per iteration.")

    if python_time_per_iter > rust_time_per_iter:
        performance = (
            (python_time_per_iter - rust_time_per_iter) / rust_time_per_iter) * 100
        print(f"\nRust code is faster than Python by {performance:.2f}%\n")
    else:
        performance = (
            (rust_time_per_iter - python_time_per_iter) / python_time_per_iter) * 100
        print(f"\nPython code is faster than Rust by {performance:.2f}%\n")


if __name__ == "__main__":
    main()
