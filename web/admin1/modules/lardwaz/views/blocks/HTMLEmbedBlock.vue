<template>
  <div class="single-component-wrapper">
    <v-layout row>
      <v-textarea
        box
        label="HTML embed code"
        rows="15"
        v-model="local.embed"
        @change="update"
        @input="update"
      ></v-textarea>
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
      embed: null
    };
    let local = Object.assign(fields, this.value);
    return {
      local
    };
  },
  methods: {
    update() {
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
