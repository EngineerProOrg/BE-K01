
-- cách 1 (cần insert thêm 3 row data, chạy đc ở dev k chạy đc ở leetcode)

insert into Accounts(account_id, income, category) values
(-1, 0,null),
(-1, 20001, null),
(-1, 50001, null);

select 
case 
	when income <20000 then 'Low Salary'
    when income >50000 then 'High Salary'
    else 'Average Salary'
end as category,
count(*) -1 as accounts_count
from Accounts
group by 
CASE 
    WHEN income < 20000 THEN 'Low Salary'
    WHEN income > 50000 THEN 'High Salary'
    ELSE 'Average Salary'
  END;

-- cách 2(chạy đươc trên leetcode k cần insert )

SELECT
    c.category
    , COUNT(a.income) AS accounts_count
FROM
    (
        SELECT 'Low Salary' AS category
        UNION ALL
        SELECT 'Average Salary' AS category
        UNION ALL
        SELECT 'High Salary' AS category
    ) c
LEFT JOIN
    accounts a ON (
        (c.category = 'Low Salary' AND a.income < 20000) OR
        (c.category = 'Average Salary' AND a.income >= 20000 AND a.income <= 50000) OR
        (c.category = 'High Salary' AND a.income > 50000)
    )
GROUP BY
    c.category
 




