-- Query-2: Which team/club has the highest total minutes of playtime among all teams/clubs?

SELECT
    player_current_club_id,
    SUM(minutes_played) AS total_minutes_played
FROM
    appereances
GROUP BY
    player_current_club_id
ORDER BY
    total_minutes_played DESC
LIMIT 1;