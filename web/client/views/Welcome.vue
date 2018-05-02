<template>
  <div class="page">
    <div class="logo-wrapper">
      <img src="mscc-logo.png" alt="">
    </div>
    <div class="dev-con-text">
        Developers Conference 2018
    </div>
    <!-- <div class="app-name">
        Sessions <br>
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Feedback
    </div> -->

    <!-- <a @click="$store.dispatch('fetchSessions')" class="btn" >
        fetch sessions
    </a>
      -->

    <router-link :to="{ name: 'sessions' }" class="btn" >
      🕒 Schedule
    </router-link>

    <router-link :to="{ name: 'login' }" class="btn" >
      🔑 Login
    </router-link>
    

    <router-link :to="{ name: 'login' }" class="btn" >
        ✅ Rate
    </router-link>
  </div>
</template>

<script>

export default {
  components: {}
}
</script>


<style lang="scss" scoped>
 .logo-wrapper {
     width: 70px;
     height: 70px;
     margin: 0 auto;
     padding-top: 5vh;
      
    img {
        display: block;
        width: 100%;
        height: 100%;
    }
 }

.dev-con-text {
    color:white;
    font-size: 18px;
    padding: 30px 0;
}

.app-name{
    color:white;
    // font-size: 
    padding: 15vh 0;
}
</style>
