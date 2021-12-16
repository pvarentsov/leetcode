-- Dialect: MySQL
-- Status: Passed

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
RETURN (SELECT sub.salary AS SecondHighestSalary
        FROM (SELECT salary,
                     DENSE_RANK() OVER (ORDER BY salary DESC) AS num
              FROM Employee e
             ) AS sub
        WHERE sub.num = N LIMIT 1
  );
END

