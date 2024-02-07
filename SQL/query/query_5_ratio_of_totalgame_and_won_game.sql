-- Query-5 : Games won to games played ratio in descending order for each club? top 10 by rank?

SELECT 
    club_id,
    total_game_count,
    total_won_games,
    ratio,
   	ratio_rank
FROM (
    SELECT
    clubs.club_id,
    count(*) AS total_game_count,
    SUM(clubs_games.is_win) AS total_won_games,
    NULLIF(SUM(clubs_games.is_win)::float, 0) /  COUNT(*)::float AS ratio,
    DENSE_RANK() OVER (ORDER BY NULLIF(SUM(clubs_games.is_win)::float, 0) /  COUNT(*)::float desc) AS ratio_rank
FROM
    clubs
JOIN
    clubs_games ON clubs.club_id = clubs_games.club_id
GROUP BY
    clubs.club_id
) AS club_stats
WHERE 
    ratio_rank <=10
ORDER BY 
    ratio_rank;