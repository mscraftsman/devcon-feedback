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
    "email": "",
    "affiliation": "",
    "name": "",
    "level": "", //level code
    "sessions": [] // list of session.id marked for attendance
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
    "reaction_summary": {
        "reaction_1": [
            {"reaction": "A", "count": 0}, //reaction code and count
            {"reaction": "B", "count": 0}
        ],
        "reaction_2": [
            {"reaction": "A", "count": 0}, //reaction code and count
            {"reaction": "B", "count": 0}
        ],
        "reaction_3": [
            {"reaction": "A", "count": 0}, //reaction code and count
            {"reaction": "B", "count": 0}
        ]
    },
    "attendance": 0, //number of people expected to attend
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
    "reaction_1": "", //Reaction Code
    "reaction_2": "", //Reaction Code
    "reaction_3": "", //Reaction Code
    "reaction_4": "",
    "created_at": "2018-01-31T14:01:10.001Z"
}
```

**Feedback Reaction Codes**

Feedback  | Code | Meaning
----------|:----:|----
`*`       | L    | Like

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

## Profile
```json
{
    "name": "", //visitor.name
    "affiliation": "", //visitor.affiliation
    "email": "" //visitor.email
}
```

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
    "messages": [], //list of messages,
    "profile": "" //profile
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Token generated
400       | Request Body invalid
401       | Code not found
403       | User access denied

## GET /me
Return current visitor details

**Request Header**
```
Authorization: Bearer <jwt>
```

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages,
    "profile": {} //profile
}
```

## GET /visitors
Get list of visitors. **Only admin will have access to this route**

**Request Header**
```
Authorization: Bearer <jwt>
```

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages,
    "visitors": {} //profile
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
401       | Token verification failed

## POST /visitors/{id}
Update visitor details. **Only admin will have access to this route**

**Request Header**
```
Authorization: Bearer <jwt>
```

**Request Body**
```json
{
    "email": "",
    "affiliation": "",
    "name": ""
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
400       | Request Body invalid
401       | Token verification failed
404       | Session not found

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages,
    "profile": {} //profile
}
```

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

## POST /sessions
Create a new session. **Accessible only to admin level**

**Request Header**
```
Authorization: Bearer <jwt>
```

**Request Body**
```json
{
    "title": "",
    "description": "",
    "level": "",
    "language": "",
    "format": "",
    "room": "",
    "speakers": "",
    "start_at": "2018-01-31T13:01:10.001Z",
    "end_at": "2018-01-31T14:01:10.001Z",
    "status": true //whether session is still on - false for cancelled
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
400       | Request Body invalid
401       | Token verification failed
404       | Session not found

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages,
    "session": {} //session
}
```  

## PUT /sessions/{id}
Update session. **Accessible only to admin level**

**Request Header**
```
Authorization: Bearer <jwt>
```

**Request Body**
```json
{
    "title": "",
    "description": "",
    "level": "",
    "language": "",
    "format": "",
    "room": "",
    "speakers": "",
    "start_at": "2018-01-31T13:01:10.001Z",
    "end_at": "2018-01-31T14:01:10.001Z",
    "status": true //whether session is still on - false for cancelled
}
```

**Response Codes**

HTTP Code | Meaning
:--------:|-----------
200       | Status Ok
400       | Request Body invalid
401       | Token verification failed
404       | Session not found

**Response Body**
```json
{
    "status": true,
    "messages": [], //list of messages,
    "session": {} //session
}
```  


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
    "reaction_4": "",
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
            "feedback": {}//feedback object
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
affiliation      | `VARCHAR(255)` | 
email            | `VARCHAR(255)` | 
level            | `CHAR(1)`      |

## Session

Field            | Type           | Constraints
-----------------|----------------|----------
id               | `INT`          | `PRIMARY KEY` `Serial`
title            | `VARCHAR(255)` |         
description      | `Text`         |        
level            | `CHAR(1)`      |    
language         | `CHAR(2)`      |        
format           | `CHAR(1)`      |      
room             | `VARCHAR(255)` |         
speakers         | `TEXT`         |     
ratings_count    | `INT`          |         
score            | `INT`          | 
reaction_summary | `JSONB`        |              
start_at         | `TIMESTAMP`    |          
end_at           | `TIMESTAMP`    |        
status           | `BOOLEAN`      |  

## Feedback

Field            | Type        | Constraints
-----------------|-------------|----------
id               | `INT`       | `PRIMARY KEY` `Serial`
visitor_id       | `INT`       | `FOREIGN KEY visitor(id)` 
reaction_1       | `CHAR(1)`   |     
reaction_2       | `CHAR(1)`   |     
reaction_3       | `CHAR(1)`   |     
reaction_4       | `TEXT`      |   
created_at       | `TIMESTAMP` |        

## Attendance

Field            | Type        | Constraints
-----------------|-------------|----------
id               | `INT`       | `PRIMARY KEY` `Serial`
session_id       | `INT`       | `FOREIGN KEY session(id)` 
visitor_id       | `INT`       | `FOREIGN KEY visitor(id)` 
