# RESM

RESM which stands for RESponse Manager is an SDK for handling response generation in a dialogue system. It can
for example be used to build NLG servers for Rasa to separate the response generation from the dialogue control.

It is a personal project in early development that has been motivated by my previous nlg-server implementation in Java
to replicate it in Go to deepen my understanding of Go and potentially to add more features to RESM compared to the Java version.
The Java version I maintain can be found here [NLG-Server](https://github.com/Kai-Karren/nlg-server)

## Use Case

When you have a dialogue system in which you want to separate the actual response generation from the dialogue control you can achieve this with RESM.

E.g. when you are using Rasa and you do not want to deploy a new model for just simple changes to the response text.

RESM is written in Go and is therefore lightweight and fast.

## Get RESM

```bash
go get github.com/Kai-Karren/resm
```

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

The request for the Rasa API has to look as follows:
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

## Response Generators

The Rasa-compatible API provides different so-called Response Generators that handle the response generation of RESM.

### Static Response Generator

The StaticResponseGenerator is the primary used Response Generator that allows you to specify a mapping between
a response name and one response text or a set of response variations. Slot values can be injected in the text
responses with `$slotName` or with `{slotName}`.

### Custom Response Generator

If you want to handle responses by executing custom Go code, you can use the CustomResponseGenerator
which allows you to pass a mapping between response name and a function that should be run to generate the response.
This allows dyamic response generation at runtime with almost unlimited possibilities.

```go
generator := NewCustomResponseGenerator()

exampleHandler := func(request NlgRequest) (NlgResponse, error) {
  return NewNlgResponse("This is a custom response."), nil
}

generator.AddHandler("test", exampleHandler)
```

### Distributed Response Generator

If you want to combine e.g. the Static and the Custom Response Generators in one API, you can do this with the
DistributedResponseGenerator. It handles the routing of the requests to the corresponding generator. If you are
using your own Response Generator implementation please make sure to correctly implement the `HandlesResponse` method.
Currently, the DistributedResponseGenerator queries the generators in the order that have been added each time a request
is handled. This may change in the future with alternative routing strategies or other alternative implementations.

### Your Response Generator

Of course, it is also possible to create your own Response Generator by implementing the Response Generator interface!
