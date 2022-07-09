# RESM

RESM which stands for RESponse Manager is an SDK for handling response generation in a dialogue system. It can
for example be used to build NLG servers for Rasa to separate the response generation form the dialogue control.

It is a personal project in early development that has been motivated by my previous nlg-server implementation in Java
to replicate it in Go to deepen my understanding of Go and potentially to add more features to RESM compared to the Java version.
The Java version I maintain can be found here [NLG-Server](https://github.com/Kai-Karren/nlg-server)

## Use Case

When you have a dialogue system in which you want to separate the actual response generation from the dialogue control you can achieve this with RESM.

E.g. when you are using Rasa and you do not want to deploy a new model for just simple changes to the response text.

RESM is written in Go and is therefore lightweight and fast.

## Run

```bash
go run .
```

## Tests

To run the test cases recursively to run the tests in the submodules run

```bash
go test -v ./...
```

As a reminder, test files have to end with _test in Go.

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

The Rasa API implementation can act as NLG server following [Rasa NLG](https://rasa.com/docs/rasa/nlg/)
Has been last tested with Rasa 3.2.1

### Rasa API Request

The request for the Rasa API has to look as following:
For more details please see [Rasa NLG](https://rasa.com/docs/rasa/nlg/)

```json
{
"response":"utter_example",
  "arguments":{
    
  },
  "tracker":{
    "sender_id":"user_0",
    "slots":{
      "number": "42"
    },
    "latest_message":{
      "intent":{
        "id":3014457480322877053,
        "name":"greet",
        "confidence":0.9999994039535522
      },
      "entities":[
        
      ],
      "text":"Hello",
      "message_id":"94838d6f49ff4366b254b6f6d23a90cf",
      "metadata":{
        
      },
      "intent_ranking":[
        {
          "id":3014457480322877053,
          "name":"greet",
          "confidence":0.9999994039535522
        }
      ]
    },
    "latest_event_time":1599476297.694504,
    "followup_action":null,
    "paused":false,
    "events":[
      {
        "event":"action",
        "timestamp":1599476297.68784,
        "name":"action_session_start",
        "policy":null,
        "confidence":null
      },
      {
        "event":"session_started",
        "timestamp":1599476297.6878452
      }
    ],
    "latest_input_channel":"rest",
    "active_loop":{
      
    },
    "latest_action_name":"action_listen"
  },
  "channel":{
    "name":"collector"
  }
}
```

### Rasa API Response

```json
{
    "text": "This is an example response from RESM."
}
```

## Examples

Examples and experiments with RESM can be found in this [repo](https://github.com/Kai-Karren/resm-examples)
