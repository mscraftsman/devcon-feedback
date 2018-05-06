<template>
    <v-container>
        <v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Feedback</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn mr-0 color="primary" :to="{name: 'api/feedbacksEdit', params:{id: 0}}">
                <v-icon dark>add</v-icon> Add
            </v-btn>
        </v-toolbar>
        
        <v-alert :type="message.type === 'E' ? 'error' : message.type" :value="true" v-for="(message, index) in messages" :key="index">
            {{ message.text }}
        </v-alert>

        <v-alert type="info" value="true"  color="primary" outline icon="info" v-if="entities.length === 0">
            No Feedback exist. Would you like to create one now?
            <v-btn :to="{name: 'api/feedbacksEdit', params:{id: 0}}" color="primary">create new</v-btn>
        </v-alert>
        <template v-else>
            <v-text-field mb-4 append-icon="search" label="Search" single-line hide-details v-model="search"></v-text-field>            
            <v-data-table :headers="headers" :items="entities" class="elevation-1" :search="search">
                <template slot="items" slot-scope="props">
					
					<td>{{ props.item.visitor_id}}</td>
					
					<td>{{ props.item.reaction_1}}</td>
					
					<td>{{ props.item.reaction_2}}</td>
					
					<td>{{ props.item.reaction_3}}</td>
					
					<td>{{ props.item.reaction_4}}</td>
					
					<td>{{ props.item.created_at}}</td>
					
                    <td class="justify-center layout px-0">
                        <v-btn icon class="mx-0" :to="{name: 'api/feedbacksEdit', params: {'id': props.item.id}  }">
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
                            No Feedback found.
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
		
		{text: "Visitor ID", value: "visitor_id"},
		
		{text: "Reaction 1", value: "reaction_1"},
		
		{text: "Reaction 2", value: "reaction_2"},
		
		{text: "Reaction 3", value: "reaction_3"},
		
		{text: "Reaction 4", value: "reaction_4"},
		
		{text: "Created At", value: "created_at"},
		
        {'text': 'Action', 'value': null}
      ],
      entities: []
    };
  },
  created() {
    axios
      .get("/api/api/feedbacks")
      .then(response => {
        this.entities = response.data.entities;
      })
      .catch(error => {
        this.messages = [...this.messages, ...error.response.data.messages];
      });
  }
};
</script>