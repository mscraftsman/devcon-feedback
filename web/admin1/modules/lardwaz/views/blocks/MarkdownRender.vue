<template>
  <div v-if="content && content.value" v-html="convertToMarkdown(content.value)"></div>
</template>
<script>
let marked = require("marked");

export default {
  data() {
    return {};
  },
  props: ["content"],
  methods: {
    convertToMarkdown(val) {
      if (val) {
        return marked(val, { sanitize: true });
      }
    }
  }
};
</script>
<style lang="scss">
</style>
