<template>
  <div class="page page-session">

    <div class="page-content" v-if="session">
      <!-- <span>{{id}}</span> -->
      <div class="session-title">{{session.title}}</div>
      <div class="speakers-wrapper" v-if="session.speakers">
        <div class="speaker-wrapper" v-for="speaker in session.speakers" :key="speaker.id">
          <div class="avatar"><img :src="getSpeaker(speaker.id)" alt=""></div>
          <p class="name">{{ speaker.name }}</p>
        </div>
      </div>

      <div class="descriptions">
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
          <p>{{ time(session.startsAt) + getDay(session.startsAt) }} - {{ time(session.endsAt) }}</p>
        </div>

        <div class="des-wrap" v-if="session.level">
          <label>Level</label>
          <p>{{ session.level }}</p>
        </div>
      </div>

      <div class="des-wrap full">
        <label>Description</label>
        <p>{{ session.description }}</p>
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
      console.log(id);
      console.log(this.speakers);
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
  max-width: 900px;
  margin: 0 auto;
  width: 100%;
  margin-top: 50px;
}

.page-content {
  grid-area: session;
  background: white;
  padding: 20px;
}

.session-title {
  color: #333333;
  text-transform: uppercase;
  font-family: var(--font-glacial);
  font-size: 40px;
  font-weight: 700;
  margin: 0 auto;
  padding: 20px 10px;
  text-align: center;
}

.speakers-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;

  .speaker-wrapper {
    --width: 70px;
    display: grid;
    width: 200px;
    grid-template-areas: "avatar name";
    grid-template-columns: var(--width) 1fr;
    align-items: center;
    grid-gap: 10px;
    margin-bottom: 10px;
    margin-right: 20px;
    font-family: var(--font-shentox);
    font-weight: 500;
    text-transform: uppercase;

    .avatar {
      grid-area: avatar;
      width: var(--width);
      height: var(--width);
      border-radius: var(--width);
      box-shadow: 0 0 20px rgba(0,0,0,.1);
      overflow: hidden;

      img {
        width: var(--width);
        height: var(--width);
        border-radius: var(--width);
        display: block;
        object-position: 50% 50%;
        object-fit: cover;
      }
    }

    p {
      grid-area: name;
      text-align: left;
      font-size: 15px;
      color: var(--color-blue);
      font-weight: 700;
    }
  }

}


.descriptions {
  padding: 0 30px 30px;
  text-align: left;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  grid-row-gap: 20px;
  color: black;

  .full {
    grid-column: 1 / -1;
  }
  label {
    text-transform: uppercase;
    font-size: 12px;
    border-bottom: 2px solid white;
    margin-bottom: 10px;
    color: black;
  }
  p {
    margin: 0;
    font-size: 20px;
    // color: #222;
    color: black;
  }
}

.footer {
  grid-area: footer;
  color: black;

  text-align: center;
  font-size: 13px;
  align-self: center;
}
</style>

