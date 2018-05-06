<template>
  <div id="app">
    <transition name="page" mode="out-in">
      <router-view></router-view>
    </transition>
    <div class="error-log">
      <ul>
        <li v-for="error in errors" :class="error.type" :key="error.timestamp">
          {{ error.text }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters({
      errors: "getErrors"
    })
  }
};
</script>


<style lang="scss">
html,
body,
#app,
.page {
  min-height: 100vh;
  /* contain: strict; */
}

.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease-in-out, transform 0.3s ease-in-out;
}
.page-enter {
  opacity: 0;
  transform: translateX(30%);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-30%);
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
  font-family: -apple-system, BlinkMacSystemFont, "avenir next", avenir,
    helvetica, "helvetica neue", Ubuntu, "segoe ui", arial, sans-serif;
}

.page {
  text-align: center;
  /* nesting for the need to test postcss */
  code {
    background-color: #f0f0f0;
    padding: 3px 5px;
    border-radius: 2px;
  }

  max-width: 800px;
  margin: 0 auto;
}

#app {
  background: linear-gradient(166deg, #31e8b7 0%, #2847d9 100%);
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
}

.tabs-component-tabs {
  border: solid 1px #ddd;
  border-radius: 6px;
  margin-bottom: 5px;
}

@media (min-width: 700px) {
  .tabs-component-tabs {
    border: 0;
    align-items: stretch;
    display: flex;
    justify-content: flex-start;
    margin-bottom: -1px;
  }
}

.tabs-component-tab {
  color: #999;
  font-size: 14px;
  font-weight: 600;
  margin-right: 0;
  list-style: none;
}

.tabs-component-tab:not(:last-child) {
  border-bottom: dotted 1px #ddd;
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

@media (min-width: 700px) {
  .tabs-component-tab {
    background-color: #fff;
    border: solid 1px #ddd;
    border-radius: 3px 3px 0 0;
    margin-right: 0.5em;
    transform: translateY(2px);
    transition: transform 0.3s ease;
  }

  .tabs-component-tab.is-active {
    border-bottom: solid 1px #fff;
    z-index: 2;
    transform: translateY(0);
  }
}

.tabs-component-tab-a {
  align-items: center;
  color: inherit;
  display: flex;
  padding: 0.75em 1em;
  text-decoration: none;
}

.tabs-component-panels {
  // padding: 4em 0;
}

@media (min-width: 700px) {
  .tabs-component-panels {
    border-top-left-radius: 0;
    background-color: #fff;
    border: solid 1px #ddd;
    border-radius: 0 6px 6px 6px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
    // padding: 4em 2em;
  }
}
</style>
