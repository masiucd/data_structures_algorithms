#[derive(Debug)]
struct User {
  id: i32,
  name: String,
  born: i32,
}

#[derive(Debug)]
struct Comment {
  id: i32,
  content: String,
  author_id: i32,
}

#[allow(dead_code)]
pub fn list_cardio() {
  let comments = vec![
    Comment {
      id: 1,
      content: "hello".to_string(),
      author_id: 1,
    },
    Comment {
      id: 2,
      content: "world".to_string(),
      author_id: 2,
    },
    Comment {
      id: 3,
      content: "hello".to_string(),
      author_id: 3,
    },
    Comment {
      id: 4,
      content: "world".to_string(),
      author_id: 4,
    },
    Comment {
      id: 5,
      content: "hello".to_string(),
      author_id: 5,
    }
  ];

  let users = vec![
    User {
      id: 1,
      name: "John".to_string(),
      born: 1990,
    },
    User {
      id: 2,
      name: "Jane".to_string(),
      born: 1991,
    },
    User {
      id: 3,
      name: "Jim".to_string(),
      born: 2018,
    },
    User {
      id: 4,
      name: "Jill".to_string(),
      born: 1993,
    },
    User {
      id: 5,
      name: "Jack".to_string(),
      born: 1994,
    }
  ];

  let todays_yaer = 2020;
  // is there users over 19 years old?
  let has_users_over_19 = users.iter().any(|x| todays_yaer - x.born < 19);
  println!("has_users_over_19: {}", has_users_over_19);
  let is_every_user_over_19 = users.iter().all(|x| todays_yaer - x.born >= 19);
  println!("is_every_user_over_19: {}", is_every_user_over_19);

  let comments_with_id_five = comments.iter().find(|x| x.id == 5); // Some(Comment { id: 5, content: "hello", author_id: 5 })
  println!("comments_with_id_five: {:?}", comments_with_id_five);

  let comments_without_id_3: Vec<Comment> = comments
    .into_iter()
    .filter(|x| x.id != 3)
    .collect();
  println!(" comments_without_id_3 {:?}", comments_without_id_3)
}