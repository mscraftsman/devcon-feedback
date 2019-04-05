<template>
  <div v-if="blocks">
    <template v-for="(component, index) in blocks">
      <component
        :is="component.type + 'Render'"
        :content="component"
        :meta="pageComponents[component.type]"
        :key="index"
      />
    </template>
  </div>
</template>

<script>
import HeadingRender from "@lardwaz/views/blocks/HeadingRender";
import TextRender from "@lardwaz/views/blocks/TextRender";
import TextareaRender from "@lardwaz/views/blocks/TextareaRender";
import QuoteRender from "@lardwaz/views/blocks/QuoteRender";
import ImageRender from "@lardwaz/views/blocks/ImageRender";
import HTMLRender from "@lardwaz/views/blocks/HTMLRender";
import YoutubeRender from "@lardwaz/views/blocks/YoutubeRender";
import GalleryRender from "@lardwaz/views/blocks/GalleryRender";
import RelatedRender from "@lardwaz/views/blocks/RelatedRender";
import FooterRender from "@lardwaz/views/blocks/FooterRender";
import TranscriptRender from "@lardwaz/views/blocks/TranscriptRender";
import IndicatorRender from "@lardwaz/views/blocks/IndicatorRender";
import MarkdownRender from "@lardwaz/views/blocks/MarkdownRender";
import LegacyRender from "@lardwaz/views/blocks/LegacyRender";
import FileUploadRender from "@lardwaz/views/blocks/FileUploadRender";
import PDFViewerRender from "@lardwaz/views/blocks/PDFViewerRender";
import HTMLEmbedRender from "@lardwaz/views/blocks/HTMLEmbedRender";

// Components
import PageComponents from "./page-components.js";

export default {
  name: "Renderer",
  props: ["blocks"],
  data() {
    return {
      pageComponents: {}
    };
  },
  components: {
    HeadingRender,
    TextRender,
    QuoteRender,
    ImageRender,
    TextareaRender,
    HTMLRender,
    YoutubeRender,
    GalleryRender,
    RelatedRender,
    FooterRender,
    MarkdownRender,
    LegacyRender,
    IndicatorRender,
    TranscriptRender,
    FileUploadRender,
    PDFViewerRender,
    HTMLEmbedRender
  },
  beforeMount() {
    PageComponents.map(block => {
      this.pageComponents[block.type] = block;
    });
  }
};
</script>
