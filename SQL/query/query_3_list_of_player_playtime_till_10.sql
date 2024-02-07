-- Query-3 : List top 10 players with highest play time. Show their rank as well.

SELECT
    player_name,
    total_minutes_played,
    player_rank
FROM (
    SELECT
        player_name,
        SUM(minutes_played) AS total_minutes_played,
        DENSE_RANK() OVER (ORDER BY SUM(minutes_played) DESC) AS player_rank
    FROM
        appearances
    GROUP BY
        player_id, player_name
) AS player_stats
WHERE
    player_rank <= 10
ORDER BY
    player_rank;