import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import store from "./store/index";
import "./registerServiceWorker";
import { Tabs, Tab } from "vue-tabs-component";

Vue.component("tabs", Tabs);
Vue.component("tab", Tab);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
