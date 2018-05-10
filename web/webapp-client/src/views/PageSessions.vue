<template>
  <div class="page page-session">

    <Tabs @changePage="changePage" indexTab="thursday">
      <TabPanel v-for="group in sessions" :key="group.groupId" :label="getDay(group.groupName)" :hash="getDay(group.groupName)" fontsize="36" tabheight="90" color="red">
        <div class="session-panes" v-for="session in group.sessions" :key="session.id">
          <router-link :to="{ name: 'session',  params: { id: session.id }}">
            <div class="date-time">{{ time(session.startsAt) }} - {{ time(session.endsAt) }}</div>
            <div class="session-title">{{ session.title }} {{ session.id }}</div>
            <div class="session-author">{{ session.speakers[0].name }} - {{ session.room }}</div>
          </router-link>
        </div>
      </TabPanel>
    </Tabs>

    <!-- <tabs :options="{ useUrlFragment: false }">
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
    </tabs> -->

  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import moment from "moment";
import Tabs from "../components/swipe-tab/index.vue";
import TabPanel from "../components/swipe-tab/tab.vue";
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
    },
    changePage(idx) {
      console.log(idx);
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
  components: {
    Tabs,
    TabPanel
  }
};
</script>

<style lang="scss" scoped>

  * {
    box-sizing: border-box;
  }

.page-session {
  // display: grid;
  // grid-template-areas: "page";
  // grid-template-rows: 90vh;
  background: rgba(0,0,0,.05);

  .tabs-container {
    height: 100%;
    display: grid;
    grid-template-rows: 50px 1fr;
    font-family: var(--font-glacial);
    box-shadow: 0 0 15px rgba(0, 0, 0, 0.5);

    .tabs-list {
      text-transform: uppercase;
      border-bottom: 0;
      background: linear-gradient(135deg, rgba(49,232,183, 1) 0%, rgba(40,71,217, 1) 70%);

      .tabs-title {
        background: transparent;
        border: 0;

        a {
          font-size: 30px;
          color: var(--color-white);
          font-family: var(--font-shentox);
          padding: 0;
          height: 50px;
          display: flex;
          align-items: center;
          justify-content: center;
        }
      }
    }

    .tabContent-wrap {
      .tabs-panel-content {
        height: 100%;
        overflow-y: scroll;

        .session-panes {
          a {
            padding: 15px 0;
          }
        }

      }
    }
  }
}

.session-id-wrapper {
  // grid-area: page;
  // align-items: center;
  // justify-content: center;
  // align-content: center;
  // display:grid;
  // grid-template-columns: 5% 1fr 5%;
}

ul.tabs-list {
  // list-style: none;
  // padding: 0;
  // margin: 0;
  // width: 100%;
  // grid-column: 2/3;

  li {
    font-size: 14px;
    list-style: none;
    margin: 0;
    border-bottom: 1px solid lightgrey;
    background: #333;
    text-align: left;

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

  @media (max-width: 480px) {
    .page-session {
      .tabs-container {
        .tabs-list {
          .tabs-title {
            a {
              font-size: 20px;
            }
          }
        }
      }
    }
  }

</style>

