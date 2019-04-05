//@ts-check
import { types } from "./types.js";
import { WebApiClient } from "@/services/passport/service_webapi_pb_service";
import {
  IsLoggedRequest,
  LogoutRequest
} from "@/services/passport/service_webapi_pb";

function cookieGet(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2)
    return parts
      .pop()
      .split(";")
      .shift();
}

function cookieDelete(name) {
  document.cookie = name + "=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
}

const actions = {
  [types.TOGGLE_DRAWER]: ({ commit }) => {
    commit("TOGGLE_DRAWER");
  },
  [types.DARKMODE_CHANGE]: ({ commit }, payload) => {
    commit("DARKMODE_CHANGE", payload);
  },
  [types.AUTH_LOGOUT]: ({ state, commit }) => {
    let cli = new WebApiClient(process.env.VUE_APP_AUTH_GRPC + "/wapi"); // send request to ID
    let req = new LogoutRequest();

    commit(types.AUTH_LOADING);

    req.setKey(state.token);
    cli.logout(req, (err, resp) => {
      if (err) {
        // FAILED
        console.log(err);
        commit(types.AUTH_ERROR);
        return;
      }

      cookieDelete("transport");
      commit("AUTH_LOGOUT");
    });
  },
  [types.AUTH_CHECK]: function({ state, commit, dispatch }) {
    dispatch(types.AUTH_TOKEN_SET);

    commit(types.AUTH_LOADING);

    let token = state.token;

    if (token === undefined || token == "") {
      commit(types.AUTH_ERROR);
      console.log("here");
      return;
    }

    let cli = new WebApiClient(process.env.VUE_APP_AUTH_GRPC + "/wapi"); // send request to ID
    let req = new IsLoggedRequest();

    req.setKey(token);
    cli.isLogged(req, (err, resp) => {
      if (err) {
        // FAILED
        console.log(err);
        commit(types.AUTH_ERROR);
        return;
      }

      let user = resp.getUser();

      user = {
        id: user.getId(),
        firstName: user.getFirstname(),
        lastName: user.getLastname(),
        locale: user.getLocale(),
        username: user.getUsername(),
        permissions: user.getPermissionsList(),
        roles: user.getRolesList()
      };

      console.log(user);

      commit(types.AUTH_LOGIN, user);

      // PASSED
      commit(types.AUTH_SUCCESS, token);
    });
  },
  [types.AUTH_TOKEN_SET]: function({ state, commit }) {
    let token = cookieGet("transport");
    commit(types.AUTH_TOKEN_SET, token);
  },
  [types.AUTH_TOKEN_DELETE]: function({ commit }) {
    commit(types.AUTH_TOKEN_DELETE);
  }
};

export default actions;
