**Problem:** https://leetcode.com/problems/nth-highest-salary/

**Table:** `Employee`

```
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+

id is the primary key column for this table.
Each row of this table contains information about the salary of an employee.
```

Write an SQL query to report the `Nth` highest salary from the `Employee` table. If there is no `Nth` highest salary, the query should report `null`.

The query result format is in the following example.

**Example 1:**

```
Input: n = 2

Employee table:

+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+

Output: 

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
```
**Example 2:**

```
Input: n = 2

Employee table:

+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+

Output: 

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| null                   |
+------------------------+
```
