-- GetPartInfo
SELECT kbuban, revision, krevision 
FROM original_schema.ok_100_part 
WHERE seppenbuban = :1;
