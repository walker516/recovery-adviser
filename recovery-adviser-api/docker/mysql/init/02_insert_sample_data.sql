INSERT INTO original_schema.ok_100_part (seppenbuban, kbuban, revision, krevision) 
VALUES ('SEPPEN001', 'KBUBAN001', 'REV001', 'KREV001');

INSERT INTO original_schema.ok_101_job_queue (process_order, status, host, register_timestamp, parameter) 
VALUES ('PO123', 2, 'host1', NOW(), 'KBUBAN001');

INSERT INTO original_schema.ok_102_job_lock (process_order, lock_timestamp) 
VALUES ('PO123', NOW());
