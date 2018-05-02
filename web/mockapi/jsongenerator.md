# Sessions
```
[
  '{{repeat(15, 20)}}',
  {
    "id": '{{objectId()}}',
    "title": '{{random("Introduction to", "Learn about", "Discover")}} {{company()}}',
    "description": '{{lorem(1, "paragraphs")}}',
    "level": '{{random("blue", "brown", "green")}}',
    "language": '{{random("French", "English", "Creol")}}',
    "format": '{{random("Format A", "Format B", "Format C")}}',
    "room": '{{random("Room 1", "Room 2", "Room 3")}}',
    "speakers": '{{firstName()}} {{surname()}}',
    "ratings_count": '{{integer(20, 40)}}',
    "score": '{{integer(20, 40)}}',
    "feedback_summary": {
        "feedback_1": [
            {"reaction": "A", "count": 0},
            {"reaction": "B", "count": 0}
        ],
        "feedback_2": [
            {"reaction": "A", "count": 0}, 
            {"reaction": "B", "count": 0}
        ],
        "feedback_3": [
            {"reaction": "A", "count": 0}, 
            {"reaction": "B", "count": 0}
        ]
    },
    "start_at": '{{date(new Date(2018, 0, 1), new Date(), "YYYY-MM-ddThh:mm:ss")}}',
    "end_at": '{{date(new Date(2018, 0, 1), new Date(), "YYYY-MM-ddThh:mm:ss")}}',
    "status": true
  }
]
```

# User 

```
[
  '{{repeat(6)}}',
  {
      "id": '{{random("abec", "asdf", "werr","sabc", "asdf", "werd")}}',
      "name": '{{firstName()}} {{surname()}}',
      "level": '{{random("blue", "brown", "green")}}'
  }

]
```