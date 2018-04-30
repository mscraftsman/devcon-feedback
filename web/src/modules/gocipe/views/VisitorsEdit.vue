<template>
    <div class="container">
		<v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Visitor</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn ml-0 small color="grey" flat :to="{name: 'visitorsList'}">
                <v-icon dark>arrow_back</v-icon> Back
            </v-btn>
        </v-toolbar>
		<v-alert :type="message.type" :value="true" v-for="(message, index) in messages" :key="index">
		{{ message.text }}
		</v-alert>
  
        

        <v-btn color="primary" @click="save()">Save</v-btn>
        <v-btn color="gray" :to="{name: 'visitorsList'}">Cancel</v-btn>
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

        axios.get("/api/visitors/" + this.id).then(response => {
            this.id = response.data.entity.id
            
            this.entity.code = response.data.entity.code
            
            
            this.entity.name = response.data.entity.name
            
            
            this.entity.affiliation = response.data.entity.affiliation
            
            
            this.entity.email = response.data.entity.email
            
            
            this.entity.level = response.data.entity.level
            
            
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
                
                code : null,
                
                name : null,
                
                affiliation : null,
                
                email : null,
                
                level : null
                
            }
        }
    },
    watch: {
        
        "select.id.search": function(val) {
            val && this.querySelections("id", "visitors", "/", val)
        },
        
        "select.code.search": function(val) {
            val && this.querySelections("code", "visitors", "/", val)
        },
        
        "select.name.search": function(val) {
            val && this.querySelections("name", "visitors", "/", val)
        },
        
        "select.affiliation.search": function(val) {
            val && this.querySelections("affiliation", "visitors", "/", val)
        },
        
        "select.email.search": function(val) {
            val && this.querySelections("email", "visitors", "/", val)
        },
        
        "select.level.search": function(val) {
            val && this.querySelections("level", "visitors", "/", val)
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
                axios.put("/api/visitors/" + this.id, this.entity).then(this.saved)
            } else {
                axios.post("/api/visitors", this.entity).then(this.saved)
            }
		},
		saved(response) {
			this.id = response.data.entity.id
			
            this.entity.code = response.data.entity.code
            
            this.entity.name = response.data.entity.name
            
            this.entity.affiliation = response.data.entity.affiliation
            
            this.entity.email = response.data.entity.email
            
            this.entity.level = response.data.entity.level
            

			this.messages.push({
				type: "success",
				text: "Visitor saved successfully"
			})
		}
    }
}
</script>