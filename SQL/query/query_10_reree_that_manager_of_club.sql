-- Query-10 : Give list of referees who have been once a manager of a team.

SELECT distinct 
		*
FROM  
    (
    SELECT
        g1.game_id,
        g1.referee,
        g2.game_id as home_game_id,
        g2.home_club_manager_name,
        null as away_game_id,
        null as away_club_manager_name
    FROM 
        games g1
    JOIN games g2 ON g1.referee = g2.home_club_manager_name 
    
    UNION ALL

    SELECT
       	g1.game_id,
        g1.referee,
        null as home_game_id,
        null as home_club_manager_name,
        g2.game_id as away_game_id,
        g2.away_club_manager_name
    FROM 
        games g1
    JOIN games g2 ON g1.referee = g2.away_club_manager_name
    )as t1;