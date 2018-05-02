<template>
  <div class="page page-session">
    <div class="logo-wrapper">
      <router-link :to="{ name: 'welcome' }">
        <img src="@/assets/mscc-logo.png" alt="">
      </router-link>
    </div>
    <div class="user-info">
      <span class="title">Logged in as</span>
      <span>User Name</span>
    </div>

    <div class="session-id-wrapper">
      <div class="form-label">Enter your unique id</div>
      <input type="text" placeholder="1234" v-model="userid">
      <a class="submit-wrapper" @click="checkUser()">
        submit
      </a>
      <router-link :to="{name: 'feedback'}" class="submit-wrapper">
        🗝
      </router-link>
    </div>

    <div class="footer">
      Developer Conference 2018
    </div>

  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  components: {},
  computed: {
    userid: {
      get: function() {
        let a = this.$store.getters.getUserId;
        return a;
      },
      set: function(newValue) {
        this.$store.commit("setUserId", newValue);
      }
    },
    ...mapGetters({
      // 'userid': 'getUserId'
    })
  },
  methods: {
    ...mapActions({
      checkUser: "checkUser"
    })
  }
};
</script>

<style lang="scss" scoped>
.logo-wrapper {
  img {
    padding: 15px;
    display: block;
    height: 100%;
    width: auto;
  }
}

.user-info {
  padding: 18px;
  span {
    display: block;
    text-align: right;

    &.title {
      text-transform: uppercase;
      font-size: 12px;
    }
  }
  font-size: 14px;
}

.page-session {
  display: grid;
  grid-template-areas:
    "logo userinfo"
    "session session"
    "footer footer";
  grid-template-columns: 100px 1fr;
  grid-template-rows: 10vh 80vh 10vh;
}

.session-id-wrapper {
  grid-area: session;
  align-items: center;
  justify-content: center;
  align-content: center;
  display: grid;
  .form-label {
    font-size: 16px;
    padding-bottom: 8px;
  }

  input {
    width: 50vw;
    text-align: center;
    max-width: 300px;
    height: 70px;
    font-size: 40px;

    &::-webkit-input-placeholder {
      color: #ddd;
    }
    // alighn
  }

  .submit-wrapper {
    background: white;
    border-radius: 50%;
    width: 100px;
    height: 100px;
    line-height: 100px;
    margin: 40px auto 0;
    text-align: center;
    text-decoration: none;
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

