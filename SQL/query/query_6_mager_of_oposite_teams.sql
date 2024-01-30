-- Query-6 : Which managers have been manager of opponent teams? (Rephrased: Identify managers who have transitioned to become the manager of teams that were once opponents of the team they initially managed.)

select DISTINCT
    t1.own_manager_name,
    t1.club_id,
    t1.opponent_id,
    t1.game_id AS original_game_id,
    t2.game_id AS new_game_id
FROM
    clubs_games t1
JOIN
    clubs_games t2 ON t1.own_manager_name IS NOT NULL
                      AND t1.own_manager_name <> ''
                      AND t2.opponent_manager_name IS NOT NULL
                      AND t2.opponent_manager_name <> ''
                      AND t1.opponent_id = t2.club_id 
where t1.own_manager_name = t2.own_manager_name;