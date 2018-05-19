import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import store from "./store/index";
import "./registerServiceWorker";
import "./assets/scss/main.scss";
import VueAnalytics from "vue-analytics";
import VueLodash from "vue-lodash";

Vue.use(VueLodash); // options is optional

Vue.use(VueAnalytics, {
  id: "UA-12103827-11",
  router
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
