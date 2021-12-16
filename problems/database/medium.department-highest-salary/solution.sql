-- Dialect: MySQL
-- Status: Passed

SELECT sub.Department,
       sub.Employee,
       sub.Salary
FROM (SELECT d.name       Department,
             e.name       Employee,
             e.salary     Salary,
             DENSE_RANK() OVER (PARTITION BY d.name ORDER BY e.salary DESC) "Rank"
      FROM Employee e
               JOIN Department d ON e.departmentId = d.id) AS sub
WHERE sub.Rank = 1
