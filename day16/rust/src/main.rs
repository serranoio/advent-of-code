use std::{fs, env};

#[derive(PartialEq, Debug)]
enum SpaceType {
    ForwardMirror,
    BackwardMirror,
    VerticalSplitter,
    HorizontalSplitter,
    EmptySpace,
    False
}

#[derive(PartialEq, Debug)]
enum Direction {
    North,
    South,
    East,
    West,
    Uninitialized,
}

#[derive(Debug)]
struct GridSpace {
    space_type: SpaceType,
    is_energized: bool,
    direction: Vec<Direction>,
}

impl GridSpace {
    fn new(c: char) -> GridSpace {
        let space_type: SpaceType;
        if c == '.' {
            space_type = SpaceType::EmptySpace;
        } else if c == '/' {
            space_type = SpaceType::ForwardMirror;
        } else if c == '\\' {
            space_type = SpaceType::BackwardMirror;
        } else if c == '|' {
            space_type = SpaceType::VerticalSplitter;
        } else if c == '-' {
            space_type = SpaceType::HorizontalSplitter;
        } else {
            space_type = SpaceType::False;
        }
        
        GridSpace{
            space_type,
            is_energized: false,
            direction: vec![], 
        }
    }

}

fn parse_file(file_name: &str) -> Vec<Vec<GridSpace>> {
    String::from_utf8(fs::read(file_name).unwrap())
    .unwrap()
    .split("\n")
    .map(|line| {

        line
        .chars()
        .map(|c| {
            GridSpace::new(c)
        })
        .collect::<Vec<GridSpace>>()
    })
    .collect()

}

fn print(grid_spaces: &Vec<Vec<GridSpace>>) {
    for row in grid_spaces.iter() {
        // println!("{:?}");

        for grid_space in row {
            // print!("{:?} ", grid_space.space_type);
            let space;
            if grid_space.is_energized {
                space = "#";
            } else {
                space = ".";
            }

            print!("{}", space);
        }
        println!("");
    }
}


fn beam(mut grid_spaces: &mut Vec<Vec<GridSpace>>, mut row: usize, mut col: usize, row_vel: i8, col_vel: i8, first: bool) {
    let mut direction: Direction = Direction::Uninitialized;

    if !first {

        if row_vel < 0 {
            row = match row.checked_sub(1) {
                None => { return },
                Some(s) => s,
            };
            // UP
            direction =Direction::North;
        } else if row_vel > 0 {
            row+=1;
            direction =Direction::South;
        }
        
        if col_vel < 0 {
            col = match col.checked_sub(1) {
                None => { return },
                Some(s) => s,
            };
            direction =Direction::West;
        } else if col_vel > 0 {
            col +=1;
            direction =Direction::East;
        }
        
        match grid_spaces.get_mut(row) {
            None => {},
            Some(row_gs) => {
                match row_gs.get_mut(col) {
                    None => {},
                    Some(gs) => {
                        if gs.direction.contains(&direction) {
                            return 
                        } else {
                            gs.direction.push(direction);
                        }
                    }
                }
            }
        }
    } 
        
    // I want to differentiate >>>

match grid_spaces.get_mut(row) {
    None => { return },
    Some(row_gs) => {
        match row_gs.get_mut(col) {
            None => { return },
            Some(ref mut grid_space) => {
                
                grid_space.is_energized = true;
                
                    match grid_space.space_type {
                        SpaceType::BackwardMirror => {
                            // moving right, \ will go DOWN. so row 1
                        if col_vel > 0 {
                            beam(&mut grid_spaces, row, col,  1, 0, false);
                        } else if col_vel < 0 {  // moving left \ will go UP
                            // moving left and hits / means it will go down
                            beam(&mut grid_spaces, row, col,  -1, 0, false);
                        } else if row_vel > 0 {  // moving down and hits \ means it will go right                        
                            beam(&mut grid_spaces, row, col,  0, 1, false);
                        } else if row_vel < 0 {  // moving up and hits \ means it will go left                    
                            beam(&mut grid_spaces, row, col,  0, -1, false);
                        }
                        },
                        SpaceType::EmptySpace => {
                            beam(&mut grid_spaces, row, col, row_vel, col_vel, false); 
                        },
                        SpaceType::False => {
                            println!("DONT BE HERE");
                        }, SpaceType::ForwardMirror => {
                            // moving right => go up /
                            if col_vel > 0 {  // right => /. go up
                                beam(&mut grid_spaces, row, col,  -1, 0, false);
                            } else if col_vel < 0 {  // left => / DOWN
                                // moving left and hits / means it will go down
                                beam(&mut grid_spaces, row, col,  1, 0, false);
                            } else if row_vel > 0 {  // down hits / means it will go left                        
                                beam(&mut grid_spaces, row, col,  0, -1, false);
                            } else if row_vel < 0 {  // moving up and hits / means it will go right                    
                                beam(&mut grid_spaces, row, col,  0, 1, false);
                            }
                            
                        }, SpaceType::HorizontalSplitter => {
                            if row_vel != 0 {  // were going up or down and hit -, making us go right & left
                                beam(&mut grid_spaces, row, col, 0, -1, false); 
                                beam(&mut grid_spaces, row, col, 0, 1, false); 
                            } else {
                                // beam()
                                beam(&mut grid_spaces, row, col, row_vel, col_vel, false); 
                            }
                        }, SpaceType::VerticalSplitter => {

                            // going left or right
                            if col_vel != 0 {  // were going left or right and hit |, making us go up & down
                                beam(&mut grid_spaces, row, col, -1, 0, false); 
                                beam(&mut grid_spaces, row, col, 1, 0, false); 
                            } else {  // then ur going up or dowm. so just pass thru
                                beam(&mut grid_spaces, row, col, row_vel, col_vel, false); 

                            }
                        }
                    }
                    return;
                }
            }
        }
    }
}

fn calc_sum(mut grid_spaces: &mut Vec<Vec<GridSpace>>) -> u64 {
    // grid_spaces.get_mut(0).unwrap().get_mut(0).unwrap().is_energized = true;

     grid_spaces.iter()
    .map(|row_gs| row_gs.iter().map(|gs| {
        if gs.is_energized {
            return 1
        } else {
            return 0
        };
    }).sum::<u64>()).sum()
    
}

fn part_1() {
    let args: Vec<String> = env::args().collect();
    
    let mut grid_spaces: Vec<Vec<GridSpace>> = parse_file(args.get(1).unwrap());
    
    if args.get(1).unwrap() == "adventday17.txt" {
        beam(&mut grid_spaces,  0, 0, 1, 0, true);
    } else {
        beam(&mut grid_spaces,  0, 0, 0, 1, true);
    }
    
    let sum = calc_sum(&mut grid_spaces);
    print(&grid_spaces);
    println!("Sum: {}", sum);
    
}

// fn create_boundaries(mut grid_spaces: &mut Vec<Vec<GridSpace>>) {
//     grid_spaces.iter_mut().for_each(|row| {
//         row.push(GridSpace::new('.'));
//         row.insert(0, GridSpace::new('.'));
//     });

//     grid_spaces.insert(0, Vec::from(grid_spaces.))
// }
fn reset(mut grid_spaces: &mut Vec<Vec<GridSpace>>) {
    grid_spaces.iter_mut().for_each(|row| {
        row.iter_mut().for_each(|gs| {
            gs.is_energized = false;
        })
    })
}

fn main() {
    // part_1();
    let args: Vec<String> = env::args().collect();
    
    let mut grid_spaces: Vec<Vec<GridSpace>> = parse_file(args.get(1).unwrap());
    let size_col = grid_spaces.len();    
    let size_row = grid_spaces[0].len();    

    let mut max: u64 = 0;

    for num in 0..grid_spaces.len() {
        beam(&mut grid_spaces,  num, 0,0, 1, true);
        let sum = calc_sum(&mut grid_spaces);         
        if sum > max {
            max = sum;
        } 
        reset(&mut grid_spaces);
        beam(&mut grid_spaces,  num, size_row-num-1,0, -1, true);
        let sum = calc_sum(&mut grid_spaces);         
        if sum > max {
            max = sum;
        } 
        reset(&mut grid_spaces);
    }
    
    for num in 0..grid_spaces[0].len() {
        beam(&mut grid_spaces,  0, num, 1, 0, true);
        let sum = calc_sum(&mut grid_spaces);         
        if sum > max {
            max = sum;
        } 
        reset(&mut grid_spaces);
        beam(&mut grid_spaces,  size_col-1-num, num,-1, 0, true);
        let sum = calc_sum(&mut grid_spaces);         
        if sum > max {
            max = sum;
        } 
        reset(&mut grid_spaces);
    }
        // print(&grid_spaces);
    println!("Sum: {}", max);
 

}


#[cfg(test)]
mod test{
    use crate::{GridSpace, SpaceType, parse_file, beam, calc_sum, print};


    #[test]
   fn test_grid_space() {
    let grid_space = GridSpace::new('\\');
    
        assert_eq!(SpaceType::BackwardMirror, grid_space.space_type);
   } 

   fn test(name: &str, row: usize, col: usize, row_vel: i8, col_vel: i8) -> u64 {

    let mut grid_spaces = parse_file(name);
    
    beam(&mut grid_spaces, row, col, row_vel, col_vel, true);
    let sum = calc_sum(&mut grid_spaces);
    println!("Grid:");
    print(&grid_spaces);

    sum
   }

    // #[test]
   fn test_beam() {
    let sum = test("test.txt", 0, 0, 0 ,1);
    assert_eq!(46, sum); 
   } 

   #[test]
   fn test_right_down() {
        let sum = test("test_down.txt", 0, 0, 1, 0);

        println!("{}", sum);
   }
}