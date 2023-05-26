select
    t1.category,
    count(t2.category) as accounts_count
from (
        SELECT 'Low Salary' AS category
        UNION ALL
        SELECT
            'Average Salary' AS category
        UNION ALL
        SELECT
            'High Salary' AS category
    ) as t1
    LEFT JOIN (
        SELECT
            CASE
                WHEN income < 20000 THEN 'Low Salary'
                WHEN income >= 20000
                AND income <= 50000 THEN 'Average Salary'
                ELSE 'High Salary'
            END AS category
        FROM Accounts
    ) AS t2 ON t1.category = t2.category
group by t2.category