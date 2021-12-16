-- Dialect: MySQL
-- Status: Passed

SELECT DISTINCT sub.lead0 ConsecutiveNums
FROM (SELECT num          lead0,
             LEAD(num, 1) OVER (ORDER BY id ASC) lead1,
             LEAD(num, 2) OVER (ORDER BY id ASC) lead2
      FROM Logs
     ) AS sub
WHERE sub.lead0 = sub.lead1
  AND sub.lead1 = sub.lead2
