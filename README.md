# ToDo list

## Rodar a API

    git clone https://github.com/Jhonjoao/api-tasks-go.git
    
    cd api-tasks-go
    
    go run .
    

### Rotas da API

   GET http://localhost:8000/task/pending
    
    example body response:
    [
        {
            "ID": 1,
            "name": "Test",
            "description": "Just a test with insomnia",
            "email": "test@test.com",
            "status": "pending",
            "TimesFinished": 0
        }
    ]
   
   GET http://localhost:8000/task/finished
   
    example body response:
    [
        {
            "ID": 1,
            "name": "Test",
            "description": "Just a test with insomnia",
            "email": "test@test.com",
            "status": "finished",
            "TimesFinished": 0
        }
    ]
   
   POST http://localhost:8000/task
   
    example body request:
    {
        "name": "Test",
        "description": "Just a test with insomnia",
        "email": "test@test.com"
    }
    
    example body response:
    {
        "ID": 1,
        "name": "Test",
        "description": "Just a test with insomnia",
        "email": "test@test.com",
        "status": "pending",
        "TimesFinished": 0
    }
   
   PATCH http://localhost:8000/task/{taskId}
   
    example body request:
    {
	    "status": "finished"
    }
    
    example body response: 
    {
        "ID": 1,
        "name": "Test",
        "description": "Just a test with insomnia",
        "email": "test@test.com",
        "status": "finished",
        "TimesFinished": 0
    }
   
   DELETE http://localhost:8000/task/{taskId}
   
    no body response

