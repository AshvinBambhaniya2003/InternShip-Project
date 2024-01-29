-- Query-5 : Games won to games played ratio in descending order for each club? top 10 by rank?

SELECT
    clubs.club_id,
    count(*) AS total_game,
    SUM(clubs_games.is_win) AS total_won_games,
    COUNT(*)::float / NULLIF(SUM(clubs_games.is_win), 0)::float AS overall_club_rank_ratio
FROM
    clubs
JOIN
    clubs_games ON clubs.club_id = clubs_games.club_id
GROUP BY
    clubs.club_id
ORDER BY
    overall_club_rank_ratio DESC
LIMIT 10;