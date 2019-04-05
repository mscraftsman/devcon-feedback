
// Feedback A Feedback
import FeedbackList from "./forms/FeedbacksList.vue";
import FeedbackEdit from "./forms/FeedbacksEdit.vue";

// Visitor A visitor
import VisitorList from "./forms/VisitorsList.vue";
import VisitorEdit from "./forms/VisitorsEdit.vue";

// Rating A rating
import RatingList from "./forms/RatingsList.vue";
import RatingEdit from "./forms/RatingsEdit.vue";


let routes = [
  {
    path: "feedbacks",
    name: "feedbacks_list",
    component: FeedbackList,
    meta: {
      icon: "work",
      entity: "Feedbacks",
      showInMenu: true
    }
  },
  {
    path: "feedbacks/edit/:id",
    name: "feedbacks_edit",
    component: FeedbackEdit,
    meta: {
      icon: "work",
      entity: "Feedbacks",
      showInMenu: false
    }
  },
  {
    path: "visitors",
    name: "visitors_list",
    component: VisitorList,
    meta: {
      icon: "work",
      entity: "Visitors",
      showInMenu: true
    }
  },
  {
    path: "visitors/edit/:id",
    name: "visitors_edit",
    component: VisitorEdit,
    meta: {
      icon: "work",
      entity: "Visitors",
      showInMenu: false
    }
  },
  {
    path: "ratings",
    name: "ratings_list",
    component: RatingList,
    meta: {
      icon: "work",
      entity: "Ratings",
      showInMenu: true
    }
  },
  {
    path: "ratings/edit/:id",
    name: "ratings_edit",
    component: RatingEdit,
    meta: {
      icon: "work",
      entity: "Ratings",
      showInMenu: false
    }
  }
];

export default routes;
