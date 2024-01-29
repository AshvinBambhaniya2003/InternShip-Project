-- Query-3 : List top 10 players with highest play time. Show their rank as well.

SELECT
    player_name,
    SUM(minutes_played) AS total_minutes_played,
    DENSE_RANK() OVER (ORDER BY SUM(minutes_played) DESC) AS player_rank
FROM
    appearances
GROUP BY
    player_id, player_name
ORDER BY
    total_minutes_played DESC
LIMIT 10;