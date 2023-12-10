use std::{path::PathBuf, fs::{File, self}, io::{BufReader, Read}, time::SystemTime};

fn get_all_lines(file_name: String) -> Vec<String> {
    let path = PathBuf::new().join("../").join(&file_name);

    
    
    let contents = String::from_utf8(fs::read(path).unwrap()).unwrap();

    contents.split("\n").map(|s| s.to_owned()).collect()
}

fn create_line_of_ints(line: &str) -> Vec<i32> {
    let mut line_of_ints: Vec<i32> = Vec::new();

    line.split(" ").for_each(|num| {
        let num: i32 = num.parse().unwrap();
        line_of_ints.push(num);          
    });

    line_of_ints
}

fn has_all_zeros(line: &Vec<i32>) -> bool {
    for num in line {
        if num != &0 { 
            return false;
        }
    }

    return true;
}

fn gather_differences(line: Vec<i32>) -> Vec<i32> {
    let mut differences: Vec<i32> = Vec::new();

    let mut current = 0;
    let mut next = 1;

    while next+1 <= line.len() {
        let difference = line[next] - line[current];

        differences.push(difference);


        current = next;
        next+=1;
    }

    differences
}

fn calculate_differences(line_of_ints: Vec<i32>) -> Vec<Vec<i32>> {
    let mut all_differences: Vec<Vec<i32>> = Vec::new();
    all_differences.push(line_of_ints.clone());

    let mut next_line = gather_differences(line_of_ints);

    while !has_all_zeros(&next_line) {
        all_differences.push(next_line.clone());

        next_line = gather_differences(next_line);        
    };


    all_differences.push(next_line);

    all_differences
}

fn calculate_history(line_of_ints: Vec<i32>) -> i32 {
    let differences = calculate_differences(line_of_ints);


    let mut last_ones: Vec<i32> = Vec::new();

    differences.iter().for_each(|difference_line: &Vec<i32>| {
        last_ones.push(difference_line[difference_line.len()-1]);
    });


    last_ones.iter().sum()
}

fn calculate_history2(line_of_ints: Vec<i32>) -> i32 {
    let differences = calculate_differences(line_of_ints);


    let mut last_ones: Vec<i32> = Vec::new();

    differences.iter().for_each(|difference_line: &Vec<i32>| {
        last_ones.push(difference_line[0]);
    });


    last_ones.into_iter().rev().reduce(|acc, x| x - acc).unwrap()
}


fn to_double_dimension_array(all_lines: &Vec<String>, part_one: bool) -> i32 {
    let mut all_histories: Vec<i32> = Vec::new();

    all_lines.into_iter().for_each(|line| {
        let line_of_ints: Vec<i32> = create_line_of_ints(line);

        let history: i32;

        if part_one {
            history = calculate_history(line_of_ints);
        } else {
            history = calculate_history2(line_of_ints);
        }

        all_histories.push(history);
    });

    all_histories.iter().sum()
} 

fn main() {
    let now = SystemTime::now();
    let all_lines = get_all_lines("adventday9.txt".to_owned());
    
    let all = to_double_dimension_array(&all_lines, true);
    println!("Sum for part 1: {all}");

    let all = to_double_dimension_array(&all_lines, false);
    println!("Sum for part 2: {all}");

    println!("Runtime: {} nanoseconds\n", now.elapsed().unwrap().as_nanos());
}

#[cfg(test)]
mod test {
    use crate::{get_all_lines, gather_differences, calculate_differences, calculate_history, calculate_history2};

#[test]
fn test_all_lines() {
    let all_lines = get_all_lines("test.txt".to_owned());

    assert_eq!(3, all_lines.len());
}

#[test]
fn test_gather_differences() {
    let line: Vec<i32> = vec![0,3,6,9,12,15];

    let differences: Vec<i32> = gather_differences(line);
    assert_eq!(vec![3, 3, 3, 3, 3], differences);


    let differences = gather_differences(differences);
    assert_eq!(vec![0,0,0,0], differences);
}

#[test]
fn test_calculate_differences() {
    let start: Vec<i32> = vec![0,3,6,9,12,15];

    let differences = calculate_differences(start);

    assert_eq!(vec![3,3,3,3,3], differences[1]);
    assert_eq!(vec![0,0,0,0], differences[2]);
}

#[test]
fn test_calculate_history() {
    let start: Vec<i32> = vec![0,3,6,9,12,15];

    let history = calculate_history(start);
    assert_eq!(18, history);

    let history = calculate_history(vec![1,3,6,10,15,21]);
    assert_eq!(28, history);

    let history = calculate_history(vec![10,13,16,21,30,45]);
    assert_eq!(68, history);
}

#[test]
fn test_calculate_history2() {
    let start: Vec<i32> = vec![0,3,6,9,12,15];

    let history = calculate_history2(start);
    assert_eq!(-3, history);

    let history = calculate_history2(vec![1,3,6,10,15,21]);
    assert_eq!(0, history);

    let history = calculate_history2(vec![10,13,16,21,30,45]);
    assert_eq!(5, history);
}



}