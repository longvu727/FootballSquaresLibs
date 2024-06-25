CREATE TABLE users (
    user_id     int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_guid   varchar(36) NOT NULL,
    ip          varchar(16),
    device_name varchar(256),
    user_name   varchar(64),
    alias       varchar(8),
    created     timestamp default now(), 
    updated     timestamp default now() on update now() 
);