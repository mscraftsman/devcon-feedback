<template>
    <div class="container">
		<v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Session</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn ml-0 small color="grey" flat :to="{name: 'sessionsList'}">
                <v-icon dark>arrow_back</v-icon> Back
            </v-btn>
        </v-toolbar>
		<v-alert :type="message.type" :value="true" v-for="(message, index) in messages" :key="index">
		{{ message.text }}
		</v-alert>
  
        

        <v-btn color="primary" @click="save()">Save</v-btn>
        <v-btn color="gray" :to="{name: 'sessionsList'}">Cancel</v-btn>
	</div>
</template>
  
<script>
import axios from "axios"

export default {
    props: ["id"],
    created() {
        if (!this.id) {
            return
        }

        axios.get("/api/sessions/" + this.id).then(response => {
            this.id = response.data.entity.id
            
            this.entity.title = response.data.entity.title
            
            
            this.entity.description = response.data.entity.description
            
            
            this.entity.level = response.data.entity.level
            
            
            this.entity.language = response.data.entity.language
            
            
            this.entity.format = response.data.entity.format
            
            
            this.entity.room = response.data.entity.room
            
            
            this.entity.speakers = response.data.entity.speakers
            
            
            this.entity.ratings_count = response.data.entity.ratings_count
            
            
            this.entity.score = response.data.entity.score
            
            
            this.entity.reaction_summary = response.data.entity.reaction_summary
            
            
            this.entity.start_at = response.data.entity.start_at
            
            
            this.entity.end_at = response.data.entity.end_at
            
            
            this.entity.status = response.data.entity.status
            
            
        })
    },
    data() {
        return {
            select: {
                
			},
			dates: {
                
			},
            messages: [],
            entity: {
                
                title : null,
                
                description : null,
                
                level : null,
                
                language : null,
                
                format : null,
                
                room : null,
                
                speakers : null,
                
                ratings_count : null,
                
                score : null,
                
                reaction_summary : null,
                
                start_at : null,
                
                end_at : null,
                
                status : null
                
            }
        }
    },
    watch: {
        
        "select.id.search": function(val) {
            val && this.querySelections("id", "sessions", "/", val)
        },
        
        "select.title.search": function(val) {
            val && this.querySelections("title", "sessions", "/", val)
        },
        
        "select.description.search": function(val) {
            val && this.querySelections("description", "sessions", "/", val)
        },
        
        "select.level.search": function(val) {
            val && this.querySelections("level", "sessions", "/", val)
        },
        
        "select.language.search": function(val) {
            val && this.querySelections("language", "sessions", "/", val)
        },
        
        "select.format.search": function(val) {
            val && this.querySelections("format", "sessions", "/", val)
        },
        
        "select.room.search": function(val) {
            val && this.querySelections("room", "sessions", "/", val)
        },
        
        "select.speakers.search": function(val) {
            val && this.querySelections("speakers", "sessions", "/", val)
        },
        
        "select.ratings_count.search": function(val) {
            val && this.querySelections("ratings_count", "sessions", "/", val)
        },
        
        "select.score.search": function(val) {
            val && this.querySelections("score", "sessions", "/", val)
        },
        
        "select.reaction_summary.search": function(val) {
            val && this.querySelections("reaction_summary", "sessions", "/", val)
        },
        
        "select.start_at.search": function(val) {
            val && this.querySelections("start_at", "sessions", "/", val)
        },
        
        "select.end_at.search": function(val) {
            val && this.querySelections("end_at", "sessions", "/", val)
        },
        
        "select.status.search": function(val) {
            val && this.querySelections("status", "sessions", "/", val)
        }
        
    },
    methods: {
        querySelections(fieldname, endpoint, filter, val) {
            this.select[fieldname].loading = true
            axios.get("/api/" + endpoint + "?" + filter + "-lk=" + encodeURIComponent(val)).then(response => {
                this.select[fieldname].loading = false
                this.select[fieldname].items = response.data.entities.map(function(e) {
                    return { text: e[filter], value: e.id }
                })
            })
        },
        save() {
            if (this.id) {
                axios.put("/api/sessions/" + this.id, this.entity).then(this.saved)
            } else {
                axios.post("/api/sessions", this.entity).then(this.saved)
            }
		},
		saved(response) {
			this.id = response.data.entity.id
			
            this.entity.title = response.data.entity.title
            
            this.entity.description = response.data.entity.description
            
            this.entity.level = response.data.entity.level
            
            this.entity.language = response.data.entity.language
            
            this.entity.format = response.data.entity.format
            
            this.entity.room = response.data.entity.room
            
            this.entity.speakers = response.data.entity.speakers
            
            this.entity.ratings_count = response.data.entity.ratings_count
            
            this.entity.score = response.data.entity.score
            
            this.entity.reaction_summary = response.data.entity.reaction_summary
            
            this.entity.start_at = response.data.entity.start_at
            
            this.entity.end_at = response.data.entity.end_at
            
            this.entity.status = response.data.entity.status
            

			this.messages.push({
				type: "success",
				text: "Session saved successfully"
			})
		}
    }
}
</script>