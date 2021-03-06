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
   
   ####Request Fields:
   1. name (string): required
   2. description (string): required
   3. status (string): required
   4. dueDate (string): required (in UTC 0 time format i.e `2021-03-12T15:00:00Z`)
   5. priority (int): required (positive integers only)
   6. userID (int): required (owner of the todo item)
   7. private (bool): optional (mark as private/public, defaults to private)
   

   cURL Request:

          curl --location --request POST "http://localhost:8080/todos/create" 
          --header "Content-Type: application/json" \
          --data "{
            \"name\": \"Read a book\",
            \"description\": \"Read Harry Potter\"
            \"status\": \"in-progress\",
            \"dueDate\": \"2021-03-12T15:22:40.793Z\",
            \"priority\": 1,
            \"userID\": 1,
            \"private\": false
          }"
        
   Example Response:
   
       {
          "id": 1,
          "name": "Read a book",
          "description": "Read Harry Potter",
          "status": "in-progress",
          "dueDate": "2021-03-12T15:22:40.793Z",
          "priority": 1,
          "userID": 2,
          "private": false
       }     
        

### Update todo item
   
   Update name/description/status by providing a valid id

   POST `http://localhost:8080/todos/update`
   
   ####Request Fields:
   1. name (string): required
   2. description (string): required
   3. status (string): required
   4. dueDate (string): required (in UTC 0 time format i.e `2021-03-12T15:00:00Z`)
   5. priority (int): required (positive integers only)
   6. todoID (int): required (id of the todo item)
   7. private (bool): optional (mark as private/public)
   8. delete (bool): optional (mark as delete)
   
   Delete todo
   
      cURL Request:
   
             curl --location --request POST "http://localhost:8080/todos/update" \
             --header "Content-Type: application/json" \
             --data "{
               \"todoID\":1,
               \"name\": \"Read a book\",
               \"description\": \"Read something\",
               \"status\": \"blocked\",
               \"dueDate\": \"2021-03-12T15:22:40.793Z\",
               \"priority\": 2
               \"delete\": true
             }"
           
   Example Response:
      
          {
             "id": 1,
             "name": "Read a book",
             "description": "Read something",
             "status": "blocked",
             "dueDate": "2021-03-12T15:22:40.793Z",
             "priority": 2
             "delete": true
          }     


### View todo items
   
   GET `http://localhost:8080/todos`
   
  ####Request Fields:
  1. currentUserID (int): required (current user id)
  2. userID (int): required (user id of user viewed by current user)
  3. status (string): optional
  4. from (string): optional (in UTC 0 time format i.e `2021-03-12T15:00:00Z`, filters for due date starting from time specified)
  5. priority (int): optional (positive integers only, filters for items with specified priority)
  6. limit (int): optional (used for pagination, specifies page size)
  7. offset (int): optional (used for pagination, specifies offset)

  Note: 
  If currentUserID is equal to userID, meaning the user is viewing his/her own todo list so all non deleted items will be shown. 
  If currentUserID is a friend of userID, only public items of the viewed user will be shown.
  If currentUserID is not a friend of userID, no items will be shown.

   cURL Request:

          curl --location --request GET "http://localhost:8080/todos" \
          --header "Content-Type: application/json" \
          --data "{
          	\"currentUserID\": 1,
          	\"userID\": 2
          }"
        
   Example Response:
   
     [
         {
             "id": 4,
             "name": "Read a book",
             "description": "Read something",
             "status": "in-progress",
             "priority": 2,
             "dueDate": "2021-03-07 01:27:14 +0000 UTC",
             "private": false
         }
     ]  