<template>
  <v-app
          id="inspire"
          :dark="darkmode"
  >

    <AppNavigation />

    <AppToolbar />

    <v-content v-if="authLoaded">
      <v-container class="pa-0">
        <router-view />
      </v-container>
    </v-content>

    <v-dialog v-model="dialog" hide-overlay persistent width="300">
      <v-card>
        <v-card-text>
          Loading...
          <v-progress-linear indeterminate class="mb-0"></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>

    <AppFooter />

  </v-app>
</template>

<script>
  import { mapGetters, mapActions } from "vuex";
  import AppNavigation from "@/gocipe/shared-ui/AppNavigation.vue";
  import AppToolbar from "@/gocipe/shared-ui/AppToolbar.vue";
  import AppFooter from "@/gocipe/shared-ui/AppFooter.vue";
  import authMixins from "@/gocipe/mixins/auth.js";

  export default {
    data: () => ({
      isActive: true
    }),
    computed: {
      ...mapGetters({
        darkmode: "auth/getDarkmode"
      }),
      dialog() {
        return !this.authLoaded;
      }
    },
    methods: {
      loadPage(item) {
        this.currentPage = item;
        console.log(this.currentPage);
      },
      setActive() {
        this.isActive = true;
      },
      setMenuItemActive(url) {
        let currentRoute = this.$route.name;

        if (currentRoute == url.name) {
          return true;
        }
      }
    },
    components: {
      AppNavigation,
      AppToolbar,
      AppFooter
    },
    props: {
      source: String
    },
    mixins: [authMixins]
  };
</script>


<style lang="scss">
  .data_table_size {
    width: 100%;
    height: 100%;
  }
  .nav-container {
  &.active {
     background-color: blue;
   }
  }

  .hidden-date {
  .vdatetime-input {
    display: none;
  }
  }
</style>
