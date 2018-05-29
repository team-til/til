<template>
  <div id="recent-notes">
    <NoteOverview 
      v-for="note in notes"
      v-bind:note="note"
    ></NoteOverview>
  </div>
</template>

<script>
import axios from 'axios';
import VueMarkdown from 'vue-markdown'
import NoteOverview from './NoteOverview'

export default {
  name: 'RecentNotes',
  data() {
    return {
      error: false,
      notes: [
        {
          title: "No Description or Tags Note",
          created_at: "2017-05-18T04:53:01Z",
          created_by: "Otis Redding",
          id: 237,
          editable: false,
          description: "",
          tags: [],
          location: {"organization": true}
        },
        {
          title: "A Test Note",
          created_at: "2018-05-28T22:55:01Z",
          created_by: "Nathan Owsiany",
          id: 234,
          editable: true,
          description: "This is just a sample note. You can type whatever you want here by adding it to the top of your note in the defined section. I think this one should be a little bit longer than it is now.",
          tags: ["test", "rails", "ruby", "testing", "best-practices"],
          location: {"private": true}
        },
        {
          title: "No Tags Note",
          created_at: "2017-05-18T04:53:01Z",
          created_by: "Otis Redding",
          id: 237,
          editable: false,
          description: "This is a description",
          tags: [],
          location: {"organization": true}
        },
        {
          title: "Another Test",
          created_at: "2018-05-18T04:53:01Z",
          created_by: "Otis Redding",
          id: 236,
          editable: false,
          description: "This is just a sample note. You can type whatever you want here by adding it to the top of your note in the defined section.",
          tags: ["test", "golang", "bdd", "testing", "tooling"],
          location: {"team": true, "team_name": "Platform"}
        },
        {
          title: "No Description Note",
          created_at: "2018-05-18T04:53:01Z",
          created_by: "Otis Redding",
          id: 236,
          editable: false,
          description: "",
          tags: ["test", "golang", "bdd", "testing", "tooling"],
          location: {"organization": true}
        },
        {
          title: "Lots of Tags",
          created_at: "2018-05-18T04:53:01Z",
          created_by: "Otis Redding",
          id: 236,
          editable: false,
          description: "This is just a sample note. You can type whatever you want here by adding it to the top of your note in the defined section.",
          tags: ["test", "golang", "bdd", "testing", "tooling", "more", "tags", "are", "always", "better", "right", "I", "think", "so"],
          location: {"organization": true}
        },
      ]
    }
  },
  created() {
    // axios.get('http://localhost:8080/note_previews')
    // .then(response => {
    //     console.log(response)
    //     this.notes = response.data.note_previews
    // })
    // .catch(error => {
    //     this.error = true
    // })
  },
  components: {
    VueMarkdown,
    NoteOverview
  }
}
</script>

<style lang="scss" scoped>

</style>