<template>
  <v-row no-gutters>
    <v-col cols="12" sm="12">
      <v-card>
        <v-card-title>
          <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Поиск"
              single-line
              hide-details
          ></v-text-field>
        </v-card-title>
        <v-data-table
            :loading=pending
            :headers="headers"
            :items="users"
            :search="search"
        >
          <template v-slot:item.rights="{ item }">
            <v-chip class="role_chip" v-for="r in item.rights" :color="getRoleColor(r)" dark>
              {{ getRoleText(r) }}
            </v-chip>
          </template>
          <template v-slot:item.active="{ item }">
            <v-simple-checkbox v-model="item.active" disabled></v-simple-checkbox>
          </template>
          <template v-slot:top>
            <v-toolbar flat>
              <v-dialog v-model="dialog" max-width="500px">
                <v-card>
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="6">
                          <v-text-field v-model="editedItem.email" :disabled="pending"  label="Email"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                          <v-checkbox label="User active" :disabled="pending" v-model="editedItem.active"></v-checkbox>
                        </v-col>
                        <v-col cols="12" sm="6">
                          <v-text-field v-model="editedItem.first_name" :disabled="pending" label="First name"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                          <v-text-field v-model="editedItem.last_name" label="Last name" :disabled="pending"></v-text-field>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" :disabled="pending" text @click="close">Cancel</v-btn>
                    <v-btn color="blue darken-1" :disabled="!updateValid()" text @click="save">Save</v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
          </template>
          <template v-slot:no-data>
            <v-btn color="primary" @click="initialize">Reset</v-btn>
          </template>
        </v-data-table>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "AppUsers",
  data () {
    return {
      dialog: false,
      dialogDelete: false,
      pending: false,
      search: "",
      headers: [
        { text: 'Active', align: 'start', value: 'active' },
        { text: 'First name', align: 'start', value: 'first_name' },
        { text: 'Last name', value: 'last_name' },
        { text: 'Email', value: 'email' },
        { text: 'Register data', value: 'register_date' },
        { text: 'Roles', value: 'rights' },
        { text: 'Actions', value: 'actions', sortable: false },
      ],
      users: [],
      editedIndex: -1,
      editedItem: {
        user_id: 0,
        email: '',
        first_name: '',
        last_name: '',
        active: false,
      },
      defaultItem: {
        user_id: 0,
        email: '',
        first_name: '',
        last_name: '',
        active: false,
      },
      formTitle: '',
      rightsMap: [],
    }
  },
  created() {
    this.loadAll();
    this.loadRights()
  },
  methods: {
    loadRights() {
      this.askBackend('data/users/roles/list', {}).then(
          data => {
            if (data.ok) {
              this.rightsMap = data.rights;
            }
          },
          err => {
            console.error(err);
            this.showError("error load user rights");
          },
      );
    },
    loadAll() {
      this.pending = true;
      this.askBackend('data/users/list', {}).then(
          data => {
            this.pending = false;
            if (!data.ok) {
              this.showError("Error load users");
              return
            }
            this.prepareUsers(data.persons, data.person_rights)
          },
          e => {
            console.error(e);
            this.showError("Error load users");
          }
      )
    },

    prepareUsers(userData, userRights) {
      userData.forEach(user => {
        user.rights = [];
        userRights.forEach(uRights => {
          if (uRights.user_id !== user.user_id) {
            return
          }
          user.rights.push(uRights.right_id);
        })
      });
      this.users = userData;
    },

    getRoleColor(rId) {
      let name = this.rightsMap[+rId];
      switch (name) {
        case "admin":
          return "red"
        case "moderator":
          return "orange";
      }
      return "green";
    },
    getRoleText(rId) {
      return this.rightsMap[+rId];
    },

    editItem (item) {
      this.editedIndex = this.users.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.formTitle = `Edit user: ${this.editedItem.email} - ${this.editedItem.user_id}`;
      this.dialog = true
    },

    close () {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    updateValid() {
      return !this.pending &&
          this.editedItem.email.length > 0 &&
          this.editedItem.first_name.length > 0 &&
          this.editedItem.last_name.length > 0;
    },

    save() {
      this.pending = true;
      this.askBackend('data/users/update', {
        user_id: this.editedItem.user_id,
        first_name: this.editedItem.first_name,
        last_name: this.editedItem.last_name,
        email: this.editedItem.email,
        active: this.editedItem.active,
      }).then(
          data => {
            this.pending = false;
            if (data.ok) {
              this.showOk(data.msg);
              this.close();
              this.loadAll();
            } else {
              this.showError("error while update");
            }
          },
          e => {
            console.error(e);
            this.showError("Error update user");
          }
      )
    },
  }
}
</script>

<style scoped>
  .role_chip {
    margin: 5px;
  }
</style>