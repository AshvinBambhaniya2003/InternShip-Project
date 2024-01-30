-- Query-4 : Which clubs have won most games? list top 10 clubs.

SELECT 
    club_code,
    total_won_games,
    overall_club_rank
FROM (
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
) AS club_stats
WHERE 
    overall_club_rank <=10
ORDER BY 
    overall_club_rank;
