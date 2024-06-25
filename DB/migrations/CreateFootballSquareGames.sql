CREATE TABLE football_square_games (
    football_square_game_id     int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    game_id                     int,
    square_id                   int,
    user_id                     int,
    winner                      boolean default false,
    winner_quarter_number       tinyint default 0,
    row_index                   int,
    column_index                int,
    created                     timestamp default now(), 
    updated                     timestamp default now() on update now()
);