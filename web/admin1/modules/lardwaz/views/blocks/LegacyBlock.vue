<template>
    <div class="single-component-wrapper">
        <v-form v-model="valid">
            <v-textarea label="Legacy content" box v-model="local" @keyup="update()" required rows="15"></v-textarea>
        </v-form>
    </div>
</template>
<script>
export default {
  props: ["value"],
  data() {
    return {
      local: this.value,
      valid: false
    };
  },
  methods: {
    update() {
      // emit changes only if validations are passed
      if (this.valid === true) {
        this.$emit("input", this.local);
      }
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
