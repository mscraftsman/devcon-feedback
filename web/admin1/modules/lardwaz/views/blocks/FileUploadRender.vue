<template>
  <div class="file-upload__container">
    <div
      class="informative__text"
      v-if="content && content.value && content.value.label"
    >{{content.value.label}}</div>
    <div class="informative__text" v-else>File available for download</div>
    <div class="button__wrapper">
      <a
        v-if="
          content && 
          content.value && 
          content.value.buttonText"
        class="download__button"
        :href="content && content.value && content.value.filePath !== null ? content.value.filePath : '#'"
        @click="triggerGA()"
        download
      >{{content.value.buttonText}}</a>
      <a
        :href="content && content.value && content.value.filePath !== null ? content.value.filePath : '#'"
        class="download__button"
        @click="triggerGA()"
        v-else
      >Download</a>
    </div>
  </div>
</template>
<script>
export default {
  props: ["content", "meta"],
  methods: {
    triggerGA() {
      if (
        typeof this.meta.properties !== "undefined" &&
        this.meta.properties.eventCategory
      ) {
        ga(
          "send",
          "event",
          this.meta.properties.eventCategory,
          "download",
          this.content.value.label
        );
      }
    }
  }
};
</script>
<style lang="scss">
</style>
