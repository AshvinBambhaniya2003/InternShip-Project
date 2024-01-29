-- Query-6 : Which managers have been manager of opponent teams? (Rephrased: Identify managers who have transitioned to become the manager of teams that were once opponents of the team they initially managed.)

SELECT DISTINCT
    m1.own_manager_name,
    m1.club_id,
    m1.opponent_id,
    m1.game_id AS original_game_id,
    m2.game_id AS new_game_id
FROM
    clubs_games m1
JOIN
    clubs_games m2 ON m1.opponent_manager_name = m2.own_manager_name
WHERE
    m1.opponent_id = m2.club_id;