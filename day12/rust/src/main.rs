use std::fs;

#[derive(Debug)]
struct HotSpringRow {
    conditions: String,
    cont_damaged: Vec<u32>
}

#[derive(Debug)]
enum RowType {
    Found,
    Unknown,
}

#[derive(Debug)]
struct Row {
    row: String,
    row_type: RowType,
}

impl Row {
    fn new(row: &str, is_unknown: bool) -> Row {
        let mut row_type = RowType::Found;

        if is_unknown { 
            row_type = RowType::Unknown;
        }

        Row { row: row.to_string(), row_type }
    }
}

impl HotSpringRow {
    fn new(conditions: &str, cont_damaged: Vec<u32>) -> HotSpringRow {
        HotSpringRow{
            conditions: conditions.into(),
            cont_damaged
        }
    }

    fn get_rows(&self) -> Vec<Row> {
        // double dimension array
        // number of markers
        // maximum size of marker
        const num: usize = 999999;
        // get combinations at each point p
        let mut all_rows: Vec<Row> = vec![]; 
        let mut cut: usize = num;
        let mut on_hot_springs = false;
        let mut is_unknown: bool = false;
        self.conditions.chars().enumerate()
        .for_each(|(pos, c)| { 
           if !on_hot_springs && (c == '#' || c == '?') {
            cut = pos;
            on_hot_springs = true;
        }
        
        if c == '?' {
            is_unknown = true;
        }
        
        if c == '.' && cut != num {
            let row = Row::new(&self.conditions[cut..pos], is_unknown);
            all_rows.push(row);
            on_hot_springs = false;
           cut = num; 
           is_unknown = false;
           }
        });

        if on_hot_springs {
            let row = Row::new(&self.conditions[cut..self.conditions.len()], is_unknown);
            all_rows.push(row);
        }

        all_rows

    }

    fn get_combinations(&mut self) -> u32 {
        let all_rows: Vec<Row> = self.get_rows();
        let cont_left: Vec<u32> = vec![];

        self.cont_damaged.sort();

        all_rows.iter().filter(|predicate| {

        match predicate.row_type {
            RowType::Found => {
                return true;
            },
            RowType::Unknown => {
                return false;
            }
        }    
        })
        .for_each(|found_row| {
            let row: u32= found_row.row.len() as u32;

            let mut num: usize= 0;
            loop {
                if num == self.cont_damaged.len() {
                    break;
                }
                if row == self.cont_damaged[num] {
                    self.cont_damaged.remove(num);
                    num = 0;
                }

                num+=1;
            }    
        });

        let mut combinations = all_rows.iter().filter(|p| {
            match p.row_type {
                RowType::Found => {
                    return false;
                },
                RowType::Unknown => {
                    return true;
                }
            }
        }).for_each(|unknown_row| {
            // greedy alg
            let spot_combo: Vec<u32> = vec![];

            
            for num in unknown_row.row.chars() {
                
            }
        
        });
    0
    
    }
}


fn parse_file(file_name: String) -> Vec<HotSpringRow> {

    String::from_utf8(fs::read(file_name).unwrap()).unwrap()
    .split("\n")
    .map(|line| {
        let mut contents = line.split_whitespace();
       let left = contents.next().unwrap(); 
       let right = contents.next().unwrap(); 


        HotSpringRow{
            conditions: left.to_owned(),
            cont_damaged: right.split(",").map(|num|  num.parse::<u32>().unwrap()).collect(),
        }        
    })
    .collect()


}



fn main() {
    let content = parse_file("test.txt".into());


    println!("{:?}", content);
    println!("Hello, world!");
}


#[cfg(test)]
mod test {
    use crate::{HotSpringRow, Row, RowType};

    #[test]
    fn test_get_row() {
        let rows = test_rows("???.###", vec![]);
        assert_eq!("???", rows[0].row);
        assert_eq!("###", rows[1].row);
    
        let row_string = "?#?#?#?#?#?#?#?#?"; 
         let rows = test_rows(row_string ,vec![]);
        assert_eq!(row_string, rows[0].row);
   
        let row_string = "????.#...#..."; 
         let rows = test_rows(row_string ,vec![]);
        assert_eq!("????", rows[0].row);
        assert_eq!("#", rows[1].row);
        assert_eq!("#", rows[2].row);
    }

     
    fn test_rows(conditions: &str, cont_damaged: Vec<u32>) -> Vec<Row> {
        let hsr = HotSpringRow::new(conditions, cont_damaged); 
        let combos = hsr.get_rows();
        
        
        combos
    }


    fn test(name: &str, conditions: &str, cont_damaged: Vec<u32>) -> u32 {
        let mut hsr = HotSpringRow::new(conditions, cont_damaged); 
        let combos = hsr.get_combinations();
        println!("{}: {}",name, combos);
        combos
    }

    #[test]
    fn test_combinations() {
        let combos = test("First test", "???.###", vec![1,1,3]);
        assert_eq!(0, combos);
    }
}