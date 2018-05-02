<template>
    <div class="page">
        <div class="logo-wrapper">
            <img src="@/assets/mscc-logo.png" alt="">
        </div>
        <div class="dev-con-text">
            Developers Conference 2018
        </div>
        <div class="app-name">
            Sessions <br> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Feedback
        </div>

        <router-link :to="{ name: 'sessionid' }" class="login-btn">
            login to continue
        </router-link>
    </div>
</template>

<script>
export default {
  components: {}
};
</script>


<style lang="scss" scoped>
.logo-wrapper {
  width: 70px;
  height: 70px;
  margin: 0 auto;
  padding: 40px 0 20px;

  img {
    display: block;
    width: 100%;
    height: 100%;
  }
}

.dev-con-text {
  color: white;
  font-size: 18px;
  padding: 30px 0;
}

.app-name {
  color: white;
  // font-size:
  padding: 15vh 0;
}

.login-btn {
  height: 50px;
  background: white;
  max-width: 250px;
  width: 85vw;
  font-size: 18px;
  line-height: 50px;
  margin: 0 auto;
  cursor: pointer;
  display: block;
  text-decoration: none;
  color: black;
  box-shadow: 0px 4px 21px rgba(0, 0, 0, 0.5);
}
</style>
