//@ts-check
import { types } from "./types.js";

const mutations = {
  [types.TOGGLE_DRAWER]: function(state) {
    state.drawer = !state.drawer;
  },
  [types.DARKMODE_CHANGE]: function(state, payload) {
    state.darkmode = payload;
  },
  [types.AUTH_LOGIN]: (state, user) => {
    state.user = user;
  },
  [types.AUTH_LOADING]: state => {
    state.status = "loading";
  },
  [types.AUTH_SUCCESS]: state => {
    state.status = "success";
    state.isAuthenticated = true;
  },
  [types.AUTH_ERROR]: state => {
    state.status = "error";
    state.isAuthenticated = false;
  },
  [types.AUTH_LOGOUT]: state => {
    state.status = "logged out";
    state.isAuthenticated = false;
  },
  [types.AUTH_TOKEN_SET]: (state, token) => {
    state.token = token;
  },
  [types.AUTH_TOKEN_DELETE]: state => {
    state.token = null;
  }
};

export default mutations;
