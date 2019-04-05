<template>
  <div class="mt-2 mb-2 file__uploader">
    <div class="caption v-label v-label--active theme--light pb-2">{{$attrs.label}}</div>
    <div class="uploader-wrapper">
      <input type="file" @change="onChange" name="fileUpload" ref="fileUpload">

      <template v-if="uploadStatus.uploading">
        <v-progress-linear :indeterminate="uploadStatus.uploading"></v-progress-linear>
      </template>
      <v-alert
        :value="uploadStatus.success"
        type="success"
        dismissible
        class="pa-2"
        transition="scale-transition"
      >file upload successful</v-alert>
      <v-alert
        :value="uploadStatus.error"
        type="error"
        dismissible
        small
        transition="scale-transition"
      >{{ uploadStatus.error | capitalize }}</v-alert>
      <v-alert
        :value="$attrs.hint"
        color="gray lighten-3"
        dark
        icon="info"
        class="pa-2"
      >{{ $attrs.hint }}</v-alert>
    </div>
  </div>
</template>

<script>
import { AdminClient } from "@/services/service_admin_pb_service";
import { UploadRequest, UploadOpts } from "@/services/service_admin_pb";
import { File } from "@/services/service_admin_pb";

function convertFileToBinary(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsArrayBuffer(file);
    reader.onload = () => {
      let array = new Uint8Array(reader.result);
      resolve(array);
    };
    reader.onerror = error => reject(error);
  });
}

let cli = new AdminClient("/api");

export default {
  inheritAttrs: false,
  props: ["field", "rpc", "entityid"],
  data() {
    return {
      file: "",
      uploadStatus: {
        success: false,
        error: null,
        uploading: false
      },
      uploadOptions: null,
      prefillopts: ""
    };
  },
  components: {},
  created() {
    this.uploadOptions = new UploadOpts();
    this.uploadOptions.setWatermark(false);
  },
  filters: {
    capitalize: function(value) {
      if (!value) return "";
      value = value.toString();
      return value.charAt(0).toUpperCase() + value.slice(1);
    }
  },
  computed: {},
  methods: {
    onChange(file) {
      console.log("file upload triggered");
      if (!file) {
        console.log("FileReader API not supported: use the <form>, Luke!");
        return;
      }

      // Resets alerts
      this.uploadStatus.error = null;
      this.uploadStatus.success = false;

      convertFileToBinary(this.$refs.fileUpload.files[0])
        .then(data => {
          console.log(data);

          // Finds file name
          let req = new UploadRequest();
          req.setId(this.entityid);
          req.setField(this.field);
          req.setFilename(this.$refs.fileUpload.files[0].name);

          req.setContent(data);
          req.setUploadoptions(this.uploadOptions);
          this.uploadStatus.uploading = true;

          cli[this.rpc](req, (err, resp) => {
            this.uploadStatus.uploading = false;
            if (err) {
              console.log(err);
              this.uploadStatus.error = err.message;
              return;
            }
            this.uploadStatus.success = true;
            this.file = resp.getUri();
            this.$emit("gocipe", this.file);
          });
        })
        .catch(e => {
          console.log(e);
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.uploader-wrapper {
  background: #c9c9c9;
}
.file-container {
  height: 200px;
  width: 200px;
  img {
    display: block;
    height: 200px;
    width: auto;
  }
}

.file__uploader {
  width: 100%;
}
</style>
