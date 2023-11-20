# Write your MySQL query statement below
SELECT 'High Salary' as category,
COUNT(
    account_id
) as accounts_count
FROM Accounts
WHERE income > 50000
UNION
SELECT 'Average Salary' as category,
COUNT(
    account_id
) as accounts_count
FROM Accounts
WHERE income >= 20000 AND income <= 50000
UNION
SELECT 'Low Salary' as category,
COUNT(
    account_id
) as accounts_count
FROM Accounts
WHERE income < 20000

