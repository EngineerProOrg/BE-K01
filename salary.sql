SELECT
    CASE
        WHEN income < 20000 THEN 'Low Salary'
        WHEN income >= 20000 AND income <= 50000 THEN 'Average Salary'
        WHEN income > 50000 THEN 'High Salary'
        END AS category,
    COUNT(*) AS accounts_count
FROM
    Accounts
GROUP BY
    CASE
        WHEN income < 20000 THEN 'Low Salary'
        WHEN income >= 20000 AND income <= 50000 THEN 'Average Salary'
        WHEN income > 50000 THEN 'High Salary'
        END;

#optimize
create table Accounts
(
    account_id int auto_increment
        primary key,
    income     int null
);

create index Accounts_income_index
    on Accounts (income);

