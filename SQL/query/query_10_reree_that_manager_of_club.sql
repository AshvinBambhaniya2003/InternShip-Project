-- Query-10 : Give list of referees who have been once a manager of a team.

SELECT DISTINCT
    g1.game_id,
    g1.referee,
    g1.home_club_id,
    g1.home_club_manager_name,
    g1.away_club_id,
    g1.away_club_manager_name
FROM games g1
JOIN games g2 ON g1.referee = g2.home_club_manager_name OR g1.referee = g2.away_club_manager_name
WHERE g1.referee IS NOT null;