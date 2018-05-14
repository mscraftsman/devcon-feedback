import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import store from "./store/index";
import "./registerServiceWorker";
import "./assets/scss/main.scss";
import VueAnalytics from "vue-analytics";

Vue.use(VueAnalytics, {
  id: "UA-12103827-11"
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
