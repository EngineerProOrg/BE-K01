-- link problem: https://leetcode.com/problems/capital-gainloss/description/

--------------------SOLVE--------------------

-- solve problem with CASE
SELECT
  stock_name,
  SUM(
    CASE
      WHEN operation = 'Buy' THEN -price
      ELSE price
    END
  ) AS capital_gain_loss
FROM Stocks
GROUP BY stock_name;

-- solve problem with IF
SELECT
  stock_name,
  SUM(
    IF(operation = 'Buy', -price, price)
  ) AS capital_gain_loss
FROM Stocks
GROUP BY stock_name;