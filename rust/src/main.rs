// Name of the project: rust
// use rust::algos::bits::count;
// use rust::common;
// use rust::math;

#[derive(Debug)]
struct Item {
  id: u8,
  title: String,
  parent_id: Option<u8>,
}

#[derive(Debug, Clone)]
struct TreeNode {
  id: u8,
  title: String,
  parent_id: Option<u8>,
  children: Vec<TreeNode>,
}

fn main() {
  let xs: Vec<Item> = vec![
    Item { id: 1, title: "Europe".to_string(), parent_id: None },
    Item { id: 2, title: "Asia".to_string(), parent_id: None },
    Item { id: 3, title: "France".to_string(), parent_id: Some(1) },
    Item { id: 4, title: "Thailand".to_string(), parent_id: Some(2) },
    Item { id: 5, title: "Paris".to_string(), parent_id: Some(3) }
  ];

  let r = build_tree(&xs, None);
  println!("{:#?}", r);
}

fn build_tree(xs: &Vec<Item>, parent_id: Option<u8>) -> Vec<TreeNode> {
  xs.iter()
    .filter(|x| x.parent_id == parent_id)
    .map(|x| TreeNode {
      title: x.title.clone(),
      id: x.id,
      parent_id: x.parent_id,
      children: build_tree(xs, Some(x.id)),
    })
    .collect()
}