<template>
  <div class="single-component-wrapper">
    <v-textarea
      box
      label="Markdown code"
      rows="15"
      v-model="local"
      @change="update"
      @input="update"
    ></v-textarea>
  </div>
</template>
<script>
export default {
  props: ["value"],
  components: {},
  data() {
    return {
      local: this.value,
      valid: false
    };
  },
  computed: {},
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
<style lang="scss">
</style>
