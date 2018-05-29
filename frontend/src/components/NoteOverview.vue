<template>
  <div class="note-preview white">
    <div class="title-section clearfix">
      <h3 class="note-title">{{note.title}}</h3>
      <div class="note-action-btns">
      <a v-if="note.location.private" href="/">
        <v-chip label small color="red lighten-3" text-color="blue-grey darken-2">
          <v-icon left class="note-tag">lock</v-icon>Private
        </v-chip>
      </a>
      <a v-if="note.location.organization" href="/">
        <v-chip label small color="green lighten-3" text-color="blue-grey darken-3">
          <v-icon left class="note-tag">location_city</v-icon>Organization
        </v-chip>
      </a>
      <a v-if="note.location.team" href="/">
        <v-chip label small color="orange lighten-3" text-color="blue-grey darken-3">
          <v-icon left class="note-tag">group</v-icon>{{note.location.team_name}}
        </v-chip>
      </a>
      <v-icon>star</v-icon>
      </div>
    </div>
    <div class="note-metadata">
      <p><v-icon size=18>person</v-icon><a rel="/">{{note.created_by}}</a> - <span>{{toRelativeTime(note.created_at)}}</span></p>
    </div>
    <div class="note-description">
      <p>{{note.description}}</p>
    </div>
    <div class="note-tags">
      <a v-for="tag in note.tags" href="/">
        <v-chip small color="grey" text-color="white">
            <v-icon left class="note-tag">label</v-icon>{{tag}}
        </v-chip>
      </a>
    </div>
    <v-btn block color="teal" dark class="view-note-btn">View Note</v-btn>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  name: 'RecentNotes',
  props: ['note'],
  methods: {
    toRelativeTime: function(time) {
      return moment(time).fromNow()
    }
  },
}
</script>

<style lang="scss">
.note-preview {
  border-radius: 2px;
  min-height: 100px;
  box-shadow: rgba(10, 10, 10, 0.1) 0px 2px 3px, rgba(10, 10, 10, 0.1) 0px 0px 0px 1px;
  margin-bottom: 16px;
  padding: 24px 24px 16px 24px;

  .note-title {
    float: left;
    margin-bottom: 2px;
    font-weight: 200;
    font-size: 24px;
  }

  .note-metadata {
    margin-left: -3px;
    margin-bottom: 5px;
    font-weight: 200;
    font-size: 14px;

  }

  .note-action-btns {
    float: right;
  }

  .note-description {
    margin-top: 15px;
    font-weight: 200;
    font-style: italic;
  }

  .note-tags {
    margin-top: 20px;
    margin-left: -3px;
  }

  .view-note-btn {
    margin-top: 20px;
  }
}
</style>