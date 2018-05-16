<template>
  <div id="app">
    <div class="header">
      <div class="logo-wrapper">
        <router-link :to="{ name: 'welcome' }">
          <img src="@/assets/logo.svg" alt="">
        </router-link>
      </div>
      <div class="user-info" v-if="user.status">
        <span class="title">Logged in as {{ user.data.name }}</span>
      </div>
    </div>

    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  computed: {
    ...mapGetters({
      errors: "getErrors",
      user: "getUser"
    })
  },
  methods: {
    ...mapActions(["checkUser", "triggerLogout"])
  },
  mounted() {
    // see if user is logged in
    this.checkUser();
  }
};
</script>


<style lang="scss">
html,
body,
#app,
.page {
  min-height: 90vh;
  /* contain: strict; */
}

.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease-in-out, transform 0.3s ease-in-out;
}
.page-enter {
  opacity: 0;
  // transform: translateX(30%);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-30%);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.btn {
  --size: 50px;
  height: var(--size);
  background: white;
  max-width: 250px;
  width: 85vw;
  font-size: 18px;
  line-height: var(--size);
  margin: 0 auto;
  cursor: pointer;
  display: block;
  text-decoration: none;
  color: black;
  box-shadow: 0px 4px 21px rgba(0, 0, 0, 0.5);

  &.small {
    font-size: 14px;
    --size: 30px;
  }
}

body {
  margin: 0;
  font-size: 2rem;
  font-family: "Nunito", -apple-system, BlinkMacSystemFont, "avenir next",
    avenir, helvetica, "helvetica neue", Ubuntu, "segoe ui", arial, sans-serif;
  background: linear-gradient(
      135deg,
      rgba(49, 232, 183, 1) 0%,
      rgba(40, 71, 217, 0.8) 70%
    ),
    url("./assets/home-bg-2.jpg");
  background-size: 100%, cover;
  background-attachment: fixed;
  background-position: center center;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.page {
  text-align: center;
  /* nesting for the need to test postcss */
  grid-area: page;
  code {
    background-color: #f0f0f0;
    padding: 3px 5px;
    border-radius: 2px;
  }

  // max-width: 800px;
  width: 100%;
  margin: 0 auto;
}

#app {
  // background: linear-gradient(166deg, #31e8b7 0%, #2847d9 100%);
  // background: linear-gradient(166deg, #333 0%, #31e8b7 100%);

  display: grid;
  grid-template-areas:
    "header"
    "page";
  grid-template-rows: 33px calc(100vh - 33px);
}

.header {
  grid-area: header;
  display: grid;
  background: #333;

  grid-template-areas: "logo userinfo";
  grid-template-columns: 1fr;
  // justify-items: center;

  .logo-wrapper {
    height: 100%;

    a {
      display: block;
      height: 100%;
    }
    img {
      padding: 5px 15px;
      display: block;
      height: 33px;
      width: auto;
      box-sizing: border-box;
    }
  }

  .user-info {
    // padding: 18px;
    span {
      display: block;
      text-align: right;

      &.title {
        text-transform: uppercase;
        font-size: 12px;
        line-height: 33px;
        color: white;
        padding-right: 20px;
      }
    }
    font-size: 14px;
  }
}

.error-log {
  position: fixed;
  bottom: 0;
  width: 100%;

  ul,
  li {
    padding: 0;
    margin: 0;
  }

  li {
    font-size: 14px;
    padding: 0 10px;

    &.success {
      background: lightgreen;
      color: green;
    }
    &.error {
      background: lightcoral;
      color: red;
    }
    &.warning {
      background: lightyellow;
      color: orange;
    }
    &.info {
      background: lightblue;
      color: blue;
    }
  }
}

.tabs-component {
  margin: 0;
  display: grid;
  grid-template-areas:
    "top"
    "bottom";
  grid-template-rows: 8vh 82vh;
}

.tabs-component-tabs {
  // border: solid 1px #ddd;
  // border-radius: 6px;
  // margin-bottom: 5px;
  height: 8vh;
}

.tabs-component-tabs {
  grid-area: top;
}
.tabs-component-panels {
  grid-area: bottom;
}

@media (max-width: 768px) {
  .tabs-component {
    grid-template-rows: 82vh 8vh;
  }
  .tabs-component-tabs {
    grid-area: bottom;
  }
  .tabs-component-panels {
    grid-area: top;
  }
}
.tabs-component-tabs {
  border: 0;
  align-items: stretch;
  display: flex;
  justify-content: space-between;
  // margin-bottom: -1px;
}
// }

ul.tabs-component-tabs {
  margin: 0 !important;
  padding: 0;
}
.tabs-component-tab {
  color: #999;
  // font-size: 14px;
  font-weight: 600;
  margin-right: 0;
  font-size: 18px;
  list-style: none;
}

.tabs-component-tab:not(:last-child) {
  // border-bottom: dotted 1px #ddd;
}

.tabs-component-tab:hover {
  color: #666;
}

.tabs-component-tab.is-active {
  color: #000;
}

.tabs-component-tab.is-disabled * {
  color: #cdcdcd;
  cursor: not-allowed !important;
}

// @media (min-width: 700px) {
.tabs-component-tab {
  // background-color: #fff;
  // border: solid 1px #ddd;
  // border-radius: 3px 3px 0 0;
  // margin-right: 0.5em;
  transform: translateY(2px);
  transition: transform 0.3s ease;
}

.tabs-component-tab.is-active {
  border-bottom: solid 10px #31e8b7;
  border-bottom: solid 10px #2847d9;
  z-index: 2;
  // transform: translateY(0);
}
// }

.tabs-component-tab-a {
  align-items: center;
  color: inherit;
  display: flex;
  padding: 0.75em 1em;
  text-decoration: none;
}

.tabs-component-panels {
  // margin: 10px;
}

.tabs-component-panel {
  max-height: 82vh;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
}

// @media (min-width: 700px) {
.tabs-component-panels {
  // border-top-left-radius: 0;
  // background-color: #fff;
  // border: solid 1px #ddd;
  // border-radius: 0 6px 6px 6px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
  // padding: 4em 2em;
}
// }
</style>
