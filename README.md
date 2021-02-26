# Todo list 
Simple todo list go server
# Running the server

This todo list server is coded using go and mysql. As a pre-requisite to starting the server, please ensure that 
go environment and mysql is installed on your machine. 

Alternatively if go is not installed, you can run the executable binary with `./todo` after applying database changes.

1. Apply the database creation script on your localhost mysql `database/0001-create-table.sql`
   Note: Database connection uses root and no password by default, to make changes change the dsn in `line 13` 
   of `storage/init.go`

2. Once database is created, run server by 

         go run main.go

4. If server is running you should be able to connect to `localhost:8080`

# APIs

### Create todo item  
   
   POST `http://localhost:8080/todos/create`

   cURL Request:

          curl --location --request POST "http://localhost:8080/todos/create" 
          --header "Content-Type: application/json" \
          --data "{
            \"name\": \"Read a book\",
            \"description\": \"Read Harry Potter\"
          }"
        
   Example Response:
   
       {
          "id": 1,
          "name": "Read a book",
          "description": "Read Harry Potter",
          "status": "active"
       }     
        

### Update todo item
   
   Update name/description/status by providing a valid id

   POST `http://localhost:8080/todos/update`
   
   Delete todo
   
      cURL Request:
   
             curl --location --request POST "http://localhost:8080/todos/update" \
             --header "Content-Type: application/json" \
             --data "{
               \"id\":1,
               \"name\": \"Read a book\",
               \"description\": \"Read something\",
               \"status\":\"delete\"
             }"
           
   Example Response:
      
          {
             "id": 1,
             "name": "Read a book",
             "description": "Read something",
             "status": "deleted"
          }     
           

   Mark as done
   
      cURL Request:
   
             curl --location --request POST "http://localhost:8080/todos/update" \
             --header "Content-Type: application/json" \
             --data "{
               \"id\":1,
               \"name\": \"Read a book\",
               \"description\": \"Read something\",
               \"status\":\"done\"
             }"
           
   Example Response:
      
          {
             "id": 1,
             "name": "Read a book",
             "description": "Read something",
             "status": "done"
          }    

### View todo items
   
   GET `http://localhost:8080/todos`

   cURL Request:

          curl --location --request GET "http://localhost:8080/todos" \
          --header "Content-Type: application/json" \
          --data ""
        
   Example Response:
   
       {
          "id": 1,
          "name": "Read a book",
          "description": "Read Harry Potter",
          "status": "active"
       }   