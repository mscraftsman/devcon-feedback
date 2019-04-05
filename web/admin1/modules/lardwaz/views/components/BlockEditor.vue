<template>
  <div class="block-editor-wrapper">
    <div class="components-editor">
      <v-expansion-panel focusable class="expansion-panel" v-model="panels">
        <draggable
          class="full-width"
          :list="blocks"
          :options="{handle:'.handle'}"
          @start="log"
          @end="log"
        >
          <v-expansion-panel-content
            text-md-center
            v-for="(block, index) in blocks"
            :key="index.id"
          >
            <div slot="header">
              <v-icon class="handle">drag_indicator</v-icon>
              <strong>{{ block.name }}</strong>
              -
              <span>
                <template v-if="block.value !== null">
                  <template
                    v-if="typeof pageComponents[block.type].properties.labelSelector === 'undefined'"
                  >{{trimContent(block.value)}}</template>

                  <template
                    v-else-if="
                  typeof pageComponents[block.type]  !== 'undefined' &&
                  typeof pageComponents[block.type].properties.labelSelector !== 'undefined' && 
                  block.value[(pageComponents[block.type].properties.labelSelector)] !==  ''"
                  >{{trimContent(block.value[pageComponents[block.type].properties.labelSelector])}}</template>
                </template>
              </span>
            </div>
            <v-card>
              <v-card-text>
                <component
                  :is="block.type + 'Block'"
                  v-model="block.value"
                  :information="information"
                  v-bind="$attrs"
                ></component>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn
                    small
                    color="error"
                    class="white--text"
                    @click="deleteItem(index)"
                    outline
                  >Delete</v-btn>
                </v-card-actions>
              </v-card-text>
            </v-card>
          </v-expansion-panel-content>
        </draggable>
      </v-expansion-panel>

      <BlockPicker/>

      <v-dialog v-model="deleteDialog" persistent max-width="290">
        <v-card>
          <v-card-title class="headline">Are you sure to delete this block?</v-card-title>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn light @click.native="deleteDialog = false">No</v-btn>
            <v-btn color="error" @click.native="deleteItemBlock()">Yes</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import FontAwesomeIcon from "@fortawesome/vue-fontawesome";
import BlockPicker from "./BlockPicker";
import draggable from "vuedraggable";

// Components
import PageComponents from "@lardwaz-config/page-components.js";

import TextBlock from "../blocks/TextBlock";
import TextareaBlock from "../blocks/TextareaBlock";
import HeadingBlock from "../blocks/HeadingBlock";
import ImageBlock from "../blocks/ImageBlock";
import QuoteBlock from "../blocks/QuoteBlock";
import HTMLBlock from "../blocks/HTMLBlock";
import YoutubeBlock from "../blocks/YoutubeBlock";
import GalleryBlock from "../blocks/GalleryBlock";
import RelatedBlock from "../blocks/RelatedBlock";
import IndicatorBlock from "../blocks/IndicatorBlock";
import TranscriptBlock from "../blocks/TranscriptBlock";
import FooterBlock from "../blocks/FooterBlock";
import MarkdownBlock from "../blocks/MarkdownBlock";
import LegacyBlock from "../blocks/LegacyBlock";
import FileUploadBlock from "../blocks/FileUploadBlock";
import PDFViewerBlock from "../blocks/PDFViewerBlock";
import HTMLEmbedBlock from "../blocks/HTMLEmbedBlock";

/** Mailefficient **/
// import MailHeaderBlock from "../blocks/mail/MailHeaderBlock";
// import MailJumbotronBlock from "../blocks/mail/MailJumbotronBlock";
// import MailColumnBlock from "../blocks/mail/MailColumnBlock";
// import MailTextHeadingBlock from "../blocks/mail/MailTextHeadingBlock";
// import MailImageTextBlock from "../blocks/mail/MailImageTextBlock";
// import MailArticlesListingBlock from "../blocks/mail/MailArticlesListingBlock";
// import MailAdBlock from "../blocks/mail/MailAdBlock";

export default {
  props: {
    information: {
      default: null,
      type: Object
    }
  },
  data() {
    return {
      deleteDialog: false,
      blockID: null,
      blocksOnPage: null,
      pageComponents: {},
      panels: null
    };
  },
  computed: {
    blocks: {
      get() {
        return this.$store.getters["lardwaz/getBlocks"];
      },
      set(value) {
        this.$store.commit("lardwaz/dragAndDrop", value);
      }
    }
  },
  components: {
    draggable,
    BlockPicker,
    TextBlock,
    TextareaBlock,
    HeadingBlock,
    ImageBlock,
    QuoteBlock,
    FontAwesomeIcon,
    HTMLBlock,
    YoutubeBlock,
    GalleryBlock,
    RelatedBlock,
    FooterBlock,
    MarkdownBlock,
    LegacyBlock,
    IndicatorBlock,
    TranscriptBlock,
    FileUploadBlock,
    PDFViewerBlock,
    HTMLEmbedBlock
    // MailHeaderBlock,
    // MailJumbotronBlock,
    // MailColumnBlock,
    // MailTextHeadingBlock,
    // MailImageTextBlock,
    // MailArticlesListingBlock,
    // MailAdBlock
  },
  beforeMount() {
    PageComponents.map(block => {
      this.pageComponents[block.type] = block;
    });
  },
  methods: {
    deleteItem(item) {
      this.deleteDialog = true;
      this.blockID = item;
    },
    deleteItemBlock() {
      this.blocks.splice(this.blockID, 1);
      this.deleteDialog = false;
    },
    trimContent(val) {
      if (val && val.length) {
        let strippedContent = val.replace(/(<([^>]+)>)/gi, "");
        return strippedContent.substring(0, 30) + "...";
      } else {
        return val;
      }
    },
    log(e) {
      //console.log(e);
    }
  },
  watch: {
    blocks: {
      handler(val) {
        setTimeout(() => {
          window.dispatchEvent(new Event("resize"));
        }, 200);
      },
      deep: true,
      immediate: true
    },
    panels(val) {
      setTimeout(() => {
        window.dispatchEvent(new Event("resize"));
      }, 200);
    }
  }
};
</script>

<style lang="scss" scoped>
.expansion-panel {
  li {
    margin: 0;
    //max-width: 100%;
  }
}

.handle {
  position: relative;
  left: -10px;
  background-color: lightgray;
  border-radius: 2px;
}

.block-editor-wrapper {
}

.components-editor {
  padding: 10px;
  position: relative;

  .block-list {
    text-align: left;
    // padding-bottom: 20px;
    &:focus-within {
      .block-header {
        background: #bbb;
      }
    }
  }
  .full-width {
    width: 100%;
  }
  .block-header {
    --height: 30px;
    display: grid;
    grid-template-columns: 1fr 50px;
    grid-auto-rows: var(--height);
    background: #ddd;
    .block-name {
      font-size: 12px;
      text-transform: uppercase;
      padding-left: 20px;
      color: black;
      line-height: var(--height);
    }
    .block-handle {
      font-size: 20px;
      color: white;
      align-self: center;
      justify-self: center;
      cursor: -webkit-grab;
      // width: 100%;
      // height: 100%;
      // line-height: 100%;
      // height: 50px;
    }
  }
}

.sortable-chosen {
  outline-offset: 0px;
  outline: 2px dashed black;
  .block-handle {
    cursor: -webkit-grabbing;
  }
}

.sortable-ghost {
  cursor: -webkit-grab;
}

.btn-wrapper {
  padding: 20px;
}
</style>
