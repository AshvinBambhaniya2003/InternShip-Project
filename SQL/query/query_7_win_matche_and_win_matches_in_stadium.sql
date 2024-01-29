-- Query-7 : Which team has won more matches in which stadium? give team/club name, stadium name, number of total matches they won, number of matches they won in that stadium.

SELECT
    clubs.club_id ,
    clubs.stadium_name,
    COUNT(*) AS total_matches_won,
    SUM(CASE WHEN clubs_games.is_win = 1 THEN 1 ELSE 0 END) AS matches_won_in_stadium
FROM
    clubs
JOIN
    clubs_games ON clubs.club_id = clubs_games.club_id
GROUP BY
    clubs.club_id , clubs.stadium_name
ORDER BY
    matches_won_in_stadium DESC
LIMIT 10;