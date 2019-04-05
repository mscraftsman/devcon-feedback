<template>
    <div>
        <div class="listing--container">
            <!-- Normal Toolbar -->
            <v-toolbar class="transparent listing-toolbar blue-grey darken-3 px-3" dark flat ml-0 v-if="!nested">
                <v-btn icon :to="{name: 'visitors_list'}">
                    <v-icon>arrow_back</v-icon>
                </v-btn>
                <v-toolbar-title>Edit Visitor </v-toolbar-title>

                <v-spacer></v-spacer>
                
                <v-btn color="primary"  @click="update">
                    Save
                    <v-icon right>save</v-icon>
                </v-btn>
            </v-toolbar>
            <!-- Nested Form Toolbar -->
            <v-toolbar class="primary listing-toolbar px-3" dense dark v-else>
                <v-toolbar-title>
                    <h3>Edit Visitor </h3>
                </v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn flat @click.stop="emitCloseForm">
                    <v-icon dark>close</v-icon> Close
                </v-btn>
            </v-toolbar>
        </div>

        <div v-if="loading" centered>
            <v-progress-linear :indeterminate="true"></v-progress-linear>
            <p class="text-xs-center">contacting server...</p>
        </div>

        <v-card v-else flat class="push-content">
            <v-card-text>
                <v-form>
                    <div class="gocipe-form-grid" v-if="entity"> 
                            <div class="gocipe-field-container">
                                <component  :is="'EditWidgetTextfield'"
                                            :label="'Name'"
                                            :hint="''" 
                                            :value="this.entity.getName()" 
                                            @gocipe="(e) => this.entity.setName(e)"></component>
                            </div> 
                            <div class="gocipe-field-container">
                                <component  :is="'EditWidgetTextfield'"
                                            :label="'Bookmark'"
                                            :hint="''" 
                                            :value="this.entity.getBookmark()" 
                                            @gocipe="(e) => this.entity.setBookmark(e)"></component>
                            </div> 
                            <div class="gocipe-field-container">
                                <component  :is="'EditWidgetTextfield'"
                                            :label="'Photo Link'"
                                            :hint="''" 
                                            :value="this.entity.getPhotolink()" 
                                            @gocipe="(e) => this.entity.setPhotolink(e)"></component>
                            </div>
                        <template v-if="componentLoaded">
                            <div class="data-table-responsive-wrapper">
                                <component
                                    :is="'FeedbacksList'"
                                    :key="'Feedbacks_key'"
                                    :filter_id="id" 
                                    :filter_field="'visitor_id'"
                                    :rpc="'setVisitorid'"
                                    :nested="true" 
                                />
                            </div>
                        </template>

                    <component  :is="'EditWidgetStatus'"
                        :label="'Status'"
                        :value="this.entity.getStatus()" 
                        @gocipe="(e) => this.entity.setStatus(e)" >
                    </component>

                    </div>
                </v-form>
            </v-card-text>
        </v-card>
        
        
        
         
        <v-toolbar class="transparent listing-toolbar blue-grey darken-4 px-3" dark flat ml-0  v-if="!nested">
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="update">
                Save
                <v-icon right>save</v-icon>
            </v-btn>
        </v-toolbar>
        <v-toolbar class="transparent listing-toolbar blue-grey darken-4 px-3 sticky" dark flat ml-0  v-else>
            <v-btn color="gray" @click="emitCloseForm">
                Cancel
                <v-icon right>close</v-icon>
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="emitSaveAndCloseForm">
                Save &amp; Close
                <v-icon right>assignment_return</v-icon>
            </v-btn>
        </v-toolbar>

        <v-snackbar v-model="snackbar.show" :bottom="true" :right="true" auto-height :color="snackbar.color" :timeout="6000">
            {{ snackbar.text }}
            <v-btn dark flat v-if="snackbar.color !== 'error'" :to="{name: 'visitors_list', params : { track : id}}">
                <v-icon>arrow_back</v-icon>
            </v-btn>
            <v-btn dark flat @click="snackbarHide">
                <v-icon>close</v-icon>
            </v-btn>
        </v-snackbar>
    </div>
</template>


<script>
import { AdminClient } from "@/services/service_admin_pb_service";
import { GetRequest } from "@/services/service_admin_pb";
import { UpdateVisitorRequest } from "@/services/service_admin_pb";
import { mapGetters } from 'vuex'

let cli = new AdminClient("/api");
export default {
  data() {
    return {
      messages: [],   
      snackbar: {
          show: false, 
          text: ''
      },
      entity: null,
      loading: false,
      id: null,
      componentLoaded: false,
      
    };
  },
  
  props: ["filter_id", "nested", "visible"],
  mounted() {
    this.request()  
  },
  computed: {
    ...mapGetters({
      'token': 'auth/getToken'
    })
  },
  methods: {
        
        log: function(e) {
            console.log(e)
        },
        request: function() {
            if (this.filter_id) {
                this.id = this.filter_id
            } else {
                this.id = this.$route.params.id;
            }

            if (this.nested && typeof this.id === 'undefined') {
                return;
            }

            let req = new GetRequest();
            req.setId(this.id);
            req.setKey(this.token);

            this.loading = true;

            cli.getVisitor(req, (err, resp) => {
                if (err) {
                    console.log(err);
                    this.snackbarShow(err.message, "error");
                    return;
                }
                this.loading = false;
                this.entity = resp.getVisitor();
                this.componentLoaded = true;

                
            });
        },
        update: function() {
    
            let req = new UpdateVisitorRequest();
            req.setVisitor(this.entity);
            req.setKey(this.token);
            // this.loading = true;

            /* this.debug(); */

            cli.updateVisitor(req, (err, resp) => {
                if (err) {
                    console.log(err);
                    this.snackbarShow(err.message, "error");
                    return;
                }

                this.snackbarShow('Visitor Saved');
                
                // this.loading = false;
                if (this.nested) {
                    this.$emit("closeform", true);
                }
            });
        },
        emitSaveAndCloseForm: function() {
            /* console.log("emitting close form to parent"); */
            this.update();
        },
        emitCloseForm: function() {
            /* console.log("emitting close form to parent"); */
            this.$emit("closeform", true);
        },
        debug: function() {
            console.log("ID", this.entity.getId());
            console.log("Status", this.entity.getStatus());
            console.log("UserID", this.entity.getUserid());
            console.log("CreatedAt", this.entity.getCreatedat());
            console.log("Name", this.entity.getName());
            console.log("Bookmark", this.entity.getBookmark());
            console.log("PhotoLink", this.entity.getPhotolink());
        },
        snackbarShow: function(text, color) {
            this.snackbar.show = true;
            this.snackbar.color = color || 'info';
            this.snackbar.text = text || 'something happened';
        },
        snackbarHide: function() {
            this.snackbar.show = false
        },
        
  },
  watch: {
    tab(val) {
      setTimeout(() => {
        window.dispatchEvent(new Event("resize"));
      }, 200);
    }
  }
  /* watch: {
    visible: function(val) {
      if (val) {
        console.log("watching");
        this.request();
      }
    }
  } */
};
</script>

<style lang="scss" scoped>
.gocipe-form-grid {
  width: 100%;
  //   max-width: 800px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(650px, 1fr));
  grid-column-gap: var(--gutter, 60px);
  .gocipe-field-container {
    // grid-column: 1/2;
  }
}

.data-table-responsive-wrapper {
  max-width: 100%;
  overflow-y: auto;
  /* padding: 5px; */
}

.push-content {
  padding-bottom: 100px;
}

.sticky {
  position: fixed;
  bottom: 0;
}
</style>
