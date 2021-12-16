-- Dialect: MySQL
-- Status: Passed

SELECT email AS Email
FROM Person
GROUP BY email
HAVING count(email) > 1
