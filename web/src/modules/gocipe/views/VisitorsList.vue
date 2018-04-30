<template>
    <v-container>
        <v-toolbar color="transparent" flat>
            <v-toolbar-title class="grey--text text--darken-4 ml-0"><h2>Visitor</h2></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn mr-0 color="primary" :to="{name: 'visitorsEdit', params:{id: 0}}">
                <v-icon dark>add</v-icon> Add
            </v-btn>
        </v-toolbar>
        
        <v-alert :type="message.type === 'E' ? 'error' : message.type" :value="true" v-for="(message, index) in messages" :key="index">
            {{ message.text }}
        </v-alert>

        <v-alert type="info" value="true"  color="primary" outline icon="info" v-if="entities.length === 0">
            No Visitor exist. Would you like to create one now?
            <v-btn :to="{name: 'visitorsEdit', params:{id: 0}}" color="primary">create new</v-btn>
        </v-alert>
        <template v-else>
            <v-text-field mb-4 append-icon="search" label="Search" single-line hide-details v-model="search"></v-text-field>            
            <v-data-table :headers="headers" :items="entities" class="elevation-1" :search="search">
                <template slot="items" slot-scope="props">
					
					<td>{{ props.item.id}}</td>
					
					<td>{{ props.item.code}}</td>
					
					<td>{{ props.item.name}}</td>
					
					<td>{{ props.item.affiliation}}</td>
					
					<td>{{ props.item.email}}</td>
					
					<td>{{ props.item.level}}</td>
					
                    <td class="justify-center layout px-0">
                        <v-btn icon class="mx-0" :to="{name: 'visitorsEdit', params: {'id': props.item.id}  }">
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
                            No Visitor found.
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
		
		{text: "ID", value: "id"},
		
		{text: "Code", value: "code"},
		
		{text: "Name", value: "name"},
		
		{text: "Affiliation", value: "affiliation"},
		
		{text: "", value: "email"},
		
		{text: "Level", value: "level"},
		
        {'text': 'Action', 'value': null}
      ],
      entities: []
    };
  },
  created() {
    axios
      .get("/api/visitors")
      .then(response => {
        this.entities = response.data.entities;
      })
      .catch(error => {
        this.messages = [...this.messages, ...error.response.data.messages];
      });
  }
};
</script>