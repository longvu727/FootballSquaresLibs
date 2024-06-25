CREATE TABLE games (
    game_id     int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    game_guid   varchar(36) NOT NULL,
    sport       varchar(64),
    team_a      varchar(64),
    team_b      varchar(64),
    created     timestamp default now(), 
    updated     timestamp default now() on update now() 
);