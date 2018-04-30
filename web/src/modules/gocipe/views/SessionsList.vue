<template>
    <v-container>
        <v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Session</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn mr-0 color="primary" :to="{name: 'sessionsEdit', params:{id: 0}}">
                <v-icon dark>add</v-icon> Add
            </v-btn>
        </v-toolbar>
        
        <v-alert :type="message.type === 'E' ? 'error' : message.type" :value="true" v-for="(message, index) in messages" :key="index">
            {{ message.text }}
        </v-alert>

        <v-alert type="info" value="true"  color="primary" outline icon="info" v-if="entities.length === 0">
            No Session exist. Would you like to create one now?
            <v-btn :to="{name: 'sessionsEdit', params:{id: 0}}" color="primary">create new</v-btn>
        </v-alert>
        <template v-else>
            <v-text-field mb-4 append-icon="search" label="Search" single-line hide-details v-model="search"></v-text-field>            
            <v-data-table :headers="headers" :items="entities" class="elevation-1" :search="search">
                <template slot="items" slot-scope="props">
					
					<td>{{ props.item.id}}</td>
					
					<td>{{ props.item.title}}</td>
					
					<td>{{ props.item.description}}</td>
					
					<td>{{ props.item.level}}</td>
					
					<td>{{ props.item.language}}</td>
					
					<td>{{ props.item.format}}</td>
					
					<td>{{ props.item.room}}</td>
					
					<td>{{ props.item.speakers}}</td>
					
					<td>{{ props.item.ratings_count}}</td>
					
					<td>{{ props.item.score}}</td>
					
					<td>{{ props.item.reaction_summary}}</td>
					
					<td>{{ props.item.start_at}}</td>
					
					<td>{{ props.item.end_at}}</td>
					
					<td>{{ props.item.status}}</td>
					
                    <td class="justify-center layout px-0">
                        <v-btn icon class="mx-0" :to="{name: 'sessionsEdit', params: {'id': props.item.id}  }">
                            <v-icon color="teal">edit</v-icon>
                        </v-btn>
                    </td>
                </template>

                <template slot="no-data">
                    <v-flex ma-4>
                        <v-alert slot="no-results" :value="true" color="primary" outline icon="info" v-if="search.length > 0">
                        Your search for "{{ search }}" found no results.
                        </v-alert>
                        <v-alert slot="no-results" :value="true" color="primary" outline icon="info" v-else>
                            No Session found.
                        </v-alert>
                    </v-flex>
                </template>
            </v-data-table>
        </template>
    </v-container>
</template>

<script>
import axios from "axios"
export default {
  data() {
    return {
      messages: [],
      search: "",
      headers: [
		
		{text: "Id", value: "id"},
		
		{text: "Title", value: "title"},
		
		{text: "Description", value: "description"},
		
		{text: "Level", value: "level"},
		
		{text: "Language", value: "language"},
		
		{text: "Format", value: "format"},
		
		{text: "Room", value: "room"},
		
		{text: "Speakers", value: "speakers"},
		
		{text: "Ratings", value: "ratings_count"},
		
		{text: "Score", value: "score"},
		
		{text: "Reaction", value: "reaction_summary"},
		
		{text: "StartAt", value: "start_at"},
		
		{text: "EndAt", value: "end_at"},
		
		{text: "Status", value: "status"},
		
        {'text': 'Action', 'value': null}
      ],
      entities: []
    };
  },
  created() {
    axios
      .get("/api/sessions")
      .then(response => {
        this.entities = response.data.entities;
      })
      .catch(error => {
        this.messages = [...this.messages, ...error.response.data.messages];
      });
  }
};
</script>