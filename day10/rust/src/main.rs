use std::{collections::HashMap, path::PathBuf, fs, time::SystemTime};

#[derive(Debug, Hash, PartialEq, Eq, Clone, Copy)]
enum Cardinal {
    North,
    South,
    East,
    West,
}

fn parse_file(file_name: &str) -> Vec<String> {
   let path = PathBuf::from(file_name);
    let contents = String::from_utf8(fs::read(path).unwrap()).unwrap();
    let lines: Vec<String> = contents.split('\n').map(|s| s.to_owned()).collect();
    
    lines 
}

fn move_to_pipe(place: (char, Cardinal, (usize, usize)), lines: &Vec<String>,
connections: &HashMap<Cardinal, Cardinal>, pipes: &HashMap<char, Vec<Cardinal>>) -> (char, Cardinal, (usize, usize)) {
        // Cardinal::East -> Cardinal::West
        let start = *connections.get(&place.1).unwrap();
        let openings = pipes.get(&place.0).unwrap();
    
        let mut next: (char, Cardinal, (usize, usize)) = ('0', Cardinal::East, (0,0));
        if start == openings[0] {
            next.1 = openings[1];
        } else {
            next.1 = openings[0];
        }

    // if we have cardinal east, it means travel east
    match next.1 {
        Cardinal::East => {
            next.0 = lines[place.2.0].chars().nth(place.2.1+1).unwrap();
            next.2 = (place.2.0, place.2.1+1);
        },
        Cardinal::West => {
            next.0 = lines[place.2.0].chars().nth(place.2.1-1).unwrap();
            next.2 = (place.2.0, place.2.1-1);
        },
        Cardinal::North => {
            next.0 = lines[place.2.0-1].chars().nth(place.2.1).unwrap();
            next.2 = (place.2.0-1, place.2.1);
        },
        Cardinal::South  => {
            next.0 = lines[place.2.0+1].chars().nth(place.2.1).unwrap();
            next.2 = (place.2.0+1, place.2.1);
        },
    };
    // direction you come from
    
    
    next
}

fn find_matching(to_direction: Cardinal, row: usize, col: usize, lines: &Vec<String>, pipes: &HashMap<char, Vec<Cardinal>>) -> bool {
   let mut available: Vec<char> = vec![];
    match to_direction {
        Cardinal::West => {
        available.push('F');
        available.push('-');
        available.push('L');
        },
        Cardinal::East => {
            available.push('7');
            available.push('J');
            available.push('-');
        },
        Cardinal::North => {         
            available.push('F');
            available.push('7');
            available.push('|');
        },
        Cardinal::South  => { 
            available.push('|');
            available.push('L');
            available.push('J');
        },
    }

    if (row) as usize >= 0 && row < lines.len() && (col) as usize >= 0 && col < lines[0].len() {

        
        let c = &lines[row].chars().nth(col).unwrap();
        
        if  pipes.contains_key(c) && available.contains(c) {
            return true;
        }
    }

    return false;
}

fn find_starting_pipe(starting_position: (usize, usize), lines: Vec<String>,
connections: HashMap<Cardinal, Cardinal>, pipes: HashMap<char, Vec<Cardinal>>) -> i32 {
    let row: usize = starting_position.0 as usize;
    let col: usize = starting_position.1 as usize;
    let mut looper: Vec<(char, Cardinal, (usize, usize))> = vec![];

    
        if find_matching(Cardinal::South, row+1, col, &lines, &pipes) {
                looper.push((lines[row+1].chars().nth(col).unwrap(), Cardinal::South, (row+1, col)));
        };
        
        if find_matching(Cardinal::North, row-1, col, &lines, &pipes) {
            looper.push((lines[row-1].chars().nth(col).unwrap(), Cardinal::North, (row-1, col)));
        };
        
        if find_matching(Cardinal::East, row, col+1, &lines, &pipes) {                
            looper.push((lines[row].chars().nth(col+1).unwrap(), Cardinal::East, (row, col+1)));
        };
        
        if find_matching(Cardinal::West, row, col-1, &lines, &pipes) {               
            looper.push((lines[row].chars().nth(col-1).unwrap(), Cardinal::West, (row, col-1)));

        };


        // if we go north, then we can only match with...
        let mut counter = 1;
        loop { 
            
            let next = move_to_pipe(looper.pop().unwrap(), &lines, &connections, &pipes);
            let next1 = move_to_pipe(looper.pop().unwrap(), &lines, &connections, &pipes);
            
            counter+=1;
            if next.2.0 == next1.2.0 &&
                        next.2.1 == next1.2.1{
                            break;
                    }             

                    looper.push(next);  
                    looper.push(next1);  
        }   


        counter

    }
    
    
    fn main() {
    let beginning = SystemTime::now();
    let pipes = HashMap::from([
        ('7', vec![Cardinal::South, Cardinal::West]),
        ('|', vec![Cardinal::North, Cardinal::South]),
        ('L', vec![Cardinal::North, Cardinal::East]),
        ('F', vec![Cardinal::South, Cardinal::East]),
        ('J', vec![Cardinal::North, Cardinal::West]),
        ('-', vec![Cardinal::East, Cardinal::West]),
    ]);

    let connections = HashMap::from([
        (Cardinal::North, Cardinal::South),
        (Cardinal::South, Cardinal::North),
        (Cardinal::East, Cardinal::West),
        (Cardinal::West, Cardinal::East),
    ]);


    let lines = parse_file("adventday10.txt");

    let mut starting_position = (0, 0);

    lines.iter().enumerate().for_each(|(row, line)| {
        line.char_indices().for_each(|(column, c)| { 
            if c == 'S' {
                starting_position.0 = row;
                starting_position.1 = column;
                return;
            }
        })
    });

    let steps = find_starting_pipe(starting_position, lines, connections, pipes);
    
    println!("Steps: {}", steps);
   println!("{:?}", beginning.elapsed().unwrap()); 
}

#[cfg(test)]
mod test {
    use std::time::SystemTime;

    use crate::main;


    #[test]
    fn test_main() {
    
        main();
    }
}
