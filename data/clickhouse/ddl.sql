create table if not exists default.clicks
(
    user_id UUID,
    username String,
    message String,
    created_at Timestamp
    )
    engine=Kafka('kafka:29092','clicks','rushawx','JSONEachRow')
;

create table if not exists daily (
    day Date,
    total UInt64
)
engine = SummingMergeTree()
primary key day
order by day
;

create materialized view clicks_daily to daily
as select toDate(toDateTime(created_at)) AS day, count() as total
from clicks group by day
;
