use std::fs;

fn parse_file(file_name: &str) -> (Vec<String>, Vec<usize>, Vec<usize>) {
    let mut columns: Vec<usize> = vec![];
    let mut rows: Vec<usize> = vec![];
    let mut counter: u16 = 0;
    let galaxy = String::from_utf8(fs::read(file_name).unwrap())
    .unwrap()
    .split("\n")
    .into_iter()
    .enumerate()
    .map(|(row, s)| {
       let mut contains_no_galaxies: bool = true;
        let s: String = s.chars().enumerate().map(|(col, s)| {
        if s == '#' { 
            if !columns.contains(&col) {
                columns.push(col);
                
            }
            contains_no_galaxies = false;
        }

        s.to_string()
        }).collect();

        if contains_no_galaxies {
            rows.push(row);
        }

        s
    }).collect();

(galaxy, columns, rows)
}

fn print(content: &Vec<String>) {
    content.iter().for_each(|s| {
        s.chars().for_each(|c| {
    
            print!("{}", c);
        
        });
        println!();
    });
}

fn get_non_cols(cols: Vec<usize>, size_of_line: usize) -> Vec<usize> {
    let mut other_cols: Vec<usize> = vec![];

    for num in 0..size_of_line {
        if !cols.contains(&num) {
            other_cols.push(num);
        }
    }

    other_cols
}

fn widen_vertically(content: &mut Vec<String>, cols: Vec<usize>) {
    let cols = get_non_cols(cols, content[0].len()); 

    let length = content[0].len();
    for vertical in 0..length {
        let vertical_backwards = length - 1 - vertical;
        
        for add_space in cols.iter() {
            if vertical_backwards == *add_space {
                for line in content.iter_mut() {
                    line.insert(vertical_backwards, '.'); 
                }
            }
        }
    }
}

fn fill_galaxies(content: Vec<String>) -> Vec<Galaxy> {
    let mut all_galaxies: Vec<Galaxy>= vec![];
    content
    .iter()
    .enumerate()
    .for_each(|(row, line)| {
        line
        .chars()
        .enumerate()
        .for_each(|(col, c)| {
            if c != '.' {
            all_galaxies.push(Galaxy{
                position: (row as i32, col as i32),
                name: String::from(c),    
            })
            } 
        })
    });

 all_galaxies
}

fn calculate_shortest_path(all_galaxies: Vec<Galaxy>, cols: Vec<usize>, rows: Vec<usize>) -> Vec<i32>{
let mut all_shortest_paths: Vec<i32> = vec![];

for current_galaxy in 0..all_galaxies.len() {
    for num in current_galaxy..all_galaxies.len() {
        if num == current_galaxy {
            continue;
        }

        let gx = all_galaxies[current_galaxy].position.0;
        let gy = all_galaxies[current_galaxy].position.1;

        let g1x = all_galaxies[num].position.0;
        let g1y = all_galaxies[num].position.1;

        let y = g1y - gy;
        let x = g1x - gx;

        let range2;
        let range1 = if gy > g1y {
            gy
        } else {
            
        };

        for num_between in gy..g1y {

        }

        let distance = y + x;
        all_shortest_paths.push(distance); 
    }
}

all_shortest_paths
}

#[derive(Debug, Clone)]
struct Galaxy {
    position: (i32, i32),
    name: String,
}

fn main() {
    let (mut content, cols, rows) = parse_file("adventday11.txt"); 
    widen_vertically(&mut content, cols);
    let all_galaxies = fill_galaxies(content);
    let all_values =  calculate_shortest_path(all_galaxies, cols, rows);
    let total: i32 = all_values.iter().sum();

    println!("total: {}", total);
}

#[cfg(test)]
mod test {
    use crate::{parse_file, get_non_cols, widen_vertically, fill_galaxies, calculate_shortest_path, Galaxy, print};

    #[test]
    fn test_parse_file() {
        let (contents, cols, rows) = parse_file("test.txt");

        println!("Cols: {:?}", cols);
        assert_eq!(contents[0].len(), 10);
        assert_eq!(contents[0].len(), 10);
   }

   #[test]
   fn test_cols_with_no_galaxies() {
        let (content, cols, rows) = parse_file("test.txt");

        let cols = get_non_cols(cols, content[0].len());

        assert_eq!(vec![2,5,8], cols);
   } 
   #[test]
   fn test_widening() {
        let (mut content, cols, rows) = parse_file("test.txt");

        widen_vertically(&mut content, cols);

        assert_eq!('1', content.get(0).unwrap().chars().nth(4).unwrap());
        assert_eq!('6', content.get(7).unwrap().chars().nth(12).unwrap());
      }

      #[test]
      fn test_all_galaxies() {
        let (mut content, cols, rows) = parse_file("test.txt");
        widen_vertically(&mut content, cols);

       let all_galaxies = fill_galaxies(content);
        let position_of_first = &all_galaxies[0];

       assert_eq!((0,4), position_of_first.position); 
      }

      fn setup(file_name: String) -> (Vec<i32>, Vec<Galaxy>) {
        let (mut content, cols, rows) = parse_file(&file_name);
        widen_vertically(&mut content, cols);
       let all_galaxies = fill_galaxies(content);

       println!("{:?}", all_galaxies);

        let all_paths = calculate_shortest_path(all_galaxies.clone());
      
        (all_paths, all_galaxies)
      }

      #[test]
      fn test_calculate_shortest_paths() {
        println!("test2");
        let (all_paths, _) = setup("test2.txt".into());
        assert_eq!(17, all_paths[0]);
        
        println!("test3");
        let (all_paths, _) = setup("test3.txt".into());
        assert_eq!(34, all_paths[0]);
        
        println!("test4");
        let (all_paths, _) = setup("test4.txt".into());
        assert_eq!(24, all_paths[0]);

        // let (all_paths, _) = setup("test3.txt".into());
        // assert_eq!(34, all_paths[0]);



    //      let (all_paths, galaxies) = setup("test2.txt".into());
    //     assert_eq!(17, all_paths[0]);

    //     let (all_paths, galaxies) = setup("test2.txt".into());
    //     assert_eq!(17, all_paths[0]);
      }


    }