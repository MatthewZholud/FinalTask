CREATE TABLE groups(
    group_id Serial primary key,
    title varchar(255) not null
);
CREATE TABLE tasks(
    task_id serial primary key,
    title varchar(255) not null,
    group_id int REFERENCES groups(group_id) ON DELETE CASCADE
);
CREATE TABLE timeframes(
    task_id int REFERENCES tasks(task_id) ON DELETE CASCADE,
    time_from timestamp without time zone,
    time_to timestamp without time zone
);