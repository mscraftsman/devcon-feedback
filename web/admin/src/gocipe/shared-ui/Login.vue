<template>
  <div class="login-container" v-if="authStatus() == 'error' || authStatus() == 'logged out'">
    <div class="login-message-wrapper">
      <h1>You are not authorised to access this page. <br>Please login if you
        want to continue.</h1>
    </div>
    <div>
      <v-btn color="primary" large class="mt-4" :href="authURL">
        Login
      </v-btn>
    </div>

  </div>
</template>

<script>
import authMixins from "@/gocipe/mixins/auth.js";

export default {
  mixins: [authMixins]
};
</script>

<style lang="scss" scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  flex-direction: column;
  padding: 10px;
}
.login-message-wrapper {
  h1 {
    text-align: center;
    color: #bcbcbc;
    text-transform: uppercase;
  }
}
/*.login-button-wrapper {*/
/*margin: 20px 0;*/
/*button {*/
/*background: #0d47a1;*/
/*border: 1px solid #0d47a1;*/
/*color: #fff;*/
/*font-size: 20px;*/
/*padding: 10px 30px;*/
/*border-radius: 5px;*/
/*transition: all 0.08s ease-in-out;*/
/*text-transform: uppercase;*/
/*font-weight: 500;*/

/*&:hover {*/
/*background: #072656;*/
/*}*/

/*&:focus {*/
/*outline: none;*/
/*}*/
/*}*/
/*}*/
</style>