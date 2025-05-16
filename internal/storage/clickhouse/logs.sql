CREATE TABLE logs (
    ip_address String,
    method String,
    path String,
    query String,
    headers String,
    body String,
    status_code Int32,
    response_time Float64,
    created_at DateTime
) ENGINE = MergeTree()
ORDER BY created_at;
