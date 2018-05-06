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
import  { mapGetters } from 'vuex';

export default {
  computed: {
    ...mapGetters({
      'errors' : 'getErrors'
    })
  }
}
</script>


<style lang="scss">

html, body, #app, .page {
  min-height: 100vh;
  /* contain: strict; */
}

.page-enter-active, .page-leave-active {
  transition: opacity .3s ease-in-out, transform 0.3s ease-in-out;
}
.page-enter{
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
    color:black;
    box-shadow: 0px 4px 21px rgba(0, 0, 0, 0.5);

    &.small {
      font-size: 14px;
       --size: 30px;
    }
}

body {
  margin: 0;
  font-size: 2rem;
  font-family: -apple-system, BlinkMacSystemFont,
               'avenir next', avenir,
               helvetica, 'helvetica neue',
               Ubuntu,
               'segoe ui', arial,
               sans-serif;
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
  background: linear-gradient(166deg, #31E8B7 0%, #2847D9 100%);
}

.error-log {
  position:fixed;
  bottom:0;
  width: 100%;

  ul, li {
    padding:0;
    margin:0;
  }

  li {
    font-size: 14px;
    padding: 0 10px;

    &.success {background: lightgreen ; color: green ;}
    &.error {background: lightcoral ; color: red;}
    &.warning {background: lightyellow; color: orange ;}
    &.info {background: lightblue ; color: blue ;}
  }
}
</style>
