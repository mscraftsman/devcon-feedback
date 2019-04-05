import { mapGetters } from "vuex";
import { namespaced as auth } from "@/gocipe/store/modules/auth/types.js";

export default {
  created() {
    this.$store.watch(this.$store.getters["auth/authStatus"], status => {
      if (status == "logged out") {
        this.$router.push({ name: "Login" });
        return;
      }
    });
  },
  computed: {
    ...mapGetters({
      isAuthenticated: "auth/isAuthenticated",
      authStatus: "auth/authStatus",
      user: "auth/getUser"
    }),
    fullName() {
      return `${this.user.firstName}  ${this.user.lastName}`;
    },
    authLoaded() {
      return this.authStatus() == "success" || this.authStatus() == "error";
    },
    authURL() {
      return process.env.VUE_APP_AUTH_URL;
    }
  },
  methods: {
    login() {
      console.log(auth.AUTH_LOGIN);
      this.$store.dispatch(auth.AUTH_LOGIN);
    },
    logout() {
      this.$store.dispatch(auth.AUTH_LOGOUT);
    }
  }
};
