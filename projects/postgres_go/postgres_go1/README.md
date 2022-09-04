# POC inserted 1M data to postgres

we use majestic_million as data source and you can get it here: 
[majestic_million](https://datahub.io/rising.odegua/majestic-millions#resource-majestic_million)

docker-compose -f [./services/postgres.yml](https://github.com/estherjanesamuel/database-playground/blob/main/services/postgres.yml) up -d

- Check pg already created (just in case you don’t have psql / dbms to connect postgres)
- Open http://localhost:5050/ , login using username = admin@admin.com password = root
- Run command at CLI docker network inspect services_default | grep Gateway to get the bridge IP
- Click create new server at pgadmin
- Fill the general tab, name = workshopsearch
- Fill the connection tab, host = bridge IP, port = 5432, username = postgres

If connected, then pg ready to used

