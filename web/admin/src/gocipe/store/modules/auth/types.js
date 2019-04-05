const module_name = "auth";
const types = {
  TOGGLE_DRAWER: "TOGGLE_DRAWER", // Toggle the sidebar drawer
  DARKMODE_CHANGE: "DARKMODE_CHANGE", // Trigger the theme change between light and dark
  AUTH_LOGIN: "AUTH_LOGIN", // Check if logged in
  AUTH_CHECK: "AUTH_CHECK", // Check if logged in
  AUTH_LOADING: "AUTH_LOADING",
  AUTH_SUCCESS: "AUTH_SUCCESS",
  AUTH_ERROR: "AUTH_ERROR",
  AUTH_LOGOUT: "AUTH_LOGOUT", // Trigger applicated level logout
  AUTH_TOKEN_SET: "AUTH_TOKEN_SET",
  AUTH_TOKEN_DELETE: "AUTH_TOKEN_DELETE"
};

// Prefix all module names with module name for easy dispatch in the app
let namespaced = Object.keys(types)
  .map(function(curr, index) {
    return [curr, module_name + "/" + types[curr]];
  })
  .reduce(function(prev, curr) {
    prev[curr[0]] = curr[1];
    return prev;
  }, {});

// when importing use as follows :
// import { namespaced as auth } from "@/store/modules/auth/types.js";

export { types, namespaced };
