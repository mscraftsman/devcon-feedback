<template>
  <div class="file-upload__container">
    <div
      class="informative__text"
      v-if="content && content.value && content.value.label"
    >{{content.value.label}}</div>
    <div class="informative__text" v-else>PDF available for viewing</div>
    <div class="button__wrapper">
      <a
        v-if="
          content && 
          content.value && 
          content.value.buttonText"
        @click="triggerGA()"
        target="_blank"
        class="download__button"
        :href="content && content.value && content.value.filePath !== null ? meta.properties.viewerPath + this.baseURL + encodeURI(content.value.filePath) : '#'"
      >{{content.value.buttonText}}</a>
      <a
        @click="triggerGA()"
        :href="content && content.value && content.value.filePath !== null ? meta.properties.viewerPath + this.baseURL + encodeURI(content.value.filePath) : '#'"
        class="download__button"
        target="_blank"
        v-else
      >View PDF</a>
    </div>
  </div>
</template>
<script>
export default {
  props: ["content", "meta"],
  data() {
    return {
      baseURL: null
    };
  },
  beforeMount() {
    this.baseURL =
      window.location.protocol +
      "//" +
      window.location.hostname +
      (window.location.port.trim().length > 0
        ? ":" + window.location.port
        : "");
  },
  methods: {
    triggerGA() {
      if (
        typeof this.meta !== "undefined" &&
        this.meta.properties.eventCategory
      ) {
        ga(
          "send",
          "event",
          this.meta.properties.eventCategory,
          "view_pdf",
          this.content.value.label
        );
      }
    }
  }
};
</script>
<style lang="scss">
</style>
