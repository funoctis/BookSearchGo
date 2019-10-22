CREATE TABLE `sessions` (
    `session_id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `session_token` TEXT,
    `username` TEXT,
    foreign key(username) references users(username)
)

CREATE TABLE `users` (
    `user_id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` TEXT, `password` TEXT
)