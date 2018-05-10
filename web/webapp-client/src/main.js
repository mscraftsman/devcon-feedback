import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import store from "./store/index";
import "./registerServiceWorker";
import "./assets/scss/main.scss";

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
