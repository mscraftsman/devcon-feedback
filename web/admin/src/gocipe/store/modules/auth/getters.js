//@ts-check
const getters = {
  getDrawer: state => state.drawer,
  getDarkmode: state => state.darkmode,
  isAuthenticated: state => state.isAuthenticated,
  authStatus: state => () => state.status,
  getToken: state => state.token,
  getUser: state => state.user
};

export default getters;
