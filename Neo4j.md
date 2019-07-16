

Change Relationship Type
```
MATCH (u:User)-[r]-(c:Car{name:"Audi"})
CREATE (u)-[r2:NEWRELATIONSHIP]->(c)
SET r2 = r
WITH r
DELETE r
```
