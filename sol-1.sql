-- Solution for Capital Gain/Loss
select
    stock_name,
    sum(
        case
            when operation = "Buy" then - price
            else price
        end
    ) as capital_gain_loss
from
    Stocks
group by
    stock_name;

-- Solution for Count Salary Categories
select
    "Low Salary" as "category",
    count(*) as "accounts_count"
from
    Accounts
where
    Accounts.income < 20000
union
select
    "Average Salary" as "category",
    count(*) as "accounts_count"
from
    Accounts
where
    income >= 20000
    and income <= 50000
union
select
    "High Salary" as "category",
    count(*) as "accounts_count"
from
    Accounts
where
    income > 50000;