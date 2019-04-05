import actions from "./actions.js";
import mutations from "./mutations.js";
import getters from "./getters.js";

const state = {
  drawer: true,
  darkmode: false,
  token: null,
  status: "",
  isAuthenticated: false,
  user: {
    id: "",
    firstName: "",
    lastName: "",
    locale: "",
    username: "",
    permissions: [],
    roles: []
  }
};

const namespaced = true;

export default {
  namespaced,
  state,
  actions,
  getters,
  mutations
};
