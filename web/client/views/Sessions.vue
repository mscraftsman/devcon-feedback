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

    <div class="session-id-wrapper">
        <ul>
            <li v-for="session in sessions" :key="session.id">
                <router-link :to="{ name: 'session',  params: { id: session.id }}">
                    <div class="date-time">{{ time(session.start_at) }} - {{ time(session.start_at) }}</div>
                    <div class="session-title">{{ session.title }}</div>
                    <div class="session-author">{{ session.speakers }} - {{ session.room }}</div>
                </router-link>
            </li>
        </ul>
    </div>

  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import moment from 'moment';

export default {
    methods: {
        ...mapActions(['fetchSessions']),
        time: function(date){
          // console.log()
          // return new Date(date).toDateString();
          return moment(date).format('LT')
        }
    },
    computed: {
        ...mapGetters({
            'sessions': 'getSessions'
        })
    },
    mounted: function() {
        this.fetchSessions()
    },
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
    "session session";
  grid-template-columns: 100px 1fr;
  grid-template-rows: 10vh 90vh;
}

.session-id-wrapper {
  grid-area: session;
  // align-items: center;
  // justify-content: center;
  // align-content: center;
  // display:grid;
  // grid-template-columns: 5% 1fr 5%;
  overflow-y: scroll;
}

ul {
    list-style: none;
    padding: 0;
    margin: 0;
    width: 100%;
    grid-column: 2/3; 
    li {
        font-size: 14px;
        list-style: none;
        margin: 0;
        border-bottom: 1px solid lightgrey;
        // color:white;
        background:white;
        text-align: left;
        padding: 0 20px;

        .date-time{
          color:lightgray;
          text-transform: uppercase;
          font-size: 12px;
          font-family: 'Courier New', Courier, monospace;
          font-weight: 600;

        }

        .session-title {
          padding: 5px 0;
          font-size: 20px;
          font-weight: bold;
        }

        .session-author {
          font-size: 14px;
          color:gray;
        }

        a {
            display: block;
            padding: 8px 0;
            text-decoration: none;
            color: black;
            transition: transform 0.2s;

            &:hover {
                transition: transform 0.2s;
                transform: translateX(5px)
            }
        }
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

