use pyo3::prelude::*;

#[pyclass]
pub struct MyRustClass {
    value: i32,
}

#[pymethods]
impl MyRustClass {
    #[new]
    fn new(value: i32) -> Self {
        MyRustClass {value}
    }

    fn get_value(&self) -> i32 {
        self.value
    }
}