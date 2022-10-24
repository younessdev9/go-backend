## Project setup

before yo start running the project you must have go installed into your machine 

### setup env 

for this project I've used mongodb as my database using mongo atlas service (mongosb in the cloud).
to roll the app you either use Mongodb Atlas or use the on your local machine if you have it installed jusr got to ``/pkg/database/setup.go`` 
and change the connection string 


### Using MongoDB Atlas 

first go to your root directory and add ``.env`` file and add ``MONGODB_PASSWORD=*your_mongodb_atlss_password*``
and yeah that's it there you go!!

from the project directory run
```sh
go run cmd/main.go
```
