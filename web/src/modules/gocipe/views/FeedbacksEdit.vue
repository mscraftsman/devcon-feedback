<template>
    <div class="container">
		<v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Feedback</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn ml-0 small color="grey" flat :to="{name: 'api/feedbacksList'}">
                <v-icon dark>arrow_back</v-icon> Back
            </v-btn>
        </v-toolbar>
		<v-alert :type="message.type" :value="true" v-for="(message, index) in messages" :key="index">
		{{ message.text }}
		</v-alert>
  
        

        <v-btn color="primary" @click="save()">Save</v-btn>
        <v-btn color="gray" :to="{name: 'api/feedbacksList'}">Cancel</v-btn>
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

        axios.get("/api/api/feedbacks/" + this.id).then(response => {
            this.id = response.data.entity.id
            
            this.entity.visitor_id = response.data.entity.visitor_id
            
            
            this.entity.reaction_1 = response.data.entity.reaction_1
            
            
            this.entity.reaction_2 = response.data.entity.reaction_2
            
            
            this.entity.reaction_3 = response.data.entity.reaction_3
            
            
            this.entity.reaction_4 = response.data.entity.reaction_4
            
            
            this.entity.created_at = response.data.entity.created_at
            
            
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
                
                visitor_id : null,
                
                reaction_1 : null,
                
                reaction_2 : null,
                
                reaction_3 : null,
                
                reaction_4 : null,
                
                created_at : null
                
            }
        }
    },
    watch: {
        
        "select.visitor_id.search": function(val) {
            val && this.querySelections("visitor_id", "api/feedbacks", "/api/", val)
        },
        
        "select.reaction_1.search": function(val) {
            val && this.querySelections("reaction_1", "api/feedbacks", "/api/", val)
        },
        
        "select.reaction_2.search": function(val) {
            val && this.querySelections("reaction_2", "api/feedbacks", "/api/", val)
        },
        
        "select.reaction_3.search": function(val) {
            val && this.querySelections("reaction_3", "api/feedbacks", "/api/", val)
        },
        
        "select.reaction_4.search": function(val) {
            val && this.querySelections("reaction_4", "api/feedbacks", "/api/", val)
        },
        
        "select.created_at.search": function(val) {
            val && this.querySelections("created_at", "api/feedbacks", "/api/", val)
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
                axios.put("/api/api/feedbacks/" + this.id, this.entity).then(this.saved)
            } else {
                axios.post("/api/api/feedbacks", this.entity).then(this.saved)
            }
		},
		saved(response) {
			this.id = response.data.entity.id
			
            this.entity.visitor_id = response.data.entity.visitor_id
            
            this.entity.reaction_1 = response.data.entity.reaction_1
            
            this.entity.reaction_2 = response.data.entity.reaction_2
            
            this.entity.reaction_3 = response.data.entity.reaction_3
            
            this.entity.reaction_4 = response.data.entity.reaction_4
            
            this.entity.created_at = response.data.entity.created_at
            

			this.messages.push({
				type: "success",
				text: "Feedback saved successfully"
			})
		}
    }
}
</script>