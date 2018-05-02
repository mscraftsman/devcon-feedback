<template>
  <div class="page page-session">
    <div class="logo-wrapper">
      <router-link :to="{ name: 'welcome' }" >
        <img src="mscc-logo.png" alt="">
      </router-link>
    </div>
    <div class="user-info">
      <span class="title">Logged in as</span>
      <span>User Name</span>
    </div>

    <div class="page-content">
        <!-- <span>{{id}}</span> -->
        <div class="session-title">{{session.title}}</div>

        <div class="descriptions">
            <div class="des-wrap">
                <label>Speaker(s)</label>
                <p>{{ session.speakers }}</p>
            </div>

            <div class="des-wrap">
                <label>Format</label>
                <p>{{ session.format }}</p>
            </div>

            <div class="des-wrap">
                <label>Language</label>
                <p>{{ session.language }}</p>
            </div>

            <div class="des-wrap">
                <label>Room</label>
                <p>{{ session.room }}</p>
            </div>

            <div class="des-wrap">
                <label>Starts at</label>
                <p>{{ session.start_at }}</p>
            </div>

            <div class="des-wrap">
                <label>Ends at</label>
                <p>{{ session.end_at }}</p>
            </div>

            <div class="des-wrap">
                <label>Level</label>
                <p>{{ session.level }}</p>
            </div>

            <div class="des-wrap full">
                <label>Description</label>
                <p>{{ session.description }}</p>
            </div>
        </div>

        <router-link :to="{ name: 'sessions' }" class="btn small" >
          Back
        </router-link>
      
    </div>

    <div class="footer">
      Developer Conference 2018
    </div>


  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
    props: ['id'],
    methods: {
        // ...mapActions(['fetchOneSession'])
    },
    computed: {
        ...mapGetters({
            'sessions': 'getSessions'
        }),
        session: function () {
            let session = _.filter(this.sessions, {id: this.id})[0]
            console.log(session)
            return session
        }
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
    box-sizing: border-box;
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
  grid-template-rows: 10vh 1fr 10vh;
}

.page-content {
  grid-area: session;
}

.session-title {
    color:white;
    padding: 20px 0;
}
.descriptions {
    padding: 0 10px 30px;
    text-align: left;
    display:grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    grid-row-gap: 20px;
    

    .full {
        grid-column: 1 / -1;
    }
    label {
        text-transform: uppercase;
        font-size: 12px;
        color:white;
    }
    p{
        margin:0;
        font-size: 20px;
        color:#222;
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

