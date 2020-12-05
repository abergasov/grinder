<template>
  <v-container>
    <v-row no-gutters>
      <v-col cols="12" sm="12" md="4" lg="3">
        <h1>Tasks</h1>
      </v-col>

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
              :items="bots"
              :search="search"
          ></v-data-table>
        </v-card>
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
  export default {
    name: 'AppHome',
    data () {
      return {
        search: '',
        pending: false,
        headers: [
          { text: 'Имя', align: 'start', value: 'first_name' },
          { text: 'Фамилия', value: 'last_name' },
          { text: 'Логин', value: 'login' },
          { text: 'Дата регистрации', value: 'register_date' },
        ],
        bots: [

        ]
      }
    },
    created() {
      this.loadAll();
    },
    methods: {
      loadAll() {
        this.pending = true;

      },
    }
  };
</script>

<style scoped lang="scss">
    h1  {
        color: green;
    }

</style>