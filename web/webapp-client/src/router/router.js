import Vue from "vue";
import Router from "vue-router";
// import Home from '../views/Home'
import Welcome from "../views/PageWelcome.vue";
import Login from "../views/PageLogin.vue";
import Feedback from "../views/PageFeedback.vue";
// import Sessions from "../views/PageSessions.vue";
import Sessions from "../views/SessionsPage.vue";
import Speakers from "../views/SpeakersPage.vue";
import SessionSingle from "../views/SessionSingle.vue";
import SpeakerSingle from "../views/SpeakerSingle.vue";
import LoginError from "../views/LoginError.vue";

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
      path: "/speakers",
      name: "speakers",
      component: Speakers
    },
    {
      path: "/session/:id",
      name: "session",
      props: true,
      component: SessionSingle
    },
    {
      path: "/speaker/:id",
      name: "speaker",
      props: true,
      component: SpeakerSingle
    },
    {
      path: "/feedback/:id",
      props: true,
      name: "feedback",
      component: Feedback
    },
    {
      path: "/loginerror",
      name: "loginerror",
      component: LoginError
    }
  ]
});
