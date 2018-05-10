<template>
  <div class="page page-session">

    <div class="page-content" v-if="session">
      <!-- <span>{{id}}</span> -->
      <div class="session-title">{{session.title}}</div>

      <div class="descriptions">
        <div class="des-wrap" v-if="session.speakers">
          <label>Speaker(s)</label>
          <div class="speaker-wrapper" v-for="speaker in session.speakers" :key="speaker.id">
            <div class="avatar"><img :src="getSpeaker(speaker.id)" alt=""></div>
            <p>{{ speaker.name }}</p>
          </div>
        </div>

        <div class="des-wrap" v-if="session.format">
          <label>Format</label>
          <p>{{ session.format }}</p>
        </div>

        <div class="des-wrap" v-if="session.language">
          <label>Language</label>
          <p>{{ session.language }}</p>
        </div>

        <div class="des-wrap">
          <label>Room</label>
          <p>{{ session.room }}</p>
        </div>

        <div class="des-wrap">
          <label>Starts at</label>
          <p>{{ time(session.startsAt) + getDay(session.startsAt) }}</p>
        </div>

        <div class="des-wrap">
          <label>Ends at</label>
          <p>{{ time(session.endsAt) }}</p>
        </div>

        <div class="des-wrap" v-if="session.level">
          <label>Level</label>
          <p>{{ session.level }}</p>
        </div>

        <div class="des-wrap full">
          <label>Description</label>
          <p>{{ session.description }}</p>
        </div>
      </div>

      <router-link :to="{ name: 'sessions' }" class="btn small">
        Back
      </router-link>

    </div>
    <div class="page-content" v-else>
      calling Ish..
    </div>
    <div class="footer">
      Developer Conference 2018
    </div>

  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import moment from "moment";

export default {
  props: ["id"],
  methods: {
    ...mapActions(["fetchSessions", "fetchSpeakers"]),
    getSpeaker: function(id) {
      if (this.speakers.length === 0) {
        this.fetchSpeakers();
      }
      let theSpeaker = this.speakers.filter(speaker => speaker.id === id);
      if (theSpeaker.length > 0) {
        return theSpeaker[0].profilePicture;
      }
    },
    time: function(date) {
      // console.log()
      // return new Date(date).toDateString();
      return moment(date).format("LT");
    },
    getDay: function(str) {
      return str.split(",")[0];
    }
  },
  computed: {
    ...mapGetters({
      sessions: "getSessions",
      speakers: "getSpeakers"
    }),
    session: function() {
      let sessions = this.sessions
        .map(groups => groups.sessions)
        .reduce(function(acc, curr) {
          return [...acc, ...curr];
        }, []);
      let session = _.filter(sessions, { id: this.id })[0];
      // console.log(session);
      return session;
    }
  },
  beforeMount() {
    if (this.$store.state.sessions.length === 0) {
      // console.error("no sessions found");
      this.fetchSessions();
    } else {
      // console.info("sessions found !");
    }
  }
};
</script>

<style lang="scss" scoped>
.page-session {
  display: grid;
  grid-template-areas:
    "session session"
    "footer footer";
  grid-template-columns: 100px 1fr;
  grid-template-rows: 1fr 10vh;
}

.page-content {
  grid-area: session;
  background: #333;
}

.session-title {
  color: white;
  padding: 20px 10px;
}

.speaker-wrapper {
  --width: 100px;
  display: grid;
  grid-template-areas: "avatar name";
  grid-template-columns: var(--width) 1fr;
  align-items: center;
  grid-gap: 10px;
  margin-bottom: 10px;

  .avatar {
    grid-area: avatar;
    width: var(--width);
    height: var(--width);
    border-radius: var(--width);
    overflow: hidden;
    img {
      width: 100%;
      height: 100%;
      display: block;
    }
  }

  p {
    grid-area: name;
    text-align: center;
  }
}

.descriptions {
  padding: 0 30px 30px;
  text-align: left;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  grid-row-gap: 20px;

  .full {
    grid-column: 1 / -1;
  }
  label {
    text-transform: uppercase;
    font-size: 12px;
    color: white;
    border-bottom: 2px solid white;
    margin-bottom: 10px;
  }
  p {
    margin: 0;
    font-size: 20px;
    // color: #222;
    color: white;
  }
}

.footer {
  grid-area: footer;
  color: white;
  text-align: center;
  font-size: 13px;
  align-self: center;
}
</style>

