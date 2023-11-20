SELECT stock_name, SUM(
  CASE 
    WHEN operation = 'Sell' THEN price
    ELSE  - price
  END 
) as capital_gain_loss
FROM Stocks
GROUP BY stock_name