# resm
RESM which stands for response manager is a service for handling response generation in a dialogue system.


## Run

```
go run .
```

## API 

```
curl http://localhost:8080/request -X POST -d '{"type":"test"}'
```

```
{
    "type": "test",
    "slots": [
        {"name": "John Doe"},
        {"turns": "4"}
    ]
}
```