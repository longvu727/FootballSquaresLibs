CREATE TABLE squares (
    square_id       int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    square_guid     varchar(36),
    square_size     int default 10,
    row_points      varchar(256),
    column_points   varchar(256),
    created         timestamp default now(), 
    updated         timestamp default now() on update now() 
);