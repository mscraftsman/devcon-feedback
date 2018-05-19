import Vue from "vue";
import Vuex from "vuex";
import _ from "lodash";
import axios from "axios";

Vue.use(Vuex);

const state = {
  api: "http://localhost:8080",
  count: 0,
  userid: "abcs",
  votes: [],
  stats: [],
  user: {
    status: false, // true or false
    data: {
      id: "",
      name: "",
      photo: ""
    }
  },
  questions: [
    "How would you rate the speaker?",
    "How would you rate the session content?",
    "Did you learn something new?",
    "Remarks for speaker or organisation?"
  ],
  sessions: [],
  sessionsFlat: [],
  speakers: [],
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
    log: []
  }
};

const mutations = {
  CURRENT_SESSION(state, payload) {
    state.viewingSession = payload;
  },
  INCREMENT(state) {
    state.count++;
  },
  DECREMENT(state) {
    state.count--;
  },
  NEXT_QUESTION(state) {
    // let session = {...state.sessions[activeIndex]}

    // if (session.questionCurrent < 3) {
    //   session.questionCurrent++
    // }

    state.sessionCurrent.questionCurrent++;
  },
  UPDATE_SESSIONS(state, payload) {
    let sessionsFlat = payload.data
      .map(groups => groups.sessions)
      .reduce(function(acc, curr) {
        return [...acc, ...curr];
      }, []);
    state.sessions = payload.data;
    state.sessionsFlat = sessionsFlat;
  },
  UPDATE_SPEAKERS(state, payload) {
    state.speakers = payload.data;
  },
  ERROR_LOG_ADD(state, payload) {
    let error = {
      timestamp: new Date().toISOString(),
      type: payload.type || info,
      text: payload.text
    };
    state.error.log.push(error);
  },
  UPDATE_STATS(state, payload) {
    state.stats = payload;
  },
  USER_LOGIN(state, payload) {
    state.user.status = true;
    state.user.data = payload;
  },
  USER_LOGOUT(state) {
    state.user.status = false;
    state.user.data = {
      name: "",
      id: "",
      photo: ""
    };
  },
  USER_VOTES(state, payload) {
    state.votes = payload.data;
  },
  setUserId(state, payload) {
    state.userid = payload;
  }
};

const getters = {
  getName(state) {
    return state.user.name;
  },
  getUser(state) {
    return state.user;
  },
  getSessionCurrent(state) {
    return state.sessionCurrent;
    // return _.filter(state.sessions, {active: true})[0]
  },
  getQuestionText(state) {
    return state.questions;
  },
  getSessions(state) {
    return state.sessions;
  },
  getSessionsFlat(state) {
    return state.sessions;
  },
  getSessionDetail(state) {
    return;
  },
  getSpeakers(state) {
    return state.speakers;
  },
  getVotes(state) {
    return state.votes;
  },
  getSpeaker(state, id) {},
  getStats(state) {
    return state.stats;
  },
  getErrors(state) {
    return state.error.log;
  }
};

const actions = {
  incrementAsync({ commit }) {
    setTimeout(() => {
      commit("INCREMENT");
    }, 200);
  },
  fetchSessions({ commit, state }) {
    // let url = state.api + "/sessions.json";
    let url = "https://sessionize.com/api/v2/m1l86vhf/view/Sessions";
    axios
      .get(url)
      .then(function(response) {
        // console.log(response);
        commit("UPDATE_SESSIONS", response);
      })
      .catch(function(error) {
        console.log(error);
      });
  },
  fetchSpeakers({ commit, state }) {
    let url = "https://sessionize.com/api/v2/m1l86vhf/view/Speakers";
    axios
      .get(url)
      .then(function(response) {
        commit("UPDATE_SPEAKERS", response);
      })
      .catch(function(error) {
        console.log(error);
      });
  },
  checkUser({ commit, state }) {
    let url = "/b/me";
    axios
      .get(url)
      .then(function(response) {
        if (response.data) {
          commit("USER_LOGIN", response.data.data);
        }
      })
      .catch(function(error) {
        console.log(error);
      });
  },
  triggerLogout({ commit }) {
    commit("USER_LOGOUT");
    // TODO clear jwt
  },
  fetchVotes({ commit }) {
    let url = "b/api/feedbacks/me";
    axios
      .get(url)
      .then(function(response) {
        if (response.status) {
          commit("USER_VOTES", response.data);
        }
      })
      .catch(function(error) {
        //jo bolé so nihal !
      });
  },
  submitVote({ commit, dispatch }, payload) {
    let url = "/b/api/feedbacks";

    axios
      .post(url, payload)
      .then(response => {
        // thanks
        // console.log(response);
        dispatch("fetchVotes");
      })
      .catch(function(error) {
        console.log(error);
      });
  },
  fetchStats({ commit, dispatch }) {
    let url = "/b/api/stats";

    axios
      .get(url)
      .then(response => {
        commit("UPDATE_STATS", response.data.data);
      })
      .catch(function(error) {
        console.log(error);
      });
  }
};

const store = new Vuex.Store({
  state,
  mutations,
  actions,
  getters
});

export default store;
