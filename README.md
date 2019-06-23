# Go-ToDo-List
Go service to handle day to day To-Do tasks by making use of RabbitMQ & Redis. This service exposes a POST api endpoint which receives POST requests from users to add the todo items to Redis database. The POST request should follow below format: <br/>
````js
Headers:
  Content-Type:application/json
Body:
  {
    "title": "To-Do Title",
    "desc": "To-Do description!!"
  }
````
On successful insertion, the service publishes messages on RabbitMQ messaging queue to inform a node service about presence of data. 


# Steps to Setup
 
 1. Clone the project. <br/>
 
 2. Install the Go package dependencies.
    
 3. Configure following values in conf.json file: <br />
    { <br />
        "PORT" : "< Port on which service should run >", <br />
        "REDIS_NETWORK" : "tcp", <br />
        "REDIS_ADDR" : "< Connection url of Redis server >", <br />
        "RABBIT_CONN_URL" : "< RabbitMQ connection url >", <br />
        "RABBIT_QNAME" : "< Queue name configured (on which service will publish messages) in RabbitMQ >" <br />
    } <br />
    
    Sample config values are as follows: <br />
    ````js
    {
      "PORT" : "8080",
      "REDIS_NETWORK" : "tcp",
      "REDIS_ADDR" : "localhost:6379",
      "RABBIT_CONN_URL" : "amqp://guest:guest@localhost:5672/",
      "RABBIT_QNAME" : "todolistqueue"
    }
    ````
    
 4. Start the service using following command: <br />
    ````bash 
    go run *.go
    ````
<br />
You can browse the apis POST endpoint at http://< Go Service Host Address >:< Configured Port> / todoitem. 
Eg: http://localhost:8080/todoitem 
