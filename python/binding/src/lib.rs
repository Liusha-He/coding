pub mod find_prime_funcs;
pub mod my_classes;
pub mod syntax;

use pyo3::prelude::*;
use find_prime_funcs::functions::find_nth_prime_rust;
use my_classes::rust_class::MyRustClass;
use syntax::loops::*;

#[pymodule]
fn libext_rs(_py: Python, m: &PyModule) -> PyResult<()> {
    // register functions here
    m.add_function(wrap_pyfunction!(find_nth_prime_rust, m)?)?;
    m.add_function(wrap_pyfunction!(demo_loops, m)?)?;
    m.add_function(wrap_pyfunction!(functional_iter, m)?)?;

    // register classes here
    m.add_class::<MyRustClass>()?;

    Ok(())
}
