<template>
  <div class="single-component-wrapper">
    <v-layout row>
      <v-text-field label="Label" v-model="local.label" @change="update"></v-text-field>
    </v-layout>
    <v-layout row>
      <v-text-field label="Button Text" v-model="local.buttonText" @change="update"></v-text-field>
    </v-layout>
    <v-layout row>
      <component
        :is="'EditWidgetFileUpload'"
        :hint="''"
        :value="local.url"
        @gocipe="(e) => this.updateUrl(e)"
        :field="'File'"
        :rpc="'upload'+information.entityName"
        :entityid="information.id"
      ></component>
    </v-layout>
    <v-layout row>
      <v-text-field label="File path" v-model="local.filePath" disabled></v-text-field>
    </v-layout>
  </div>
</template>

<script>
export default {
  components: {},
  props: {
    information: {
      default: null,
      type: Object
    },
    value: Object
  },
  data() {
    const fields = {
      label: null,
      buttonText: null,
      filePath: null
    };
    let local = Object.assign(fields, this.value);
    return {
      local
    };
  },
  methods: {
    update() {
      this.$emit("input", this.local);
    },
    updateFilePath(val) {
      this.filePath = val.target.files[0];
    },
    updateUrl(e) {
      this.local.filePath = e;
      this.$emit("input", this.local);
    }
  },
  watch: {
    value: {
      handler(val) {
        if (val !== null) {
          this.local = val;
        }
      },
      deep: true,
      immediate: true
    }
  }
};
</script>
<style lang="scss" scoped>
</style>
