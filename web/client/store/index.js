import Vue from 'vue'
import Vuex from 'vuex'
import _  from 'lodash'

Vue.use(Vuex)

const state = {
  count: 0,
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
  sessionCurrent: {
      id: 0,
      name: "", // name of session
      active: false, //true or false
      questionCurrent: 1, // 0-3
      reaction1: null, // -2 -1 0 1 2
      reaction2: null,
      reaction3: null,
      reaction4: null
  }
}

const mutations = {
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
  }
}

const getters = {
  getName (state) {
    return state.user.name
  },
  getSessionCurrent (state) {
    return state.sessionCurrent
    // return _.filter(state.sessions, {active: true})[0]
  },
  getQuestionText (state) {
    return state.questions
  }
}

const actions = {
  incrementAsync ({ commit }) {
    setTimeout(() => {
      commit('INCREMENT')
    }, 200)
  }
}

const store = new Vuex.Store({
  state,
  mutations,
  actions,
  getters
})

export default store
