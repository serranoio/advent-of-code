use std::{fs, collections::HashMap};


#[derive(Debug)]
enum RockType {
    Rounded,
}

#[derive(Debug)]
struct Coordinates {
    x: usize,
    y: usize,
}

#[derive(Debug)]
struct Rock {
    position: Coordinates
}

const ROUNDED: char = 'O';

impl Rock {
    fn new(position: (usize, usize)) -> Rock {
        Rock{
            position: Coordinates { x: position.0, y: position.1 }
        }
    }
}

fn parse_file(file_name: &str) -> (Vec<Vec<char>>, Vec<Rock>){
    let mut rounded_rocks: Vec<Rock> = vec![];
    let rocks: Vec<Vec<char>> = String::from_utf8(fs::read(file_name).unwrap()).unwrap()
    .split("\n")
    .enumerate()
    .map(|(row, line)| {
        line
        .chars()
        .enumerate()
        .for_each(|(col, c)| {
            if c == ROUNDED {
                rounded_rocks.push(Rock::new((row, col)));
            }
        });
        line
        .chars()
        .collect::<Vec<char>>()
    })
    .collect();

    (rocks, rounded_rocks)
}
#[derive(PartialEq)]
enum Direction {
    North,
    South,
    East,
    West
}

fn get_border(direction: &Direction, position: &Coordinates, length: usize, width: usize) -> bool {
    match direction {
       Direction::North => {
        if position.x > 0 {
            return true;
        } else {
            return false;
        }
       }, Direction::South => {
        if position.x < length - 1 {
            return true;
        } else {
            return false;
        } 
       }, Direction::West => {
        if position.y > 0 {
            return true;
        } else {
            return false;
        }
       }, Direction::East => {
        if position.y < width - 1 {
            return true;
        } else {
            return false;
        }
       }
    }
}

fn get_base_case(direction: &Direction, position: &Coordinates, length: usize, width: usize) -> bool {
    match direction {
       Direction::North => {
        if position.x == 0 {
            return true;
        } else {
            return false;
        }
       }, Direction::South => {
        if position.x >= length -1 {
            return true;
        } else {
            return false;
        } 
       }, Direction::West => {
        if position.y == 0 {
            return true;
        } else {
            return false;
        }
       }, Direction::East => {
        if position.y >= width -1 {
            return true;
        } else {
            return false;
        }
       }
    }
}

fn check_and_move(contents: &mut Vec<Vec<char>>, rock: &mut Rock, direction: &Direction) {


    
    let mut position: Coordinates = Coordinates{x: rock.position.x, y: rock.position.y};
    let mut next_position = '.';
    
    if get_base_case(direction, &position, contents.len(), contents[0].len()) {
        return
    }

while next_position != ROUNDED &&
    next_position != '#' &&
    get_border(&direction, &position, contents.len(), contents[0].len()) { 
        
    match direction {
        Direction::North => {
            position.x = position.x-1; 
        }, Direction::South => {
            position.x = position.x+1; 
        }, Direction::East => {
            position.y = position.y+1;
        }, Direction::West => {
            position.y = position.y-1;
        }
    }

    next_position = contents[position.x][position.y];
}

if next_position == ROUNDED ||
    next_position == '#' {
        match direction {
            Direction::North => {
            position.x+=1;
            }, Direction::South => { 
            position.x-=1;
            }, Direction::West => {  
            position.y+=1;
        }, Direction::East => { 
            position.y-=1;
            }
        }
}

let temp = contents[position.x][position.y];
contents[position.x][position.y] = contents[rock.position.x][rock.position.y];
contents[rock.position.x][rock.position.y] = temp;

rock.position.x = position.x;
rock.position.y = position.y;
}

fn move_up(mut contents: &mut Vec<Vec<char>>, rocks: &mut Vec<Rock>, direction: &Direction) {

    rocks.iter_mut()
    .for_each(|rock| { 
        check_and_move(&mut contents, rock, &direction);
    });
}

fn count(rocks: Vec<Rock>, south_edge: usize) -> usize {

    rocks
    .iter()
    .map(|rock| {
        south_edge - rock.position.x
    })
    .sum()
}

const BILLION: u64 = 1000000000;
fn main() {
    let (mut contents, mut rocks)= parse_file("adventday14.txt");

    let mut cycle: HashMap<Vec<Vec<char>>, u32>= HashMap::new();

    for _ in 0..BILLION {
        move_up(&mut contents, &mut rocks, &Direction::North);
        rocks.sort_by(|a, b| a.position.y.cmp(&b.position.y)); 
        move_up(&mut contents, &mut rocks, &Direction::West);
        rocks.sort_by(|a, b| b.position.x.cmp(&a.position.x)); 
        move_up(&mut contents, &mut rocks, &Direction::South);
        rocks.sort_by(|a, b| b.position.y.cmp(&a.position.y)); 
        move_up(&mut contents, &mut rocks, &Direction::East);
        rocks.sort_by(|a, b| a.position.x.cmp(&b.position.x)); 

        let cycle_clone = cycle.clone();
        *cycle.entry(contents.clone()).or_insert(0) += 1;
    
        let found: HashMap<Vec<Vec<char>>, u32>  =cycle_clone 
        .into_iter()
        .filter(|p| p.1 > 1)
        .collect();
        if found.len() > 0 {
            break;
        }
    }

    let all = count(rocks, contents.len());
    println!("{:?}", all);
}

#[cfg(test)]
mod test {
    use crate::{parse_file, check_and_move, RockType, ROUNDED, move_up, count, Direction};

    #[test]
    fn test_main() {
        parse_file("test.txt");
        println!("TEST")
    }

    #[test]
    fn test_check_and_move() {
        let (mut contents, mut rocks) = parse_file("test.txt");
        
        check_and_move(&mut contents,  &mut rocks[9], &Direction::North);
        assert_eq!(contents[0][7], ROUNDED); // slides to edge
     
        check_and_move(&mut contents,  &mut rocks[0], &Direction::North); 
        assert_eq!(contents[0][0], ROUNDED); // on edge 
   
         check_and_move(&mut contents,  &mut rocks[7], &Direction::North); 
        assert_eq!(contents[2][9], ROUNDED); // stops on # 
     
          check_and_move(&mut contents,  &mut rocks[4], &Direction::North); 
        assert_eq!(contents[2][0], ROUNDED); // slides to edge    
       
        
        
        // check_and_move(&mut contents, &mut rocks[17]);


    }

    #[test]
    fn test_move_up() {
 
        let (mut contents, mut rocks)= parse_file("test.txt");

        move_up(&mut contents, &mut rocks);

        assert_eq!(vec![ROUNDED, ROUNDED, ROUNDED, ROUNDED, '.', '#', '.', ROUNDED, '.', '.'], contents[0])
    }

    #[test]
    fn test_all() {
 
        let (mut contents, mut rocks)= parse_file("test.txt");

        move_up(&mut contents, &mut rocks);

        assert_eq!(136, count(rocks, contents.len()));

    }

}