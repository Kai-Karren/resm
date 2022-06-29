# resm
RESM which stands for response manager is a service for handling response generation in a dialogue system.


## Run

```
go run .
```

## API 

### API Request
```
{
    "response": "example_response",
    "slots": {
        "name": "John Doe",
        "turns": "4"
    }
}
```


```
curl http://localhost:8080/request -X POST -d '{
    "response": "example_response",
    "slots": {
        "name": "John Doe",
        "turns": "4"
    }
}'
```

### API Response