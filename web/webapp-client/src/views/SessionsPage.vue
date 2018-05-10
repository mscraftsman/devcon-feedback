<template>
  <div class="page-all-sessions">
    <div class="tabs-container">
      <div class="tab-items">
        <router-link :to="{ name: 'welcome' }">
          🔙
        </router-link>
      </div>
      <div class="tab-items" :class="{active : getDay(group.groupName) === active}" @click="active = getDay(group.groupName)" v-for="group in sessions" :key="group.groupId" :label="getDay(group.groupName)">
        {{getDay(group.groupName)}}
      </div>
    </div>
    <div class="tabs-content">

      <transition name="fade" mode="out-in">
        <div class="tabs-panel-content" v-if="getDay(group.groupName) === active" v-for="group in sessions" :key="group.groupId">
          <div class="session-panes" v-for="session in group.sessions" :key="session.id">
            <router-link :to="{ name: 'session',  params: { id: session.id }}">
              <div class="date-time">{{ time(session.startsAt) }} - {{ time(session.endsAt) }}</div>
              <div class="session-title">{{ session.title }} </div>
              <div class="session-author">{{ session.speakers[0].name }} - {{ session.room }}</div>
            </router-link>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import moment from "moment";
export default {
  data() {
    return {
      tabs: ["Thursday", "Friday", "Saturday"],
      active: "Thursday"
    };
  },
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
  }
};
</script>

<style lang="scss" scoped>
.bounce-enter-active {
  animation: bounce-in 0.2s;
}
.bounce-leave-active {
  animation: bounce-in 0.2s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.page-all-sessions {
  display: grid;
  grid-template-rows: 50px calc(100vh - 50px - 33px);
}

.tabs-container,
.tabs-content {
  max-width: 900px;
  margin: 0 auto;
  width: 100%;
  //   box-shadow: 0 0 15px rgba(0, 0, 0, 0.5);
}
.tabs-container {
  display: flex;
  justify-content: space-around;
  background: white;

  .tab-items {
    font-size: 18px;
    height: 50px;
    text-transform: uppercase;
    box-sizing: border-box;
    cursor: pointer;

    &.active {
      border-bottom: 3px solid red;
    }
  }
}

.tabs-content {
  .tabs-panel-content {
    // display: none
    color: white;
    height: 100%;
    overflow-y: scroll;
  }
}

.tabs-container {
  text-transform: uppercase;
  border-bottom: 0;
  // background: linear-gradient(
  //   135deg,
  //   rgba(49, 232, 183, 1) 0%,
  //   rgba(40, 71, 217, 1) 70%
  // );
  background: transparent;

  .tab-items {
    background: transparent;
    border: 0;

    color: var(--color-white);
    font-family: var(--font-shentox);
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;

    .cursor {
      background: transparent;
    }
  }
}

.tabs-panel-content {
  background: var(--color-white);
  padding: 0 10px;
  font-family: var(--font-glacial);

  .session-panes {
    font-size: 14px;
    list-style: none;
    margin: 0;
    border-bottom: 1px solid lightgrey;
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

    .date-time {
      text-transform: uppercase;
      font-size: 13px;
      font-family: var(--font-shentox);
      font-weight: 700;
      color: var(--color-green);
      color: var(--color-blue);
    }

    .session-title {
      padding: 5px 0;
      font-size: 20px;
      color: #333;
      font-weight: 500;
      //   text-transform: uppercase;
    }

    .session-author {
      font-size: 14px;
      color: #333;
      text-transform: uppercase;
      font-weight: 300;
    }
  }
}
</style>
