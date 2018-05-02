<template>
  <div id="app">
    <transition name="fade">
      <router-view></router-view>
    </transition>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

html,
body,
#app,
.page {
  min-height: 100vh;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
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
</style>
