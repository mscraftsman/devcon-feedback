<template>
  <div :class="{ 'px-3' : !nested }" class="listing--container">
    <v-toolbar color="transparent" class="listing-toolbar blue-grey darken-3 px-3" dark flat ml-0 v-if="!nested">
      <v-btn icon>
        <v-icon>work</v-icon>
      </v-btn>
      <v-toolbar-title class="ml-0 pl-0">Rating </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-text-field class="mb-3" prepend-inner-icon="search" label="Search" single-line hide-details small v-model="search"></v-text-field>
      <v-btn-toggle flat class="transparent" v-model="toggle_status">
        <v-btn icon small :color="toggle_status === 0 ? 'success' : ''" flat>
          <v-icon small dark>check_circle</v-icon>
        </v-btn>
        <v-btn icon small :color="toggle_status === 1 ? 'info' : ''" flat>
          <v-icon small dark>bookmark</v-icon>
        </v-btn>
        <v-btn icon small :color="toggle_status === 2 ? 'warning' : ''" flat>
          <v-icon small dark>rowing</v-icon>
        </v-btn>
        <v-btn icon small :color="toggle_status === 3 ? 'error' : ''" flat>
          <v-icon small dark>visibility_off</v-icon>
        </v-btn>
      </v-btn-toggle>
      <v-spacer></v-spacer>
      <v-btn mr-0 class="success" color="white" small flat @click="create">
        <v-icon dark>add</v-icon> Add
      </v-btn>
    </v-toolbar>

    <v-toolbar color="transparent" class="listing-toolbar" flat ml-0 v-else>
      <v-btn icon>
        <v-icon>work</v-icon>
      </v-btn>
      <v-toolbar-title>Rating </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn mr-0 class="success" color="white" small flat @click="create">
        <v-icon dark>add</v-icon> Add
      </v-btn>
    </v-toolbar>

    <!-- Error Messages -->
    <v-alert :type="message.type==='E' ? 'error' : message.type" :value="true" v-for="(message, index) in messages" :key="index">
      {{ message.text }}
    </v-alert>

    <!-- Empty Check -->
    <div v-if="loading" centered>
      <v-progress-linear :indeterminate="true"></v-progress-linear>
      <p class="text-xs-center">contacting server...</p>
    </div>
    <v-alert type="info" value="true" color="primary" outline icon="info" v-else-if="entities.length===0 && !loading">
      No Ratings found. Would you like to create one now?
      <v-btn color="primary" @click="create">create new</v-btn>
    </v-alert>

    <!-- Table form listing -->
    <v-card v-else>
      <v-data-table class="elevation-0" :headers="headers" :items="entities" :pagination.sync="pagination" :total-items="pagination.totalItems" :rows-per-page-items="pagination.rowsPerPageItems">
        <template slot="items" slot-scope="props">
          <tr :class="{'highlight': recentlySaved(props.item.Id)}" :key="props.item.Id">
            <td class="justify-center layout">
              <v-btn icon class="mx-0" @click.stop="deleteTry(props.item.Id, props.item.Name)">
                <v-icon>delete</v-icon>
              </v-btn>
              <v-btn icon class="mx-0" @click.stop="duplicate(props.item.Id)">
                <v-icon>file_copy</v-icon>
              </v-btn>
              <v-tooltip top v-if="nested">
                <v-btn slot="activator" icon class="mx-0" @click="nestedFormEdit(props.item.Id, props.item.Name)">
                  <v-icon>edit</v-icon>
                </v-btn>
                <span>Edit</span>
              </v-tooltip>
              <v-tooltip top v-else>
                <v-btn slot="activator" icon class="mx-0" :to="{name: 'ratings_edit', params:{ id: props.item.Id }}">
                  <v-icon>edit</v-icon>
                </v-btn>
                <span>Edit</span>
              </v-tooltip>
            </td>
            <!-- Check if it should appear in the list -->
               
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for select types -->
            
            <td>
              <component :is="'ListWidgetSelect'" :field="'Status'" :status=" props.item.Status " />
            </td>       
            <!-- Use different rendering for dates -->
            
            <td>
              <component :is="'ListWidgetTime'" :time=" props.item.Createdat " />
            </td>   
            <!-- Use different rendering for dates -->
            
            <td>
              <component :is="'ListWidgetTime'" :time=" props.item.Updatedat " />
            </td>   
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for select types -->
            <!-- Use different rendering for dates -->
            
            <!-- Use different rendering for toggles -->
            
            <td>
              <v-tooltip top>
                <span slot="activator">
                  {{ props.item.Ratingcount}}

                </span>
                <span>Widget Type : String</span>
              </v-tooltip>
            </td>
                   
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for select types -->
            <!-- Use different rendering for dates -->
            
            <!-- Use different rendering for toggles -->
            
            <td>
              <v-tooltip top>
                <span slot="activator">
                  {{ props.item.Score}}

                </span>
                <span>Widget Type : Int64</span>
              </v-tooltip>
            </td>
                   
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for dates -->
            <!-- Use different rendering for select types -->
            <!-- Use different rendering for dates -->
            
            <!-- Use different rendering for toggles -->
            
            <td>
              <v-tooltip top>
                <span slot="activator">
                  {{ props.item.Reactionsummary}}

                </span>
                <span>Widget Type : String</span>
              </v-tooltip>
            </td>
                  
          </tr>
        </template>

        <template slot="no-data">
          <v-flex ma-4>
            <v-alert slot="no-results" :value="true" color="primary" outline icon="info" v-if="search.length> 0"> Your search for "{{ search }}" found no results.
            </v-alert>
            <v-alert slot="no-results" :value="true" color="primary" outline icon="info" v-else>
              No Rating found.
            </v-alert>
          </v-flex>
        </template>
      </v-data-table>
      <div class="text-xs-center pt-2 pb-4">
        <v-pagination v-model="pagination.page" :length="pages" circle :total-visible="7"></v-pagination>
      </div>
    </v-card>

    <v-dialog v-model="dialog" persistent max-width="300">
      <v-card dark>
        <v-card-title class="headline">Confirm Delete</v-card-title>
        <v-card-text>
          You are about to delete
          <strong>"{{deleteItemId.label}}"</strong>
        </v-card-text>
        <v-card-actions>
          <v-btn flat @click.native="dialog = false;">Cancel</v-btn>
          <v-spacer></v-spacer>
          <v-btn color="error" dark @click.native="deleteConfirm">Yes, Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Edit In Place -->
    <v-dialog
      v-model="nestedForm.visible"
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
      :persistent="true"
    >
      <v-card flat>
        <component :is="'RatingsEdit'" :key="nestedForm.id" @closeform="nestedFormClose" :visible="nestedForm.visible" :filter_id="nestedForm.id" :nested="true" />
      </v-card>
    </v-dialog>
    <!-- !# Edit In Place -->

    <v-snackbar v-model="snackbar.show" :bottom="true" color="primary" auto-height :timeout="2000">
      {{ snackbar.text }}
      <v-btn dark flat @click="snackbarHide">
        <v-icon>close</v-icon>
      </v-btn>
    </v-snackbar>
  </div>
</template>

<script>
import { AdminClient } from "@/services/service_admin_pb_service";
import { ListRequest, GetRequest, Filter} from "@/services/service_admin_pb";
import { CreateRatingRequest } from "@/services/service_admin_pb";
import { DeleteRequest } from "@/services/service_admin_pb";
import { mapGetters } from 'vuex';

import { Rating, ListOpts, ListSortOpts } from "@/services/models_pb.js";
import _ from 'lodash';


let cli = new AdminClient("/api");

export default {
  data() {
    return {
      pagination: {
        sortBy: 'id',
        descending: true,
        page: 1,
        rowsPerPage: 5,
        totalItems: 0,
        rowsPerPageItems: [5, 10, 25, 50, 100]

      },
      messages: [],
      track: null,
      search: "",
      toggle_status:null,
      snackbar: {
          show: false, 
          text: ''
      },
      dialog: null,
      deleteItemId: {
        id: null,
        label: null
      },
      nestedForm: {
        visible: null,
        id: null,
        label: null
      },
      text: "",
      headers: [
        {'text': 'Action', 'value': null, sortable: false, 'width': '80px'},
        {text: "Status", value: "Status", sortable: false},
        {text: "Created", value: "CreatedAt", sortable: false},
        {text: "Updated", value: "UpdatedAt", sortable: false},
        {text: "Rating Count", value: "RatingCount", sortable: false},
        {text: "Score", value: "Score", sortable: false},
        {text: "Reaction Summary", value: "ReactionSummary", sortable: false},
      ],
      entities: [],
      loading: false
    };
  },
  components: {
  },
  computed: {
    ...mapGetters({
      'token': 'auth/getToken'
    }),
    pages() {
      if (this.pagination.rowsPerPage == null ||
        this.pagination.totalItems == null
      ) return 0

      return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
    }
  },
  props: ["filter_id", "filter_field", "nested", "rpc"],
  mounted() {
    this.loading = true;
    this.populateData();
    this.track = this.$route.params.track;

  },
  methods: {    
      recentlySaved(id) {
          if (this.track === id) {
              return true;
          }
          return false;
      },
      nestedFormEdit(editid, name) {
        this.nestedForm = {
          visible: true,
          id: editid,
          label: name
        };
      },
      nestedFormClose() {
        // clear and reset nested form
        this.getList();
        this.nestedForm = {
          visible: false,
          id: "",
          label: ""
        };
      },
      count(req) {
        return new Promise((resolve, reject) => {
          cli.countRatings(req, (err, resp) => {
            if (err) {
              console.log(err);
              this.snackbarShow(err.message, "error");
              reject(err);
            }

            let count = resp.getCount();
            resolve(count);
          })
        })
      },
      populateData: function () {
        this.getList().then((values) => {
          this.entities = values.entities;
          this.pagination.totalItems = values.count
        })
      },
            
      fetchEntities: function(req){
        return new Promise((resolve, reject) => {
          cli.listRatings(req, (err, resp) => {
            if (err) {
              console.log(err);
              this.snackbarShow(err.message, "error");
              return;
            }
            this.loading = false;
            this.snackbarShow('Ratings Loaded');

            let entities = resp.getRatingsList().map(entity => {
              return {
                Id: entity.getId(),
                Status: entity.getStatus(),
                Userid: entity.getUserid(),
                Createdat: entity.getCreatedat().toDate(),
                Updatedat: entity.getUpdatedat().toDate(),
                Ratingcount: entity.getRatingcount(),
                Score: entity.getScore(),
                Reactionsummary: entity.getReactionsummary(),
                
              }
            })
            resolve(entities);
          });
        });
      },
      getList: function () {
        // Will fetch Total and Entities as per pagination request
        let req = new ListRequest();
        let sort = new ListSortOpts();

        req.setKey(this.token);

        if (this.filter_id) {
          let filter = new Filter();
          filter.setField(this.filter_field);
          filter.setOperation("=");
          filter.setValue(this.filter_id);
          req.addFilters(filter);
        }

        // Search term
        let searchFilter = new Filter();
        searchFilter.setField('id');
        searchFilter.setOperation("~");
        searchFilter.setValue('%' + this.search + '%');
        req.addFilters(searchFilter);

        // Status filter
        if (this.toggle_status != null) {
          let statusFilter = new Filter();
          let statusTerm = '';
          statusFilter.setField("status");
          statusFilter.setOperation("=");

          if (this.toggle_status == 1) {
            statusTerm = 'saved'
          }
          else if (this.toggle_status == 2) {
            statusTerm = 'draft'
          }
          else if (this.toggle_status == 3) {
            statusTerm = 'unpublished'
          }
          else if (this.toggle_status == 0) {
            statusTerm = 'published'
          }
          console.log(statusTerm);
          statusFilter.setValue(statusTerm);
          req.addFilters(statusFilter);
        }

        let opts = new ListOpts();
        opts.setOffset((this.pagination.page - 1) * this.pagination.rowsPerPage)
        opts.setLimit(this.pagination.rowsPerPage);

        sort.setAscending(true);
        sort.setField("updated_at");
        opts.addSort(sort);

        req.setListopts(opts);

        return Promise.all([this.count(req), this.fetchEntities(req)]).then(x => {
          return Promise.resolve({ count: x[0], entities: x[1] })
        })
      },      
      duplicate(id) {
        let toBeDuplicatedId = id;
        let toBeDuplicatedEntity = null;
        // Get the entity
        let req = new GetRequest();
        req.setId(toBeDuplicatedId);

        this.loading = true;

        cli.getRating(req, (err, resp) => {
          if (err) {
            console.log(err);
            this.snackbarShow(err.message, "error");
            return;
          }
          this.loading = false;
          ///
          toBeDuplicatedEntity = resp.getRating();
          this.create("existing", toBeDuplicatedEntity);
        });
      },
      create: function(from, existingEntity) {
        let req = new CreateRatingRequest();
        req.setKey(this.token);

        if (from === "existing" && existingEntity) {
          req.setRating(existingEntity);
        } else if (this.nested) {
          let entity = new Rating();
          // entity[this.rpc](this.filter_id);
          req.setRating(entity);
        }

        this.snackbarShow('Loading Ratings');
        this.loading = true;

        cli.createRating(req, (err, resp) => {
          if (err) {
            console.log(err);

            this.snackbarShow("Create Error " + err.message, "error");

            this.loading = false;
            return;
          }
          this.loading = false;
          if (this.nested) {
            this.nestedForm = {
              visible: true,
              id: resp.getRating().getId(),
              label: ""
            };
          } else {
            this.$router.push({
              name: "ratings_edit",
              params: { id: resp.getRating().getId() }
            });
          }
        });
      },
      deleteTry: function(id, label) {
          this.deleteItemId = { id: id, label: label };
          this.dialog = true;
      },
      deleteConfirm: function() {

          let req = new DeleteRequest();
          req.setKey(this.token);
          req.setId(this.deleteItemId.id)

          cli.deleteRating(req, (err, resp) => {
            if (err) {
                console.log(err);
                this.snackbarShow("Delete error: " + err.message, "error");

                this.loading = false;
                return;
            } else {
              this.snackbarShow('Deleted', this.deleteItemId.label);

              if (this.entities.length == 1 && this.pagination.page > 1) {
                this.pagination.page--;
              }

              this.populateData();
            }
          })

          this.dialog = false;

      },
      snackbarShow: function(text, color) {
          this.snackbar.show = true;
          this.snackbar.color = color || 'info';
          this.snackbar.text = text || 'something happened';
      },
      snackbarHide: function() {
          this.snackbar.show = false
      }
  },
  watch: {
    pagination: {
      handler() {
        this.getList().then(values => {
          this.entities = values.entities;
        })
      },
      deep: true
    },
    toggle_status: {
      handler() {
        this.populateData()
      }
    },
    search: _.debounce(function (val) {
        this.populateData()
    }, 500)
  }
};
</script>

<style lang="scss">
.listing-toolbar .v-toolbar__content {
  padding: 0;
}

.listing--container td {
  img {
    display: block !important;
    height: 48px;
    width: auto;
  }
}
</style>


<style lang="scss" scoped>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}

.highlight {
  animation: highlight 10s 1 ease-in-out;
}

@keyframes highlight {
  0% {
    background: #b9f6ca;
  }
  100% {
    background: transparent;
  }
}
</style>
