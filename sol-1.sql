-- capital_gain_loss
select stock_name, 
SUM(
  CASE
    WHEN operation = "Buy" THEN price * -1
    ELSE price 
  END
) AS capital_gain_loss 
from Stocks 
group by stock_name;

--count_salary_categories.sql
SELECT type.category AS category , 
COUNT(account_id) AS accounts_count
FROM (
    SELECT 'Low Salary' AS category
    UNION
    SELECT 'Average Salary' AS category
    UNION
    SELECT 'High Salary' AS category
) AS type
LEFT JOIN accounts AS accounts
ON  type.category = (
    CASE
        WHEN accounts.income > 50000 THEN "High Salary"
        WHEN accounts.income >= 20000 THEN "Average Salary"
        ELSE "Low Salary"
    END
) 
group by type.category;