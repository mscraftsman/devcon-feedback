<template>
  <div class="page">
    <div class="super-menu">
      <div class="dev-con-text">
        Developers Conference 2018
      </div>
      <div class="super-menu-item">
        <router-link :to="{ name: 'sessions' }">
          Schedule
        </router-link>
      </div>
      <!-- <div class="super-menu-item">
                <router-link :to="{ name: 'login' }">
                    Login
                </router-link>
            </div> -->
    </div>

    <!-- <router-link :to="{ name: 'login' }" class="btn">
            ✅ Rate
        </router-link> -->
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import moment from "moment";

export default {
  methods: {
    ...mapActions(["fetchSessions", "fetchSpeakers"]),
    time: function(date) {
      // console.log()
      // return new Date(date).toDateString();
      return moment(date).format("LT");
    }
  },
  computed: {
    ...mapGetters({
      sessions: "getSessions"
    })
  },
  mounted: function() {
    this.fetchSessions();
    this.fetchSpeakers();
  },
  components: {}
};
</script>

<style lang="scss" scoped>
.dev-con-text {
  color: white;
  font-size: 60px;
  padding: 30px 0;
}

.app-name {
  color: white;
  // font-size:
  padding: 15vh 0;
}

.super-menu {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-rows: 1fr 100px 100px 100px;
  grid-gap: 20px;
  //   grid-template-areas:
}

.super-menu-item {
  display: flex;
  align-items: center;
  justify-content: center;

  &:nth-child(1) {
    a {
      //   background-image: url("../assets/schedule_1.jpg");
    }
  }
  &:nth-child(2) {
    a {
      //   background-image: url("../assets/duotone.jpg");
      background: #31e8b7;
      color: black;
    }
  }

  a {
    // height: 200px;
    width: 80%;
    // display: block;
    padding: 30px 0;
    box-shadow: 0 15px 30px 0 rgba(0, 0, 0, 0.11),
      0 5px 15px 0 rgba(0, 0, 0, 0.08);
    background-size: cover;
    text-decoration: none;
    color: white;
    border-radius: 5px;
    background: #2847d9;
  }
}
</style>
