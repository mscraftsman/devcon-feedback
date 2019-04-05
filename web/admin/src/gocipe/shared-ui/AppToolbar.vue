<template>
  <v-toolbar :clipped-left="$vuetify.breakpoint.lgAndUp" dense dark app>
    <v-toolbar-title style="width: 300px">
      <v-toolbar-side-icon @click.stop="toggleDrawer" v-if="isAuthenticated"></v-toolbar-side-icon>
      <span class="hidden-sm-and-down">Backoffice LSL Digital</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-menu v-if="isAuthenticated" offset-y>
      <v-btn slot="activator" flat dark>Logged in as {{ fullName }}</v-btn>
      <v-list>
        <v-list-tile>
          <v-list-tile-title>Roles</v-list-tile-title>
        </v-list-tile>
        <v-list-tile v-for="role in user.roles" :key="role">
          <v-list-tile-title>{{role}}</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-menu>
    <v-btn v-if="isAuthenticated" icon @click="logoutDialog = true">
      <v-icon>exit_to_app</v-icon>
    </v-btn>
    <v-btn v-else :href="authURL">Login</v-btn>
    <v-dialog v-model="logoutDialog" width="500" :persistent="true">
      <v-card>
        <v-card-title class="headline grey lighten-2" primary-title>
          Logout
        </v-card-title>
        <v-card-text class="">
          Are you sure you want to logout?
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-btn color="primary" flat @click="logoutDialog = false">
            Cancel
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="logoutClose">
            Proceed
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-toolbar>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import { namespaced as auth } from "@/gocipe/store/modules/auth/types.js";
import authMixins from "@/gocipe/mixins/auth.js";

export default {
  data() {
    return {
      logoutDialog: false
    };
  },
  methods: {
    toggleDrawer() {
      this.$store.dispatch(auth.TOGGLE_DRAWER);
    },
    logoutClose() {
      this.logout();
      this.logoutDialog = false;
    }
  },
  mixins: [authMixins]
};
</script>
