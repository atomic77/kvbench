DROP TABLE IF EXISTS logs;

CREATE TABLE logs (Timestamp int, Label text, TotThd int, Run int, DB text, Test text, ThreadId int, Operations int, Seconds float, Rate float);

. separator ","
. import all_split.csv logs

-- The final 
-- SELECT Label, TotThd, Run, DB, Test, sum(Operations) as TotOps, avg(Seconds) as Seconds, sum(Rate) as SumRate
SELECT Label, TotThd, Run, DB, Test, sum(Rate) as SumRate FROM logs GROUP BY 1,2,3,4,5 LIMIT 10;




