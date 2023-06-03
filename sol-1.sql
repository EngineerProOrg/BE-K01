'''
Bai 1
'''
select stock_name, sum(
case 
    when operation = 'buy' then -price
    when operation = 'sell' then price   
end
)  as capital_gain_loss 
from Stocks 
group by stock_name 

''' 
Bai 2
'''
SELECT 'Low Salary' AS category , COUNT(*) AS accounts_count FROM accounts WHERE income<20000
UNION
SELECT 'Average Salary' AS category , COUNT(*) AS accounts_count FROM accounts WHERE income BETWEEN 20000 and 50000
UNION
SELECT 'High Salary' AS category , COUNT(*) AS accounts_count FROM accounts WHERE income>50000 ;
