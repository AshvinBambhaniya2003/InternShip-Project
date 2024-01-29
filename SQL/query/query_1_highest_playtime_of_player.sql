-- Query-1 : Which player has the highest total minutes of playtime among all players?

SELECT
    player_name,
    SUM(minutes_played) AS total_minutes_played
FROM
    appearances
GROUP BY
    player_name
ORDER BY
    total_minutes_played DESC
LIMIT 1;