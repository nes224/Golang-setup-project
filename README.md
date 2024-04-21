# Golang-setup-project

# Why do we need to use DB transaction?
There are 2 main reasons:

1. We want our unit of work to be reliable and consistent, even in case of system failure.
- เราต้องการให้หน่วยการทำงานของเรามีความเชื่อถือได้และสม่ำเสมอ แม้กรณีเกิดความล้มเหลวในระบบ" นั่นคือ เมื่อเกิดข้อผิดพลาดใด ๆ ในระบบ การดำเนินการหนึ่ง ๆ ที่ถูกทำจะต้องยังคงทำงานได้อย่างเสถียรและสม่ำเสมอโดยไม่มีการขัดข้องใด ๆ ในขั้นตอนการดำเนินงานหรือการเก็บรักษาข้อมูลใด ๆ ที่สำคัญในสถานการณ์ที่เกิดข้อผิดพลาดในระบบ

2. We want to provide isolation between programs that access the database concurrently.
เราต้องการให้มีการแยกกันระหว่างโปรแกรมที่เข้าถึงฐานข้อมูลพร้อมกัน" ซึ่งหมายความว่าเราต้องการให้โปรแกรมที่ทำงานพร้อมกันและเข้าถึงข้อมูลในฐานข้อมูลสามารถทำงานได้อย่างอิสระต่อกัน โดยไม่มีผลกระทบต่อผลลัพธ์หรือความสมดุลของข้อมูลที่ถูกเข้าถึงโดยโปรแกรมอื่น ๆ ที่ทำงานในเวลาเดียวกัน

# ACID properties
In order to achieve these 2 goals, a database transaction must satisfy the ACID properties, where:

A is Atomicity, which means either all operations of the transaction complete successfully, or the whole transaction fails, and everything is rolled back, the database is unchanged.
C is Consistency, which means the database state should remains valid after the transaction is executed. More precisely, all data written to the database must be valid according to predefined rules, including constraints, cascades, and triggers.
I is Isolation, meaning all transactions that run concurrently should not affect each other. There are several levels of isolation that defines when the changes made by 1 transaction can be visible to others. We will learn more about it in another lecture.
The last property is D, which stands for Durability. It basically means that all data written by a successful transaction must stay in a persistent storage and cannot be lost, even in case of system failure

# How to run a SQL DB transaction?

It’s pretty simple:

We start a transaction with the BEGIN statement.
Then we write a series of normal SQL queries (or operations).
If all of them are successful, We COMMIT the transaction to make it permanent, the database will be changed to a new state.
Otherwise, if any query fails, we ROLLBACK the transaction, thus all changes made by previous queries of the transaction will be gone, and the database stays the same as it was before the transaction