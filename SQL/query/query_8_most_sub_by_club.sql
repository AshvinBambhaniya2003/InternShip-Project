-- Query-8 : Which club has done most substitutions? also show number of substitutions.

SELECT
    clubs.name,
    COUNT(*) AS total_substitutions
FROM
    game_events 
JOIN
    clubs ON game_events.club_id = clubs.club_id
WHERE
    game_events.type = 'Substitutions'
GROUP BY
    clubs.club_id,clubs.name
ORDER BY
    total_substitutions DESC
LIMIT 1;
