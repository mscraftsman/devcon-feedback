<template>
  <div class="page page-session">
    <div class="logo-wrapper">
      <router-link :to="{ name: 'login' }" >
        <img src="mscc-logo.png" alt="">
      </router-link>
    </div>
    <div class="user-info">
      <span class="title">Logged in as</span>
      <span>User Name</span>
    </div>

    <div class="session-id-wrapper">
      <div class="form-label">Enter the Session ID</div>
      <input type="number" placeholder="1234">
      <router-link :to="{name: 'feedback'}" class="submit-wrapper">
        &raquo;
      </router-link>
    </div>

    <div class="footer">
      Developer Conference 2018
    </div>


  </div>
</template>

<script>

export default {
  components: {
  }
}
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
    display:block;
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
  display:grid;
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
    height:100px;
    line-height: 100px;
    margin: 40px auto 0;
    text-align: center;
  }
}

.footer {
  grid-area: footer;
  color:white;
  text-align: center;
  font-size: 13px;
  align-self: center;
}
</style>

