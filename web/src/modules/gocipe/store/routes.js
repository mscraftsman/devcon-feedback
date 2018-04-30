
// Feedback represents feedback from a visitor on a session
import FeedbackEdit from "../views/FeedbacksEdit.vue";
import FeedbackList from "../views/FeedbacksList.vue";

// Visitor represents a conference visitor
import VisitorEdit from "../views/VisitorsEdit.vue";
import VisitorList from "../views/VisitorsList.vue";

// Session represents a conference session
import SessionEdit from "../views/SessionsEdit.vue";
import SessionList from "../views/SessionsList.vue";


let routes = [
  
  {
    path: "/feedbacks/:id",
    name: "feedbacksEdit",
    props: true,
    icon: "dashboard",
    component: FeedbackEdit,
    entity: "Feedbacks"
  },
  {
    path: "/feedbackslist/",
    name: "feedbacksList",
    icon: "dashboard",
    component: FeedbackList,
    entity: "Feedbacks"
  },
  
  {
    path: "/visitors/:id",
    name: "visitorsEdit",
    props: true,
    icon: "dashboard",
    component: VisitorEdit,
    entity: "Visitors"
  },
  {
    path: "/visitorslist/",
    name: "visitorsList",
    icon: "dashboard",
    component: VisitorList,
    entity: "Visitors"
  },
  
  {
    path: "/sessions/:id",
    name: "sessionsEdit",
    props: true,
    icon: "dashboard",
    component: SessionEdit,
    entity: "Sessions"
  },
  {
    path: "/sessionslist/",
    name: "sessionsList",
    icon: "dashboard",
    component: SessionList,
    entity: "Sessions"
  }
  
];

let entities = [
  
  "Feedbacks",
  
  "Visitors",
  
  "Sessions"
  
];

function registerRoutes(router) {
  router.addRoutes(routes);
}

export default {
  routes,
  entities,
  registerRoutes
}
