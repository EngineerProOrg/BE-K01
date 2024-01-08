select "Low Salary" category, sum(case when income < 20000 then 1 else 0 end) accounts_count from accounts
union
select "Average Salary" category, sum(case when income >= 20000 and income <= 50000 then 1 else 0 end) accounts_count from accounts
union
select "High Salary" category, sum(case when income > 50000 then 1 else 0 end) accounts_count from accounts