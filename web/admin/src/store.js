import Vue from "vue";
import Vuex from "vuex";
import auth from "@/gocipe/store/modules/auth/index";
import lardwaz from "@/../modules/lardwaz/store/index";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
    lardwaz
  }
});
