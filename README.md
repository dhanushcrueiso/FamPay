# FamPay
Backend Youtube Api consumption
- Note the change in postman port for both the cases 
- Wait for initial 10seconds for the api to fetch the data 

# TWO WAYS : 
# Main Branch , local without docker :

- Create a database named youtube_data , in your local postgres 
- Run the migration.sql file to create table 
- use : go run . in terminal to run the server
- Use Postman , localhost:3007/v1/GetData , with query param pg (page num), q(search filter on title) 
-Done 

# Docker branch , local with docker :

- run : docker compose up --build
-  open terminal and enter psql 
- run : docker exec -it <container_id> psql -U postgres (replace the container_id with the db container id from your docker image)
- run : CREATE DATABASE youtube_data;
- run : \c youtube_data;(to connect to the databse we created just now inside the db container)
-  run : CREATE TABLE videos (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    publish_time TIMESTAMP,
    thumbnails TEXT[]
);
- go to postman : localhost:3000/v1/GetData , with query param pg (page num), q(search filter on title) 
- Done 



#Setup -Running docker ENV:

