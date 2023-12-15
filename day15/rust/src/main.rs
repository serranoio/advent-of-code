use std::{collections::HashMap, fs};

// PART 1 ONE-LINER BABY
fn part_1(file_name: &str) -> u64 { String::from_utf8(fs::read(file_name).unwrap()).unwrap().replace("\n", "").split(",").map(|string| {string.chars().map(|c| c as u64).reduce(|acc, c| c + acc * 17 % 256).unwrap() * 17 % 256 }).sum::<u64>()}

fn part_2(file_name: &str) -> u64 {
    let mut maps: Vec<(u64, Vec<String>)> = vec![];

    for boxes in 0..256 {
        let vec: Vec<String> = vec![];
        maps.push((boxes, vec));
    }
    // to have the Vec's live as long as the hashmap :/
    let mut boxes: HashMap<u64, Vec<String>> = maps.into_iter().collect();

    String::from_utf8(fs::read(file_name).unwrap())
        .unwrap()
        .replace("\n", "")
        .split(",")
        .for_each(|string| {
            let delimiter;
            if string.contains('=') {
                delimiter = '=';
            } else {
                delimiter = '-';
            }

            let parts: Vec<&str> = string.split(delimiter).collect();

            let hash = parts[0]
                .chars()
                .map(|c| c as u64)
                .reduce(|acc, c| c + acc * 17 % 256)
                .unwrap()
                * 17
                % 256;

            if delimiter == '=' {
                let mut found = false;
                // if already in, replace
                boxes.get_mut(&hash).unwrap().iter_mut().for_each(|lenses| {
                    if (*lenses).contains(parts[0]) {
                        *lenses = parts[0].to_owned() + parts[1];
                        found = true;
                    }
                });
                
                if !found {
                    boxes
                        .get_mut(&hash)
                        .unwrap()
                        .push(parts[0].to_owned() + parts[1])
                }
            } else {  // '-' signifies we need to delete instances inside box
                let lenses = boxes.get_mut(&hash).unwrap();

                let mut lens_index = 0;
                loop {
                    if lens_index >= lenses.len() {
                        break;
                    }
                    if lenses[lenses.len() - 1 - lens_index].contains(parts[0]) {
                        lenses.remove(lenses.len() - 1 - lens_index);
                    }
                    lens_index += 1
                }
            };
        });

    let total = boxes
        .iter()
        .map(|(box_number, lenses)| {
            let box_number = box_number.clone();

            lenses
                .iter()
                .enumerate()
                .map(|(pos, lens)| {
                    let total_in_box = (box_number + 1)
                        * (pos as u64 + 1)
                        * lens.chars().last().unwrap().to_digit(10).unwrap() as u64;
                    total_in_box 
                })
                .sum::<u64>()
        })
        .sum::<u64>();

    total
}

fn main() {
    let answer: u64 = part_1("adventday16.txt");
    println!("Part 1: {}", answer);

    let answer = part_2("adventday16.txt");
    println!("Part 2: {}", answer);
}
