extern crate rand;
extern crate time;

mod helper;

use std::collections::{HashMap, LinkedList};
use std::fmt::Display;
use std::io;
use rand::prelude::Rng;
use helper::utils::say_hello;
use helper::utils::fib;
use helper::utils::fib_dyn;
use helper::data_types::MyStruct;
use helper::data_types::MyData;

fn get_guess() -> u8 {
    loop {
        let mut guess: String = String::new();
        io::stdin().read_line(&mut guess).expect("Could not read from terminal!");

        match guess.trim().parse::<u8>() {
            Ok(v) => return v,
            Err(e) => println!("Could not understand input {}", e),
        }
    }
}

fn handle_guess(expected: u8) {
    for n in 0..7 {
        let ans = get_guess();
        if ans == expected {
            println!("Yep, {} is right!", ans);
            break
        } else {
            if n == 6 {
                println!("Oh, no luck today!");
                println!("haha, the correct answer is {}", expected);
            } else {
                if ans > expected {
                    println!("Opps, too high, try again...");
                } else {
                    println!("Opps, too low, try again...")
                }
            }
        }
    }
}

fn show_all(v: Vec<&dyn Display>) {
    for item in v {
        println!("{}", item)
    }
}

fn main() {
    let mut ll = LinkedList::new();

    ll.push_back(1);
    ll.push_back(2);
    ll.push_back(3);

    for item in ll {
        println!("{item}");
    }
    println!("========================================");
    say_hello();

    println!("========================================");
    println!("{}", fib(13));

    println!("========================================");
    let mut map = HashMap::new();
    for n in 1..20 {
        println!("{} -> {}", n, fib_dyn(n, &mut map))
    }

    println!("========================================");
    println!("Guess Game!!!");
    println!("Input guess:");
    let mut rng = rand::thread_rng();
    let expected: u8 = rng.gen_range(1..20);
    handle_guess(expected);

    println!("========================================");
    let mut a = 100;
    let b = &mut a;
    *b += 1;

    println!("{}", a);

    println!("========================================");
    let d = MyStruct{x: 10, y: 5};
    d.print();

    println!("========================================");
    let v = vec![&12 as &dyn Display, &"hello" as &dyn Display];
    show_all(v);

//     closure - todo work out how this work
//     let myclosure = |x| { x * 2 };
//     println!("closure result -> {}", myclosure(32));

    let items = (1..10).into_iter();
    let other_items: Vec<i32> = items.map(
        |x| {
            x + 1
        }
    ).collect();

    for item in other_items {
        println!("{}", item)
    }
}
