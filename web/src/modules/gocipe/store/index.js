import actions from "./actions";
import getters from "./getters";
import mutations from "./mutations";
import routes from "./routes";

const namespaced = true;

const state = {
  entities: routes
};

export default {
  namespaced,
  state,
  actions,
  getters,
  mutations
};