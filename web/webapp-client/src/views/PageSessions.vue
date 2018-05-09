<template>
  <div class="page page-session">
    <tabs :options="{ useUrlFragment: false }">
      <tab :name="getDay(group.groupName)" v-for="group in sessions" :key="group.groupId">
        <ul>
          <li v-for="session in group.sessions" :key="session.id">
            <router-link :to="{ name: 'session',  params: { id: session.id }}">
              <div class="date-time">{{ time(session.startsAt) }} - {{ time(session.endsAt) }}</div>
              <div class="session-title">{{ session.title }} {{ session.id }}</div>
              <div class="session-author">{{ session.speakers[0].name }} - {{ session.room }}</div>
            </router-link>
          </li>
        </ul>
      </tab>
    </tabs>

  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import moment from "moment";

export default {
  methods: {
    ...mapActions(["fetchSessions"]),
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
      sessions: "getSessions"
    })
  },
  mounted: function() {
    this.fetchSessions();
  },
  components: {}
};
</script>

<style lang="scss" scoped>
.page-session {
  // display: grid;
  // grid-template-areas: "page";
  // grid-template-rows: 90vh;
  background: white;
}

.session-id-wrapper {
  // grid-area: page;
  // align-items: center;
  // justify-content: center;
  // align-content: center;
  // display:grid;
  // grid-template-columns: 5% 1fr 5%;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
  width: 100%;
  grid-column: 2/3;

  li {
    font-size: 14px;
    list-style: none;
    margin: 0;
    border-bottom: 1px solid lightgrey;
    // color:white;
    background: #333;
    text-align: left;
    padding: 0 20px;

    .date-time {
      color: gray;
      text-transform: uppercase;
      font-size: 12px;
      font-family: "Courier New", Courier, monospace;
      font-weight: 600;
    }

    .session-title {
      padding: 5px 0;
      font-size: 20px;
      font-weight: bold;
    }

    .session-author {
      font-size: 14px;
      color: gray;
    }

    a {
      display: block;
      padding: 8px 0;
      text-decoration: none;
      color: white;
      transition: transform 0.2s;

      &:hover {
        transition: transform 0.2s;
        transform: translateX(5px);
      }
    }
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

