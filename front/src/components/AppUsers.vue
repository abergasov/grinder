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
      pending: false,
      search: "",
      headers: [
        { text: 'First name', align: 'start', value: 'first_name' },
        { text: 'Last name', value: 'last_name' },
        { text: 'Email', value: 'email' },
        { text: 'Register data', value: 'register_date' },
        { text: 'Roles', value: 'rights' },
      ],
      users: [],
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
  }
}
</script>

<style scoped>
  .role_chip {
    margin: 5px;
  }
</style>