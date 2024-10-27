use std::vec;

use pyo3::prelude::pyfunction;

#[pyfunction]
pub fn demo_loops() {
    let mut mylist = vec![1, 2, 3, 4];

    // for el in mylist.into_iter() {
    //     println!("{}", el);
    // }

    for el in mylist.iter() {
        println!("{}", el);
    }

    // println!("Complex implementation");
    // let mut iter = IntoIterator::into_iter(mylist);
    // loop {
    //     match iter.next() {
    //         Some(x) => {
    //             println!("{}", x);
    //         },
    //         None => { break }
    //     }
    // }

    for item in mylist.iter_mut() {
        *item += 1;
        println!("{}", *item);
    }

    println!("=========== query using index ==========");
    println!("{}", mylist[0]);
}

#[pyfunction]
pub fn functional_iter() {
    let members = vec![1,2,3,4,5];
    let members2 = vec![1, 2, 3, 4, 5];
    let n = 3;
    let trans_members = members.iter()
    .map(|&x| (x as i32).pow(n));

    trans_members.for_each(|x| println!("{}", x));

    println!("========== filter ==============");
    let filtered_members = members
    .into_iter()
    .filter(|&x| x > 2);

    filtered_members.for_each(|x| println!("{}", x));

    let sum: i32 = members2.iter().sum();
    println!("sum: {}", sum);

    let max = members2.iter().max().unwrap();
    println!("max: {}", max);
}
