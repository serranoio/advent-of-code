use std::{fs, collections::{HashSet, HashMap}};



fn parse_file(file_name: &str) -> Vec<Vec<u32>> {
   String::from_utf8(fs::read(file_name).unwrap()).unwrap()
   .split("\n")
   .map(|s| {
        s.chars()
        .map(|s| {
            s.to_digit(10).unwrap()
        })
        .collect()
   }).collect()
}

// create empty priority queue Q
// create empty set visited
// create array distances with size equal to the number of vertices in the graph
// set initial distance values to infinity for all vertices
// set distance[start] to 0
// enqueue (start, 0) into Q

// while Q is not empty:
//     current, current_distance = dequeue from Q

//     if current is in visited:
//         continue

//     mark current as visited

//     for neighbor in neighbors of current:
//         if neighbor is in visited:
//             continue

//         new_distance = current_distance + distance between current and neighbor

//         if new_distance < distances[neighbor]:
//             distances[neighbor] = new_distance
//             enqueue (neighbor, new_distance) into Q

#[derive(Eq, Hash, PartialEq)]
struct Node {
    coordinates: (usize, usize),
    weight: u32,
}

fn create_distances(nodes: &Vec<Node>) ->  HashMap<(usize, usize), u32> { 
    let mut distances: HashMap<(usize, usize), u32> = HashMap::new(); 
    
    for node in nodes {
            distances.insert((node.coordinates.0, node.coordinates.1), std::u32::MAX);
    }

    distances
}

fn create_nodes(graph: &Vec<Vec<u32>>) -> Vec<Node> {
    let mut nodes: Vec<Node> = vec![];

    for (y, row) in graph.iter().enumerate() {
        for (x, node) in row.iter().enumerate() {
            nodes.push(Node{
                coordinates: (x, y),
                weight: *node,
            })
        }
    }

    nodes
}

// return distances
fn asdfghjkl(graph: &Vec<Vec<u32>>) {
    
    let mut nodes: Vec<Node> = create_nodes(&graph);
    
    let mut p_q: Vec<Node> = vec![];
    
    let mut visited: HashSet<(usize, usize)> =  HashSet::new();    
    
    let mut distances: HashMap<(usize, usize), u32> = create_distances(&nodes);  
    *distances.get_mut(&(0, 0)).unwrap() = 0;
    visited.insert((0,0));
    
    p_q.push(nodes.remove(0));
    
    while !p_q.is_empty() {
        
        let vertice = p_q.pop().unwrap();
        
        // if currnet is in visited, continue
        if visited.contains(&(vertice.coordinates.0, vertice.coordinates.1)) {
            continue;
        }
        
        visited.insert((vertice.coordinates.0, vertice.coordinates.1));
        
        // for neighbor of vertice's neighbors
        
        
        
}

}
// everything is one space apart
// the weights of the nodes depends on the node itself

fn main() {
    let test = parse_file("test.txt");

    asdfghjkl(&test);

    println!("{:?}", test);
    println!("Hello, world!");
}
