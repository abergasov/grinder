<template>
  <div>
    <v-card class="mx-auto" outlined>
        <v-toolbar flat color="green darken-2">
          <v-icon>mdi-account</v-icon>
          <v-toolbar-title class="font-weight-light">User Profile</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn color="green darken-4" fab small @click="isEditing = !isEditing">
            <v-icon v-if="isEditing">mdi-close</v-icon>
            <v-icon v-else>mdi-pencil</v-icon>
          </v-btn>
        </v-toolbar>

        <v-list-item three-line>
          <v-list-item-content>
            <v-row>
              <v-col cols="12" md="4">
                  <v-text-field v-model="firstName" :disabled="!isEditing" label="First name" required></v-text-field>
              </v-col>
              <v-col cols="12" md="4">
                <v-text-field v-model="lastName" :disabled="!isEditing" label="Last name" required></v-text-field>
              </v-col>
                <v-col cols="12" md="4"></v-col>
            </v-row>
          </v-list-item-content>
        </v-list-item>

        <v-card-actions>
            <v-spacer />
            <v-btn :disabled="!formValid()" color="success" @click="updateUser" large>{{ $t('update') }}</v-btn>
        </v-card-actions>
    </v-card>
    <v-card class="mx-auto pass_change" dark>
      <v-toolbar flat>
        <v-icon>mdi-key-variant</v-icon>
        <v-toolbar-title class="font-weight-light">Update password</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn color="#385F73 darken-4" fab small @click="changePass = !changePass">
          <v-icon v-if="changePass">mdi-close</v-icon>
          <v-icon v-else>mdi-pencil</v-icon>
        </v-btn>
      </v-toolbar>

      <v-list-item v-if="changePass" three-line>
        <v-list-item-content>
          <v-row>
            <v-col cols="12" md="12">
              <v-text-field v-model="oldPass" :disabled="pending" label="Current password" required></v-text-field>
            </v-col>
            <v-col cols="12" md="12">
              <v-text-field v-model="newPass" :disabled="pending" label="New password" required></v-text-field>
            </v-col>
            <v-col cols="12" md="12">
              <v-text-field v-model="reTypeNewPass" :disabled="pending" label="ReType new password" required></v-text-field>
            </v-col>
          </v-row>
        </v-list-item-content>
      </v-list-item>

      <v-card-actions v-if="changePass">
        <v-spacer />
        <v-btn :disabled="!formPassValid()" color="success" @click="updatePass" large>{{ $t('update_password') }}</v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
  export default {
    name: "AppProfile",
    data () {
      return {
        pending: false,
        firstName: '',
        lastName: '',
        isEditing: false,
        changePass: false,

        oldPass: '',
        newPass: '',
        reTypeNewPass: '',
      }
    },
    created() {
      this.loadAll();
    },
    methods: {
      formPassValid() {
        return !this.pending && this.oldPass.length > 0 && this.newPass.length > 0 && this.newPass === this.reTypeNewPass;
      },

      formValid() {
        return this.isEditing && this.firstName.length > 0 && this.lastName.length > 0;
      },

      loadAll() {
        this.pending = true;
        this.askBackend('data/profile', {}).then(
          data => {
            this.pending = false;
            if (data.ok) {
              this.firstName = data.user.first_name;
              this.lastName = data.user.last_name;
            }
          },
          e => console.error(e),
        )
      },

      updatePass() {
        this.pending = true;
        this.askBackend('data/profile/update_password', {
          old_pass: this.oldPass,
          new_pass: this.newPass
        }).then(
          data => {
            this.pending = false;
            !data.ok ? this.showError(data.error):this.showOk("Password changed");
          },
          e => console.error(e),
        )
      },

      updateUser() {
        this.pending = true;
        this.askBackend('data/profile/update', {
          first_name: this.firstName,
          last_name: this.lastName,
        }).then(
          data => {
            this.pending = false;
            if (data.ok) {
              this.firstName = data.user.first_name;
              this.lastName = data.user.last_name;
            }
          },
          e => console.error(e),
        )
      },
    }
  }
</script>

<style scoped>
  .pass_change {
    margin-top: 20px;
  }
</style>