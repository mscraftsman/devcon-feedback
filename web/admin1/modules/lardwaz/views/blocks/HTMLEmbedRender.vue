<template>
  <div class="embed__container">
    <div class="embed__wrapper" v-if="embed !== null">
      <VueScriptComponent v-if="embed" :script="embed"/>
    </div>
  </div>
</template>
<script>
import VueScriptComponent from "vue-script-component";

export default {
  props: ["content"],
  components: {
    VueScriptComponent
  },
  data() {
    return {
      embed: null
    };
  },
  mounted() {
    if (this.content && this.content.value && this.content.value.embed) {
      this.embed = this.content.value.embed;
    }
  },
  watch: {
    content: {
      handler(val) {
        if (this.content && this.content.value && this.content.value.embed) {
          this.embed = "";
          setTimeout(() => {
            this.embed = this.content.value.embed;
          }, 100);
        }
      },
      deep: true,
      immediate: true
    }
  }
};
</script>
<style lang="scss">
.embed__container {
  width: 100%;

  blockquote {
    margin: 0 auto !important;
  }

  .embed__wrapper {
    margin: 0 auto;
    div {
      margin: 0 auto !important;
    }
  }

  iframe {
    div,
    & {
      text-align: center;
      margin: 0 auto !important;
    }
  }

  div {
    text-align: center;

    iframe {
      margin: 0 auto !important;
    }
  }

  twitter-widget {
    margin-left: auto;
    margin-right: auto;
    margin-top: 30px !important;
    margin-bottom: 30px !important;
  }
}
</style>
