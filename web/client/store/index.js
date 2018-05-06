import Vue from 'vue'
import Vuex from 'vuex'
import _  from 'lodash'
import axios from 'axios'

Vue.use(Vuex)

const state = {
  api: 'http://localhost:3333',
  count: 0,
  userid: "abcs",
  user: {
    authenticated: true, // true or false
    name: "Sandeep Ramgolam",
    jwt: ""
  },
  questions: [
    "How would you rate the speaker?",
    "How would you rate the session content?",
    "Did you learn something new?",
    "Remarks for speaker or organisation?"
  ],
  sessions: [],
  viewingSession: null,
  sessionCurrent: {
      id: 0,
      name: "", // name of session
      active: false, //true or false
      questionCurrent: 1, // 0-3
      reaction1: null, // -2 -1 0 1 2
      reaction2: null,
      reaction3: null,
      reaction4: null
  },
  error: {
    log: [
      
    ]
  }
}

const mutations = {
  CURRENT_SESSION (state, payload) {
    state.viewingSession = payload
  },
  INCREMENT (state) {
    state.count++
  },
  DECREMENT (state) {
    state.count--
  },
  NEXT_QUESTION (state) {

    // let session = {...state.sessions[activeIndex]}

    // if (session.questionCurrent < 3) {
    //   session.questionCurrent++
    // }

    state.sessionCurrent.questionCurrent++
  },
  UPDATE_SESSIONS (state, payload) {
    state.sessions = payload.data
  },
  ERROR_LOG_ADD (state, payload) {
    let error = {
      timestamp : new Date().toISOString(),
      type: payload.type || info,
      text: payload.text
    }
    state.error.log.push(error)
  },
  setUserId (state, payload) {
    state.userid = payload
  }
}

const getters = {
  getName (state) {
    return state.user.name
  },
  getUserId (state) {
    return state.userid
  },
  getSessionCurrent (state) {
    return state.sessionCurrent
    // return _.filter(state.sessions, {active: true})[0]
  },
  getQuestionText (state) {
    return state.questions
  },
  getSessions (state) {
    return state.sessions
  },
  getSessionDetail (state) {
    return 
  },
  getErrors (state) {
    return state.error.log
  }
}

const actions = {
  incrementAsync ({ commit }) {
    setTimeout(() => {
      commit('INCREMENT')
    }, 200)
  },
  fetchSessions ({ commit, state }) {
    let url = state.api + '/sessions.json';
    axios.get(url)
    .then(function (response) {
      // console.log(response);
      commit('UPDATE_SESSIONS', response)
    })
    .catch(function (error) {
      console.log(error);
    });
  },
  checkUser ({ commit, state }) {
    let url = state.api + '/users.json';
    axios.get(url)
    .then(function (response) {
      // console.log(response);
      let user = _.filter(response.data, {id: state.userid})
      if (user.length > 0) { //found someone
          commit('ERROR_LOG_ADD', {
            type: 'success',
            text: 'Succesfully Logged In'
          })
      } else { // didn't find anyone
        commit('ERROR_LOG_ADD', {
          type: 'error',
          text: 'You don\'t exist'
        })
      }
      console.log(user)
      // commit('UPDATE_SESSIONS', response)
    })
    .catch(function (error) {
      console.log(error);
    });
  }
}

const store = new Vuex.Store({
  state,
  mutations,
  actions,
  getters
})

export default store
