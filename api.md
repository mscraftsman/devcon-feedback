# Preface
Documentation of end points accessible via the API         

# Authentication
Authentication will be via JSON web tokens (JWT).

Token can be created by providing a unique `visitor.code`.

Token will expire after `6 hours`.

Token structure will be:
```json
{
    "id": 0,
    "name": "",
    "level: "" //level code
}
```

**Level Codes**

Code | Level   | Notes
:---:|---------|----------
V    | Visitor | General Visitor
A    | Admin   | Administrator - can access privileged information/actions

# Resources
## Session
Represents a conference session.

```json
{
    "id": 0,
    "title": "",
    "description": "",
    "level": "",
    "language": "",
    "format": "",
    "room": "",
    "speakers": "",
    "ratings_count": 0, //number of people who rated
    "score": 75, //overall score in percentage
    "feedback_summary": {
        "feedback_1": [
            {"reaction": "A", "count": 0}, //feedback code and count
            {"reaction": "B", "count": 0}
        ],
        "feedback_2": [
            {"reaction": "A", "count": 0}, //feedback code and count
            {"reaction": "B", "count": 0}
        ],
        "feedback_3": [
            {"reaction": "A", "count": 0}, //feedback code and count
            {"reaction": "B", "count": 0}
        ]
    },
    "start_at": "2018-01-31T13:01:10.001Z",
    "end_at": "2018-01-31T14:01:10.001Z",
    "status": true //whether session is still on - false for cancelled
}
```

**Level Codes**

Code | Level
:---:|---------------------------
O    | Introductory and overview
I    | Intermediate
A    | Advanced
E    | Expert
V    | Vendor

**Language Codes**

Code | Language
:---:|----------
en   | English
fr   | French

**Format Codes**

Code | Format
:---:|---------
S    | Session

**Notes**
> `speakers` is a `string` of comma separated values for multiple 
speakers
> `ratings_count` will be computed
> `rating` will be computed

## Feedback
Represents a single feedback

```json
{
    "id": 0,
    "visitor": "",
    "feedback_1": "", //Reaction Code
    "feedback_2": "", //Reaction Code
    "feedback_3": "", //Reaction Code
    "message": "",
    "created_at": "2018-01-31T14:01:10.001Z"
}
```

**Feedback Reaction Codes**

Feedback     | Code | Meaning
-------------|:----:|----
`feedback_1` | L    | Like

## Message
Represents a message from server
```json
{
    "type": "", //type code
    "text": ""
}
```

**Type Codes**

Code | Type
:---:|---------
D    | Debug
E    | Error
S    | Success

# Routes
## POST /auth
Create an authentication token

**Request Body**
```json
{
    "code": "" //unique visitor.code
}
```

**Response Body**
```json
{
    "status": true,
    "messages": [] //list of messages
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Token generated
400       | Request Body invalid
401       | Code not found
403       | User access denied

## GET /sessions
Return a list of sessions open for voting

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages
    "sessions": [] //list of session resources
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok

**Optional query string parameters:**
- `?status=all` : return all sessions
- `?status=upcoming` : return all sessions that are not yet closed but not yet started
- `?status=closed` : return all sessions that are closed

## POST /sessions/{id}/feedback
Create feedback for a session

**Request Header**
```
Authorization: Bearer <jwt>
```

**Request Body**
```json
{
    "reaction_1": "",
    "reaction_2": "",
    "reaction_3": "",
    "message": "",
    "created_at": "2018-01-31T14:01:10.001Z"
}
```

**Response Body**
```json
{
    "status": true,
    "messages": [] //list of messages
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Feedback saved
400       | Request Body invalid
401       | Token verification failed 
403       | Feedback not allowed for this session
404       | Session not found

## GET /feedback/me
List feedback posted by current visitor

**Request Header**
```
Authorization: Bearer <jwt>
```

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages
    "sessions": [
        {
            "session_id": 0,
            "title": "",
            "feedback": {
                "feedback_1": "", //Reaction Code
                "feedback_2": "", //Reaction Code
                "feedback_3": "", //Reaction Code
                "message": "",
                "created_at": "2018-01-31T14:01:10.001Z"
            }
        }
    ]
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Ok
401       | Token verification failed 

## Get /sessions/{id}
Get a session details

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages
    "session": {} //single session resource
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
404       | Session not found

## GET /sessions/{id}/feedback
List feedback for a specific session

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages
    "feedback": [] //list of feedback resources
}
```

**Notes:**
> In each feedback resource, `visitor` may be withheld - equal to `null`; if request does not have `admin` privileges 

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
404       | Session not found

# Schema 

## Visitor

Field            | Type           | Constraints
-----------------|----------------|----------
id               | `INT`          | `PRIMARY KEY` `Serial`
code             | `CHAR(6)`      | `UNIQUE` `INDEXED`
name             | `VARCHAR(255)` | 
level            | `CHAR(1)`      |

## Session

Field            | Type           | Constraints
-----------------|----------------|----------
id               | `INT`          | `PRIMARY KEY` `Serial`
speaker_id       | `INT`          | `FOREIGN KEY visitor(id)`
title            | `VARCHAR(255)` |         
description      | `Text`         |        
level            | `CHAR(1)`      |    
language         | `CHAR(2)`      |        
format           | `CHAR(1)`      |      
room             | `VARCHAR(255)` |         
speakers         | `TEXT`         |     
ratings_count    | `INT`          |         
score            | `INT`          | 
feedback_summary | `JSONB`        |              
start_at         | `TIMESTAMP`    |          
end_at           | `TIMESTAMP`    |        
status           | `BOOLEAN`      |  

## Feedback

Field            | Type        | Constraints
-----------------|-------------|----------
id               | `INT`       | `PRIMARY KEY` `Serial`
visitor_id       | `INT`       | `FOREIGN KEY visitor(id)` 
feedback_1       | `CHAR(1)`   |     
feedback_2       | `CHAR(1)`   |     
feedback_3       | `CHAR(1)`   |     
message          | `TEXT`      |   
created_at       | `TIMESTAMP` |        
