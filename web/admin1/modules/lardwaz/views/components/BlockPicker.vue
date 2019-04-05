<template>
  <v-menu v-model="menu" :close-on-content-click="false" :nudge-width="200" offset-x>
    <v-btn slot="activator" dark color="primary" fab hover class="block-button">
      <v-icon>add</v-icon>
    </v-btn>

    <v-card width="500px">
      <v-subheader>Add block</v-subheader>
      <v-layout row wrap>
        <v-flex v-for="block in toolbox" v-if="!block.disabled" :key="block.type" xs6>
          <v-list-tile @click="add(block); hideSheet()">
            <v-list-tile-avatar>
              <v-avatar size="32px" tile>
                <v-icon class="ml-2">{{ block.icon }}</v-icon>
              </v-avatar>
            </v-list-tile-avatar>
            <v-list-tile-title>{{ block.name }}</v-list-tile-title>
          </v-list-tile>
        </v-flex>
      </v-layout>
    </v-card>
  </v-menu>
</template>

<script>
import FontAwesomeIcon from "@fortawesome/vue-fontawesome";
import { mapGetters, mapActions } from "vuex";

export default {
  data() {
    return {
      menu: false
    };
  },
  methods: {
    ...mapActions({
      add: "lardwaz/add"
    }),
    hideSheet() {
      this.menu = false;
    }
  },
  computed: {
    ...mapGetters({
      toolbox: "lardwaz/getToolbox"
    }),
    icon: function() {
      return faCoffee;
    }
  },
  components: {
    FontAwesomeIcon
  }
};
</script>

<style lang="scss" scoped>
.add-block {
  color: #ddd;
  font-size: 30px;
  text-shadow: inset 0 0 5px #ddd;
  font-weight: 900;
  padding-top: 20px;
}

.block-button {
  margin: 0 auto;
  position: absolute;
  left: 0;
  right: 0;
  margin-top: 35px;
  /*width: 100%;*/
  /*text-align: right;*/
}

.component-wrapper {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  grid-auto-rows: minmax(30px, 1fr);
  padding: 10px 20px;
  grid-gap: 10px;
  .component-button {
    text-transform: uppercase;
    font-size: 10px;
    // background: #999;
    // display: flex;
    // align-items: center;
    // justify-content: center;
    font-weight: 600;
    transition: background 0.4s;
    color: lightgray;
    padding: 5px;
    // border-radius: 5px;
    &:not([disabled]) {
      border: 0;
      // background: linear-gradient(150deg, #0bceb7 10%, #09b3af 70%, #09b3af 94%);
      color: black;
    }
    &:hover {
      box-shadow: 0 0 5px #999;
      cursor: not-allowed;
    }
    &:hover:not([disabled]) {
      // background: white;
      box-shadow: 0 0 10px #999;
      cursor: pointer;
      transition: background 0.4s;
    }
    .icon {
      width: 100%;
      // height: 50px;
      font-size: 3em;
      display: block;
    }
    .label {
      display: block;
      font-size: 13px;
    }
  }
}
</style>
