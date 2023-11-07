### FEED AGGREGATOR SERVER




## DEVELOPMENT CYCLE

Since almost all the code is done, the only additions would typically be new endpoints to manage data. The project utilises the package dependencies goose to migrate the tables into the postgres DATABASE and sqlc creates functions automatically by using native SQL queries.

To add an endpoint follow the path:-
1. If the endpoint requires a new table create a new migration in the [sql/schema](./sql/schema) directory.
2. Navigate to [sql/schema](./sql/schema) directory and run command 
```bash
goose postgres postgres://userName:password@localhost:5432/aggregator up
```
This will run the migrations specified in the **Up** section of the migration.

3. Now we generate the SQL query to perform a particular function. Could be creation, deletion, updation anything. The query has to be added in the [sql/queries](./sql/queries) directory.

4. In the root directory run 
```bash
sqlc generate
```
If the query is syntactically correct this would generate automatic functions as specified with the given metadata name.

5. Generate models in [models](./models.go) to specify what all information to send back when the particular api is called.

6. Now generate a handler file eg [Feed Follow Handler](./handle_feed_follow.go). This would contain the business logic to handle the request and return the result.
7. Include the handler with a endpoint and HTTP Method.
8. Run 
```bash
go build && ./aggregator
``` 
to start the server.


Here is the [Postman Collection](./Go%20Aggregator.postman_collection.json) to Interact with the API.
