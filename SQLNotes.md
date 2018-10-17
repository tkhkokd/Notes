
Copy a value to another column, WHERE clause needed.
```
UPDATE tablename SET column2=column02
```
Replace string
```
UPDATE tablename
SET column01 = REPLACE(column01, 'stringToReplace', 'newString')
```

Find if there are duplicate value
```
SELECT column01, COUNT(*) c FROM tablename GROUP BY column01 HAVING c > 1;
```


ADD TIMESTAMP

```
ALTER TABLE tablename
ADD COLUMN created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
```
