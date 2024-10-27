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
