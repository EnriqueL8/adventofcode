use math::round;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

static INPUT_FILE: &str = "../input.txt";

// Calcualte fuel for a mass
fn calculate_fuel(mass: i32) -> i32 {
    let fuel = round::floor(f64::from(mass / 3), 0) as i32 - 2;
    if fuel < 0 {
        return 0
    }

    fuel
}


// Calculate the total fuel for one mass
fn total_fuel_for_mass(initial_mass: i32) -> i32 {
    let mut total = 0;
    let mut fuel = initial_mass;
    while fuel > 0 {
        fuel = calculate_fuel(fuel);
        total += fuel;
    }

    total
}


// Read file and calculate fuel for all the modules
fn main() -> io::Result<()> {
    let file = File::open(INPUT_FILE)?;
    let reader = BufReader::new(file);

    let mut total = 0;

    for line in reader.lines() {
        let my_string = line.unwrap().to_string();
        let my_int: i32 = my_string.parse().unwrap();
        total += total_fuel_for_mass(my_int)
    }

    println!("Total fuel needed is: {}", total);

    Ok(())
}
