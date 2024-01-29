-- Query-4 : Which clubs have won most games? list top 10 clubs.

SELECT
    clubs.club_code,
    SUM(clubs_games.is_win) AS total_won_games,
    DENSE_RANK() OVER (ORDER BY SUM(clubs_games.is_win) DESC) AS overall_club_rank
FROM
    clubs
JOIN
    clubs_games ON clubs.club_id = clubs_games.club_id
GROUP BY
    clubs.club_code
ORDER BY
    total_won_games DESC
LIMIT 10;