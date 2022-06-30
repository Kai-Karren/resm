# RESM

RESM which stands for response manager is a service for handling response generation in a dialogue system.

## Run

```bash
go run .
```

## Simple API

RESM's most simple API that provided basic features.

### API Request

```json
{
    "response": "utter_example",
    "slots": {
        "name": "John Doe",
        "turns": "4"
    }
}
```

```bash
curl http://localhost:8080/nlg -X POST -d '{
    "response": "utter_example",
    "slots": {
        "name": "John Doe",
        "turns": "4"
    }
}'
```

### API Response

```json
{
    "response": "utter_example",
    "text": "This is an example response from RESM"
}
```

## Rasa Compatible API

TBD
