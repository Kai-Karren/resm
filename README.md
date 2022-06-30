# RESM

RESM which stands for response manager is a service for handling response generation in a dialogue system.

It has been motivated by my previous nlg-server implementation in Java and also acts as a project to use more Go in my personal
projects.

## Use Case

When you have a dialogue system in which you want to separate the actual response generation from the dialogue control you can achieve this with RESM.

E.g. when you are using Rasa and you do not want to deploy a new model for just simple changes to the response text.

RESM is written in Go and is therefore lightweight and fast.

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
