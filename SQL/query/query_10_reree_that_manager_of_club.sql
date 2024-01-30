-- Query-10 : Give list of referees who have been once a manager of a team.

SELECT distinct 
		*
FROM  
    (
    SELECT
        g1.game_id,
        g1.referee,
        g2.home_club_id,
        g2.home_club_manager_name,
        g2.away_club_id,
        g2.away_club_manager_name
    FROM 
        games g1
    JOIN games g2 ON g1.referee = g2.home_club_manager_name 
    
    UNION ALL

    SELECT
       	g1.game_id,
        g1.referee,
        g2.home_club_id,
        g2.home_club_manager_name,
        g2.away_club_id,
        g2.away_club_manager_name
    FROM 
        games g1
    JOIN games g2 ON g1.referee = g2.away_club_manager_name
    )as t1;