-- GetPartInfo
SELECT kbuban, revision, krevision 
FROM original_schema.OK_PART 
WHERE seppenbuban = ?;
