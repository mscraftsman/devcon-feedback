import Vue from "vue";
import Router from "vue-router";
// import Home from '../views/Home'
import Login from "./views/PageLogin.vue";
// import SessionID from "../views/SessionID";
// import Feedback from "../views/Feedback";

Vue.use(Router);

export default new Router({
  mode: "hash",
  routes: [
    {
      path: "/",
      name: "login",
      component: Login
    }
    // {
    //   path: "/sessionid",
    //   name: "sessionid",
    //   component: SessionID
    // },
    // {
    //   path: "/feedback",
    //   name: "feedback",
    //   component: Feedback
    // }
  ]
});
