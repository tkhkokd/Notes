

Change Relationship Type
```
MATCH (u:User)-[r]-(c:Car{name:"Audi"})
CREATE (u)-[r2:NEWRELATIONSHIP]->(c)
SET r2 = r
WITH r
DELETE r
```

Delete Node by ID
```
MATCH (u:User) where ID(p)=32
OPTIONAL MATCH (u)-[r]-() //drops p's relations
DELETE r,u
```
