<script>
export default {
  props: {
    label: {
      required: true
    },
    hash: {
      required: true
    },
    fontsize: {
      type: String
    },
    tabheight: {
      type: String
    },
    frontIcon: {
      type: String,
      default: ""
    },
    postIcon: {
      type: String,
      default: ""
    },
    RedDot: {
      type: String,
      default: ""
    },
    color: {
      type: String
    }
  },
  data() {
    return {
      isActive: false
    };
  },
  computed: {
    tabHeader() {
      return this.frontIcon + this.label + this.postIcon + this.RedDot;
    },
    tabHash() {
      return `${this.hash.toLowerCase()}`;
    },
    fontSize() {
      return parseInt(this.fontsize) / 75 + "rem";
    },
    tabHeight() {
      return parseInt(this.tabheight) / 75 + "rem";
    }
  }
};
</script>

<template>
    <div class="tabs-panel-content">
        <slot />
    </div>
</template>

<style lang="scss">
.tabs-panel-content {
  background: #333;
  padding: 0 10px;
  .session-panes {
    font-size: 14px;
    list-style: none;
    margin: 0;
    border-bottom: 1px solid lightgrey;
    text-align: left;

    a {
      display: block;
      padding: 8px 0;
      text-decoration: none;
      color: white;
      transition: transform 0.2s;

      &:hover {
        transition: transform 0.2s;
        transform: translateX(5px);
      }
    }
    text-align: left;
    .date-time {
      color: gray;
      text-transform: uppercase;
      font-size: 12px;
      font-family: "Courier New", Courier, monospace;
      font-weight: 600;
    }

    .session-title {
      padding: 5px 0;
      font-size: 20px;
      font-weight: bold;
    }

    .session-author {
      font-size: 14px;
      color: gray;
    }
  }
}
</style>
