use std::{fs, env};

use geo::point;
use polygonical::{point::Point, polygon::Polygon};

fn parse_file(file_name: &str)  {
    
    let mut current_points = Point::new(0., 0.);

    let mut points: Vec<Point> = vec![current_points];

    let mut length = 0.;

    String::from_utf8(fs::read(file_name).unwrap())
    .unwrap()
    .split("\n")
    .for_each(|line| {
        let mut parts = line.split(" ");

        let direction = parts.next().unwrap();
        let amount = parts.next().unwrap();

        let first = current_points.clone(); 

        let amount = amount.parse::<f64>().unwrap();
        if direction == "R" {
            current_points.x += amount;
        } else if direction == "L" {
            current_points.x -= amount;
        } else if direction == "D" {
            current_points.y += amount; 
        } else if direction == "U" {
            current_points.y -= amount;
        }

        length += (first.y -current_points.y).abs() + (first.x - current_points.x).abs();

        points.push(current_points);

    });

    let polygon = polygonical::polygon::Polygon::new(points);

    let area = polygon.area();

    println!("Area: {}", area.abs() + (length/2.) + 1.);
    
}

fn parse_file_2(file_name: &str) {
    
    let mut current_points = Point::new(0., 0.);

    let mut points: Vec<Point> = vec![current_points];

    let mut length = 0.;
    
    String::from_utf8(fs::read(file_name).unwrap())
    .unwrap()
    .split("\n")
    .for_each(|line| {
        let mut parts = line.split(" ");        

        let hex = String::from(parts.last().unwrap()); 
        let direction = hex.chars().nth(hex.len()-2).unwrap(); 
        let amount = i64::from_str_radix(&hex[2..7], 16).unwrap();
        let amount: f64 = amount as f64;

        let first = current_points.clone(); 

        if direction == '0' {
            current_points.x += amount;
        } else if direction == '2' {
            current_points.x -= amount;
        } else if direction == '1' {
            current_points.y += amount; 
        } else if direction == '3' {
            current_points.y -= amount;
        }

        length += (first.y -current_points.y).abs() + (first.x - current_points.x).abs();

        points.push(current_points);
        
    });
    let polygon = polygonical::polygon::Polygon::new(points);

    let area = polygon.area();

    println!("Area: {}", area.abs() + (length/2.) + 1.);
    
}

fn main() {

    let args: Vec<String> = env::args().collect();

    parse_file(&args[1]);
    parse_file_2(&args[1]);
}
