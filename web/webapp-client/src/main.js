import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import store from "./store/index";
import "./registerServiceWorker";
// import { Tabs, Tab } from "vue-tabs-component";
import "./assets/scss/main.scss";
import Tab from "vue-swipe-tab";

Vue.use(Tab);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
