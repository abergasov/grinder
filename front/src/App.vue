<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" v-if="auth" app>
      <v-list-item @click="drawer = !drawer" class="mini_swapper">
        <v-list-item-content>
          <v-list-item-title class="title">Bot orchestrator</v-list-item-title>
          <v-list-item-subtitle>master of puppets</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list dense nav>
        <v-list-item v-for="item in items" :key="item.title" link :to="item.link">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ $t(item.title) }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <div class="pa-2">
          <v-btn block @click="logout"><v-icon>mdi-logout</v-icon>Logout</v-btn>
        </div>
      </template>
    </v-navigation-drawer>
    <v-app-bar id="app-bar" absolute app color="transparent" flat height="75">
      <v-app-bar-nav-icon v-if="auth" @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title v-if="auth" class="hidden-sm-and-down font-weight-light" v-text="$t($route.name)"/>
      <v-spacer />

      <div class="mx-3" />

      <v-btn icon @click="setMode">
        <v-icon v-if="darkMode">mdi-brightness-7</v-icon>
        <v-icon v-else>mdi-brightness-4</v-icon>
      </v-btn>
      <div class="text-center flag_wrap">
        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn tile outlined color="success" v-bind="attrs" v-on="on">
              <v-icon left>mdi-flag</v-icon>
              {{ currentLang.code }}
            </v-btn>
          </template>
          <v-list>
            <v-list-item v-for="(item, index) in lang" :key="index" @click="setLanguage(item.code)">
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
      <div style="display: inline-flex" v-if="auth">
      <v-btn class="ml-2" min-width="0" text to="/">
        <v-icon>mdi-view-dashboard</v-icon>
      </v-btn>
      <v-btn class="ml-2" min-width="0" text to="/statistic">
        <v-icon>mdi-chart-areaspline</v-icon>
      </v-btn>
      <v-btn class="ml-2" min-width="0" text to="/profile">
        <v-icon>mdi-account</v-icon>
      </v-btn>
      </div>
    </v-app-bar>
    <app-login v-if="!auth"></app-login>
    <v-main v-else>
      <v-container fluid>
        <router-view></router-view>
        <v-footer id="dashboard-core-footer" app>
          <v-container>
            <v-row align="center" no-gutters>
              <v-col class="text-center mb-sm-0 mb-5" cols="auto"></v-col>
              <v-spacer class="hidden-sm-and-down" />

              <v-col cols="12" md="auto">
                <div class="body-1 font-weight-light pt-6 pt-md-0 text-center">
                  Активных ботов: {{ footerInfo.total_tasks }}
                </div>
              </v-col>
            </v-row>
          </v-container>
        </v-footer>
      </v-container>
    </v-main>
    <v-snackbar v-if="alertData.display" :value="true" bottom :color="alertData.color" outlined right>{{ alertData.text }}</v-snackbar>
  </v-app>
</template>

<script>
  import AppLogin from "./components/AppLogin";
  export default {
    components: {AppLogin},
    data () {
      return {
        bottomNav: null,
        footerInfo: {
          total_tasks: null,
        },
        drawer: true,
        value: true,
        items: [
          { title: 'dashboard', icon: 'mdi-view-dashboard', link: '/' },
          { title: 'statistic', icon: 'mdi-chart-areaspline', link: '/statistic' },
          { title: 'profile', icon: 'mdi-account-box', link: '/profile'},
          { title: 'users', icon: 'mdi-account-box', link: '/users'},
        ],
        darkMode: false,
        lang: [
          { title: 'English', code: 'en', icon: '' },
          { title: 'Russian', code: 'ru', icon: '' },
        ],
        currentLang: {
          title: '',
          code: '',
          icon: '',
        }
      }
    },
    computed: {
      auth() {
        return +this.$store.state.auth === 1;
      },
      alertData() {
        return this.$store.state.alertData;
      }
    },
    mounted () {
      this.$store.watch(() => this.$store.getters.authGetter, n => {
        if (n === 1) {
          this.initApp();
        }
      })
    },
    created() {
      let currentSize = this.$vuetify.breakpoint.name;
      if (currentSize === 'xs' || currentSize === 'sm') {
        this.drawer = false;
      }

      this.$store.commit('setAuth', localStorage.auth || 0);

      this.setLanguage(localStorage.lang || 'en');
      this.darkMode = +localStorage.dark === 1;
      this.$vuetify.theme.dark = this.darkMode;
      if (this.auth) {
        setTimeout(this.initApp, 1000)
      }
    },
    methods: {
      initApp() {
        this.loadBulkData();
      },
      logout() {
        this.askBackend('auth/logout', {}).then(() => {
          this.$store.commit('setAuth', 0);
        })
      },

      setMode(){
        this.darkMode = !this.darkMode;
        this.$vuetify.theme.dark = this.darkMode;
        localStorage.dark = (this.darkMode ? 1 : 0);
      },

      setLanguage(code) {
        let langCode = 0;
        let userLang = this.lang.findIndex(l => l.code === code);
        if (userLang) {
          langCode = userLang;
        }
        this.currentLang = this.lang[langCode];
        this.$i18n.locale = this.currentLang.code;
        this.$vuetify.lang.current = this.currentLang.code;
        localStorage.lang = code;

      },

      loadBulkData() {
        // this.askBackend('data/bot', {}).then(
        //     data => {
        //       if (data.ok) {
        //         this.footerInfo.total_tasks = data.data.bot_count;
        //       }
        //     },
        //     e => console.error(e),
        // )
      }

    }
  }
</script>

<style scoped>
  .flag_wrap {
    padding-right: 8px;
    padding-left: 16px;
  }
.mini_swapper{
  cursor: pointer;
}
</style>