pub struct MyStruct {
    pub x: i32,
    pub y: u8
}

pub trait MyData {
    fn print(&self);
}

impl MyData for MyStruct {
    fn print(&self) {
        println!("x: {} -> y: {}", self.x, self.y);
    }
}
