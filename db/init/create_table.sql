drop table if exists TEST_USER;

create table TEST_USER (
    user_id BIGINT PRIMARY KEY,
    user_password VARCHAR(20) NOT NULL
);