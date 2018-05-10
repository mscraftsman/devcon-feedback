import Vue from "vue";
import Router from "vue-router";
// import Home from '../views/Home'
import Welcome from "../views/PageWelcome.vue";
import Login from "../views/PageLogin.vue";
import Feedback from "../views/PageFeedback.vue";
// import Sessions from "../views/PageSessions.vue";
import Sessions from "../views/SessionsPage.vue";
import SessionSingle from "../views/SessionSingle.vue";

Vue.use(Router);

export default new Router({
  mode: "hash",
  routes: [
    {
      path: "/",
      name: "welcome",
      component: Welcome
    },
    {
      path: "/login",
      name: "login",
      component: Login
    },
    {
      path: "/sessions",
      name: "sessions",
      component: Sessions
    },
    {
      path: "/session/:id",
      name: "session",
      props: true,
      component: SessionSingle
    },
    {
      path: "/feedback/:id",
      props: true,
      name: "feedback",
      component: Feedback
    }
  ]
});
