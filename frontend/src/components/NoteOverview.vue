<template>
  <div v-if="error">
    <h2>Unable to load data</h2>
  </div>
  <div v-else>
    <v-toolbar color="pink">
      <v-toolbar-title class="white--text">Recent Notes</v-toolbar-title>
    </v-toolbar>
    <v-card>
      <v-container
        fluid
        style="min-height: 0;"
        grid-list-lg
      >
        <v-layout row wrap>
          <div v-for="note in notes">
            <v-flex xs12>
              <v-card color="cyan darken-2" class="white--text">
                <v-container fluid grid-list-lg>
                  <v-layout row>
                    <v-flex xs7>
                      <div>
                        <div class="headline">{{note.name}}</div>
                        <br>
                        <vue-markdown>{{note.note_preview}}</vue-markdown>
                      </div>
                    </v-flex>
                  </v-layout>
                </v-container>
              </v-card>
            </v-flex>
          </div>
        </v-layout>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import axios from 'axios';
import VueMarkdown from 'vue-markdown'

export default {
  name: 'NoteOverview',
  data() {
    return {
        error: false,
        notes: []
    }
  },
  created() {
    axios.get('http://localhost:8080/note_previews')
    .then(response => {
        console.log(response)
        this.notes = response.data.note_previews
    })
    .catch(error => {
        this.error = true
    })
  },
  components: {
    VueMarkdown
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>