CREATE SCHEMA IF NOT EXISTS original_schema;

CREATE TABLE IF NOT EXISTS original_schema.ok_100_part (
    seppenbuban VARCHAR(255) PRIMARY KEY,
    kbuban VARCHAR(255),
    revision VARCHAR(255),
    krevision VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS original_schema.ok_101_job_queue (
    process_order VARCHAR(255) PRIMARY KEY,
    status INT,
    host VARCHAR(255),
    register_timestamp TIMESTAMP,
    parameter VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS original_schema.ok_102_job_lock (
    process_order VARCHAR(255) PRIMARY KEY,
    lock_timestamp TIMESTAMP
);
