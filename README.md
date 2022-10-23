# bookstore_oauth-api

do using ddd

## cassendra

run cassendra in docker container

`describe keyspaces;`
`docker exec -it bookstore_oauth-api_cassandra_1 cqlsh`

`Create KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};`

`Use oauth;`

`Create Table access_tokens( access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);`

`select * from access_tokens where acess_token='sdfj';`