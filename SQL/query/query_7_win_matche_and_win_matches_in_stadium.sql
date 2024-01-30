-- Query-7 : Which team has won more matches in which stadium? give team/club name, stadium name, number of total matches they won, number of matches they won in that stadium.

SELECT
    club_id,
    stadium_name,
    COUNT(*) AS total_matches_won,
    COUNT(CASE WHEN hosting = 'Home' AND is_win = 1 THEN 1 END) AS matches_won_in_stadium
FROM
    (
        SELECT
            g.home_club_id AS club_id,
            g.stadium AS stadium_name,
            'Home' AS hosting,
            CASE WHEN g.home_club_goals > g.away_club_goals THEN 1 ELSE 0 END AS is_win
        FROM
            games g

        UNION ALL

        SELECT
            g.away_club_id AS club_id,
            g.stadium AS stadium_name,
            'Away' AS hosting,
            CASE WHEN g.away_club_goals > g.home_club_goals THEN 1 ELSE 0 END AS is_win
        FROM
            games g
    ) AS t1
GROUP BY
    club_id,
    stadium_name
ORDER BY
    total_matches_won desc,matches_won_in_stadium desc;
