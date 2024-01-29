-- Query-9 : Which player is used more time as substitution?

SELECT
    players.name,
    COUNT(*) AS total_substitutions
FROM
    game_events 
JOIN
    players ON game_events.player_id = players.player_id
WHERE
    game_events.type = 'Substitutions'
GROUP BY
    players.player_id, players.name
ORDER BY
    total_substitutions DESC
LIMIT 1;