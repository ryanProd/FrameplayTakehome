# Frameplay Takehome
Setup:
1. Downlaod repo to local
2. Copy and paste .env file to root directory, since it has been .gitignore'd from source repo
3. In root directory, run "go run main.go"
4. The retrieved input from the database, and JSON response from API Gateway/Lambda should be displayed to STDOUT and also visible at localhost:3000. 

If you have additional questions please reach out and I will try to answer them. If you have concerns that parts of the prompt were not covered in this solution, please reach out and I will try to fix them. 
<br>

## Thought Process
My thought process was to build a Go app with integrations to other services. I began by thinking about the flow of data and starting drawing some diagrams. The data had to be persisted somewhere, the Go app would retrieve that data, do some things with it, and then send that data onwards to a HTTP based API. <br>
I wanted to test out integrations with the Go app, so my next step was to decide what services to use. I chose a Postgres service [neon](https://neon.tech) to serve as the database, as it has a generous free tier. Despite this being a practice programming assignemnt, I would chose Postgres in many real use cases because of its scalability in quantity of data it can manage and also number of concurrent clients it can have. <br>
For the HTTP API, I went with AWS API Gateway with a Lambda integration because it was what I was familiar with. In a real use case, these services could be good candidates because AWS has good availability, good scaling, and probably competitive pricing. <br>
For the Go app itself, I tried to create packages that aligned with functionality. For example, the config package acts as an accessor for the .env variables, database package provides util functions for connecting and querying to the database, structs package holds data structures used throughout the app, and data package has a method to validate the data retrieved from the database. <br>
I tried to be aware of error handling throughout the app. <br>

## Design Decisions
### Database
Although this is a practice programming assignment, I thought it was good practice to have a .env file that held variables needed to form the Connection String when opening the database. This increases visibility, and when .env is added to .gitignore, the password does not have to be visibile in the source repo. I will attach the .env when I turn in the assignment. There is a config package that acts as an accesor for the .env file.<br>
The Postgres service [neon](https://neon.tech) is a standard cloud PostgresSQL where I created a table and inserted some practice data. <br>
Because the data is moving through the app, I created a struct User to represent it. <br>
There is also a database package that handles initially connecting to the database. The ConnectDB() returns a pointer to the sql.DB struct, which represents access to the neon database in this case. After connecting, I can pass this sql.DB struct pointer to other Database functions, such as QueryDBforUsers(). This function returns a list of User structs from the database.<br>

### Data
After retrieving the data from the database. I use a data util function in the data package to do some validation. <br>

### HTTP Request
In a real world use case, I think it would be useful to create a http request handling package for similar apps to reduce code redunduncy and handle all kinds of request methods, not just POST. I opted to just write the http request code in main() for this practice assignment. <br>
Once the data is validated, it is printed to STDOUT. The http request is then assembled, with the retrieved input data as part of the body. The POST request is sent to the API Gateway endpoint and the response is decoded and printed to STDOUT as well. <br>

### Go Fiber
I was checking out Go Fiber, although it does not have any significant role in this project. After running main.go, you can see the retrieved input and JSON response at localhost:3000. 


