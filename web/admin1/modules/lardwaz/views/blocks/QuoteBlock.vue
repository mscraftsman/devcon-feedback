<template>
  <div class="single-component-wrapper">
    <v-form v-model="valid">
      <v-textarea v-model="local.quote" @keyup="update()" :rules="quoteRules" required></v-textarea>
      <v-text-field label="Author" v-model="local.author" @keyup="update()"></v-text-field>
    </v-form>
  </div>
</template>

<script>
export default {
  inheritAttrs: false,
  props: ["value"],
  data() {
    const fields = {
      quote: "",
      author: ""
    };
    let local = Object.assign(fields, this.value);
    return {
      local,
      valid: false,
      quoteRules: [
        v => !!v || "Quote is required",
        v => v.length >= 1 || "Quote must be greater than 1 character(s)"
      ]
    };
  },
  methods: {
    update() {
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

<style lang="scss" scoped>
.single-component-wrapper {
  textarea {
    width: 100%;
    display: block;
  }
}
</style>
