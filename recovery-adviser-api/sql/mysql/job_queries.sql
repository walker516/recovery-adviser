-- GetRecoveryJobStatus
WITH OrderedJobs AS (
    SELECT 
        process_order, 
        status, 
        host, 
        register_timestamp,
        parameter,
        ROW_NUMBER() OVER (ORDER BY register_timestamp DESC) AS row_num
    FROM 
        original_schema.ok_101_job_queue
    WHERE 
        parameter LIKE CONCAT('%', (SELECT kbuban FROM original_schema.ok_100_part WHERE seppenbuban = ?), '%')
)
SELECT
    MAX(CASE WHEN row_num = 1 THEN process_order END) AS latest_process_order,
    MAX(CASE WHEN row_num = 1 THEN register_timestamp END) AS latest_register_timestamp,
    MAX(CASE WHEN row_num = 1 THEN host END) AS latest_host,
    MAX(CASE WHEN status IN (1, 2) THEN 1 ELSE 0 END) AS needs_investigation,
    MAX(CASE WHEN row_num = 1 AND status = 3 AND parameter LIKE '%revise%' THEN 1 ELSE 0 END) AS needs_detailed_review,
    MAX(CASE WHEN row_num = 1 AND status = 3 AND parameter NOT LIKE '%revise%' THEN 1 ELSE 0 END) AS job_not_completed_correctly,
    MAX(CASE WHEN row_num = 1 AND status = 4 THEN 1 ELSE 0 END) AS error_occurred_during_job
FROM 
    OrderedJobs;

-- GetJobQueueByProcessOrder
SELECT process_order, status, host, register_timestamp, parameter 
FROM original_schema.ok_101_job_queue 
WHERE process_order = ?;

-- GetJobQueueBySeppenbuban
SELECT process_order, status, host, register_timestamp, parameter 
FROM original_schema.ok_101_job_queue 
WHERE parameter LIKE CONCAT('%', (SELECT kbuban FROM original_schema.ok_100_part WHERE seppenbuban = ?), '%') 
ORDER BY register_timestamp DESC 
LIMIT 1;

-- UpdateJobQueue
UPDATE original_schema.ok_101_job_queue 
SET status = ?, host = ? 
WHERE process_order = ?;

-- GetJobLock
SELECT process_order, lock_timestamp 
FROM original_schema.ok_102_job_lock 
WHERE process_order = ?;

-- DeleteJobLock
DELETE FROM original_schema.ok_102_job_lock 
WHERE process_order = ?;
