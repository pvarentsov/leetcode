-- Dialect: MySQL
-- Status: Passed

SELECT MAX(sub.salary) AS SecondHighestSalary
from (SELECT salary,
             DENSE_RANK() over (ORDER BY salary desc) AS num
      FROM Employee e
     ) AS sub
WHERE sub.num = 2
LIMIT 1
