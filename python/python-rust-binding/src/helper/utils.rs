use std::collections::HashMap;

pub fn say_hello() {
    println!("helper world!")
}

const FIB_0: u64 = 1;
const FIB_1: u64 = 1;

pub fn fib(n: u64) -> u64 {
    if n == 0 {
        FIB_0
    } else if n == 1 {
        FIB_1
    } else {
        fib(n - 1) + fib(n - 2)
    }
}

pub fn fib_dyn(n: u64, map: &mut HashMap<u64, u64>) -> u64 {
    match n {
        0 | 1 => 1,
        n => {
            if map.contains_key(&n) {
                *map.get(&n).unwrap()
            } else {
                let val = fib_dyn(n - 1, map) + fib_dyn(n - 2, map);
                map.insert(n, val);
                val
            }
        }
    }
}
