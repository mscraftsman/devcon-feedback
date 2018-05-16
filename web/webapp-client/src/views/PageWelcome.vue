<template>
  <div class="page">
    <div class="super-menu">
      <div class="dev-con-text">
        <div class="text">Developers Conference</div>
        <div class="year">2018</div>
      </div>
      <div class="buttons-wrapper">

        <div class="super-menu-item">
          <router-link :to="{ name: 'sessions' }">
            Schedule
          </router-link>
        </div>
        <div class="super-menu-item">
          <router-link :to="{ name: 'speakers' }">
            Speakers
          </router-link>
        </div>

        <div class="super-menu-item meetup" v-if="!user.status">
          <a href="/b/login">
            Login with meetup
          </a>
        </div>
        <div class="super-menu-item meetup" v-else>
          <a @click="triggerLogout">
            🔌 Logout
          </a>
        </div>
      </div>
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
    ...mapActions(["fetchSessions", "fetchSpeakers", "triggerLogout"]),
    time: function(date) {
      // console.log()
      // return new Date(date).toDateString();
      return moment(date).format("LT");
    }
  },
  computed: {
    ...mapGetters({
      sessions: "getSessions",
      user: "getUser"
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
.page {
  // background: linear-gradient(
  //     135deg,
  //     rgba(49, 232, 183, 1) 0%,
  //     rgba(40, 71, 217, 0.8) 70%
  //   ),
  //   url("../assets/home-bg-2.jpg");
  // background-size: 100%, cover;
  // background-position: center center;
}

.dev-con-text {
  color: white;
  font-size: 56px;
  padding: 30px 0;
  font-family: var(--font-shentox);
  font-weight: 700;
  text-transform: uppercase;
  padding: var(--gutter);
  display: flex;
  align-items: center;
  align-content: center;
  flex-wrap: wrap;
  justify-content: center;
  flex-direction: column;
  grid-area: intro;

  .year {
    display: block;
    align-items: center;
    justify-content: center;
    font-family: var(--font-shentox);
    text-transform: uppercase;
    color: var(--color-white);
    font-size: 44px;
    padding: calc(var(--gutter) / 1.5);
    height: 40px;
    line-height: 45px;
    font-weight: 700;
    background: linear-gradient(135deg, #31e8b7, #2847d9);
    border-radius: 5px;
    -webkit-box-shadow: 0 10px 50px rgba(0, 0, 0, 0.2);
    box-shadow: 0 10px 50px rgba(0, 0, 0, 0.2);
  }
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
  grid-template-rows: 1fr auto;
  grid-auto-rows: auto;
  grid-template-columns: 1fr 4fr 1fr;
  grid-gap: 20px;
  // align-content: center;
  grid-template-areas:
    ". intro ."
    ". button .";
}

.buttons-wrapper {
  grid-area: button;
  display: grid;
  // align-content: center;
  justify-content: center;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  grid-gap: 30px;
}

@media screen and (max-width: 600px) {
  .buttons-wrapper {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  }
}

.super-menu-item {
  display: flex;
  // align-items: center;
  // justify-content: center;
  font-family: var(--font-shentox);
  text-transform: uppercase;
  font-weight: 800;
  // max-width: 250px;

  span {
    font-size: 10px;
  }

  &:nth-child(1) {
    a {
      //   background-image: url("../assets/schedule_1.jpg");
    }
  }
  // &:nth-child(3) {
  //   filter: saturate(0);
  //   opacity: 0.5;
  //   a {
  //     //   background-image: url("../assets/duotone.jpg");
  //     background: #31e8b7;
  //     color: black;
  //   }
  // }

  a {
    // height: 200px;
    width: 100%;
    display: block;
    padding: 20px 0;
    margin-bottom: 20px;
    box-shadow: 0 15px 30px 0 rgba(0, 0, 0, 0.3),
      0 5px 15px 0 rgba(0, 0, 0, 0.08);
    background-size: cover;
    text-decoration: none;
    color: white;
    border-radius: 5px;
    background: #2847d9;
    font-size: 24px;
  }

  &.meetup {
    a {
      background: #fff;
      color: #f64060;
      font-size: 16px;
      align-self: center;
    }
  }
}

@media (max-width: 480px) {
  .dev-con-text {
    font-size: 30px;
  }
}
</style>
