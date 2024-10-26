use pyo3::prelude::*;

// Helper function in Rust
#[pyfunction]
fn is_prime_rust(num: u32) -> bool {
    if num < 2 {
        return false;
    }

    for i in 2..num {
        if num % i == 0 {
            return false;
        }
    }
    
    true
}

// Nth Prime finding function in Rust
#[pyfunction]
fn find_nth_prime_rust(n: u32) -> u32 {
    let mut count: u32 = 0;
    let mut num: u32 = 1;

    while count < n {
        num += 1;

        if is_prime_rust(num) {
            count += 1;
        }
    }

    num
}

/// Python module implementation.
#[pymodule]
fn find_prime_rs(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(find_nth_prime_rust, m)?)?;
    Ok(())
}
