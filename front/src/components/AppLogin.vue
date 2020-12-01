<template>
  <v-main>
    <v-container class="fill-height" fluid>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="4">
          <v-card class="elevation-12">
            <v-toolbar  flat>
                <v-toolbar-title>{{ $t((loginMode ? 'login' : 'register')) }}</v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn class="ma-2" :outlined="loginMode" @click="switchMode" color="indigo">Register</v-btn>
                <v-btn class="ma-2" :outlined="!loginMode" @click="switchMode" color="indigo">Sign in</v-btn>
            </v-toolbar>
            <v-card-text>
                <v-form v-model="formValid">
                    <v-text-field label="Email"
                          name="login"
                          type="email"
                          required
                          v-model="email"
                          :rules="emailRules"
                          prepend-icon="mdi-account"></v-text-field>
                    <v-text-field
                        id="password"
                        :label="$t('password')"
                        name="password"
                        prepend-icon="mdi-lock"
                        required
                        :rules="passwordRules"
                        v-model="pass"
                        type="password"
                    ></v-text-field>
                    <v-text-field v-if="!loginMode"
                        id="password_retype"
                        :label="$t('password_retype')"
                        name="password_retype"
                        prepend-icon="mdi-lock"
                        required
                        v-model="pass_re"
                        type="password"
                    ></v-text-field>
                </v-form>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="submitForm" :disabled="!isValid">{{ $t((loginMode ? 'login' : 'register')) }}</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
    export default {
        name: "AppLogin",
        data () {
            return {
                formValid: false,
                loginMode: true,
                email: 'a@a.ru',
                pass: '123123',
                pass_re: '123123',
                emailRules: [
                    v => !!v || 'E-mail is required',
                    v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
                ],
                passwordRules: [
                    v => !!v || 'Password is required',
                    v => (v && v.length >= 5) || 'Password must have 5+ characters',
                ],
            }
        },
        computed: {
            isValid() {
                if (!this.formValid) {
                    return false;
                }
                if (!this.loginMode) {
                    if (this.pass !== this.pass_re) {
                        return false;
                    }
                }
                return true;
            }
        },
        methods: {
            switchMode() {
                this.loginMode = !this.loginMode;
            },
            submitForm() {
                let url = this.loginMode ? 'auth/login' : 'auth/register';
                this.askBackend(url, {
                    email: this.email,
                    password: this.pass,
                    re_password: this.pass_re,
                }).then(
                    resp => {
                        if (resp.ok) {
                            let rawData = this.parseJwt(resp.token);
                            this.$store.commit('setUserId', rawData.user_id);
                            this.$store.commit('setAuthExpires', rawData.exp);
                            setTimeout(() => {
                                this.$store.commit('setAuth', 1);
                            }, 100);
                        }
                    });
            }
        }
    }
</script>

