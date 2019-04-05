import "@babel/polyfill";
import Vue from "vue";
import "./plugins/vuetify";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "./registerServiceWorker";
import VueTimeago from "vue-timeago";
import "@/gocipe/components-registration.js";
import Datetime from "vue-datetime";
import "vue-datetime/dist/vue-datetime.css";

Vue.config.productionTip = false;

Vue.use(VueTimeago, {
  name: "Timeago", // Component name, `Timeago` by default
  locale: "en", // Default locale
  // We use `date-fns` under the hood
  locales: {
    "fr-FR": require("date-fns/locale/fr")
  }
});

Vue.use(Datetime);

Vue.directive("setter", {
  bind(el, binding, vnode, oldVnode) {
    let d = binding.rawName.indexOf(".");
    if (d !== -1) {
      const s = "set" + binding.rawName.substr(d + 1, 1).toUpperCase() + binding.rawName.substr(d + 2);

      const setter = binding.value[s];
      const that = binding.value;

      el.addEventListener("input", () => setter.call(that, el.value));
    }
  }
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
