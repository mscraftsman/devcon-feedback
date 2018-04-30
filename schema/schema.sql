DROP TABLE IF EXISTS ;

CREATE TABLE  (
	"id" 
	SERIAL
	 NOT NULL
	,
	
	"visitor_id" 
	INT
	 NOT NULL
	,
	
	"reaction_1" 
	CHAR(1)
	 NOT NULL
	,
	
	"reaction_2" 
	CHAR(1)
	 NOT NULL
	,
	
	"reaction_3" 
	CHAR(1)
	 NOT NULL
	,
	
	"reaction_4" 
	TEXT
	 NOT NULL
	,
	
	"created_at" 
	TIMESTAMP
	 NOT NULL
	,
	
	PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS ;

CREATE TABLE  (
	"id" 
	SERIAL
	 NOT NULL
	,
	
	"title" 
	VARCHAR(255)
	 NOT NULL
	,
	
	"description" 
	Text
	 NOT NULL
	,
	
	"level" 
	CHAR(1)
	 NOT NULL
	,
	
	"language" 
	CHAR(2)
	 NOT NULL
	,
	
	"format" 
	CHAR(1)
	 NOT NULL
	,
	
	"room" 
	VARCHAR(255)
	 NOT NULL
	,
	
	"speakers" 
	TEXT
	 NOT NULL
	,
	
	"ratings_count" 
	INT
	 NOT NULL
	,
	
	"score" 
	INT
	 NOT NULL
	,
	
	"reaction_summary" 
	JSONB
	 NOT NULL
	,
	
	"start_at" 
	TIMESTAMP
	 NOT NULL
	,
	
	"end_at" 
	TIMESTAMP
	 NOT NULL
	,
	
	"status" 
	BOOLEAN
	 NOT NULL
	,
	
	PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS ;

CREATE TABLE  (
	"id" 
	SERIAL
	 NOT NULL
	,
	
	"code" 
	CHAR(6)
	 NOT NULL
	,
	
	"name" 
	VARCHAR(255)
	 NOT NULL
	,
	
	"affiliation" 
	VARCHAR(255)
	 NOT NULL
	,
	
	"email" 
	VARCHAR(255)
	 NOT NULL
	,
	
	"level" 
	CHAR(1)
	 NOT NULL
	,
	
	PRIMARY KEY ("id")
);


