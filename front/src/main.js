// Import Vue
import Vue from 'vue';
import VueRouter from 'vue-router';

// Import Vue App, routes, store
import App from './App';
import routes from './routes';
import store from './store';
import i18n from './i18n'
import vuetify from '@/plugins/vuetify' // path to vuetify export
Vue.use(VueRouter);
Vue.use(require('vue-moment'));
const axios = require('axios').default;
const Chart = require('chart.js');

import Fingerprint2 from "fingerprintjs2"
let hash = require('object-hash');

window.userId = -1;
if (window.requestIdleCallback) {
    requestIdleCallback(function () {
        Fingerprint2.get(function (components) {
            window.userId = hash(components);
            runApp();
        })
    })
} else {
    setTimeout(function () {
        Fingerprint2.get(function (components) {
            window.userId = hash(components);
            runApp();
        })
    }, 500)
}

Vue.prototype.parseJwt = function (token) {
    let base64Url = token.split('.')[1];
    let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    let jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

Vue.prototype.refreshManager = function () {
    let expires = this.$store.state.auth_expires;
    let auth = +this.$store.state.auth;
    if (auth !== 1) {
        return;
    }
    let nowTime = + new Date();
    let diff = (expires * 1000) - nowTime;
    diff = diff / 60000;
    if (diff > 1) {
        return;
    }

    axios.post(`/api/auth/refresh`, {user_id: this.$store.state.auth_user}, {
        headers: {
            m: window.userId,
        }
    })
        .then(resp => {
            if (resp.data.ok) {
                let rawData = this.parseJwt(resp.data.token);
                this.$store.commit('setUserId', rawData.user_id);
                this.$store.commit('setAuthExpires', rawData.exp);
                this.$store.commit('setAuth', 1);
            }
        })
        .catch(() => {
            this.$store.commit('setAuth', 0);
            this.$store.commit('setAlert', {display: true, text: 'Unauthorized', color: 'error'});
        })
}

Vue.prototype.askBackend = function (url, param) {
    console.log('ask smth!', process.env.BACK_SERVER);
    param.user_sign = window.userId;
    let config = {
        headers: {
            m: window.userId,
        }
    }
    this.$store.commit('setLoading', true);
    return new Promise((resolve, reject) => {
        axios.post(`/api/${url}`, param, config)
            .then(resp => {
                this.$store.commit('setLoading', false);
                resolve(resp.data)
            })
            .catch(error => {
                this.$store.commit('setLoading', false);
                let code = +error.response.status;
                let message = ''
                switch (code) {
                    case 401:
                        this.$store.commit('setAuth', 0);
                        message = 'Unauthorized';
                        break;
                    case 409:
                        message = 'Already exist';
                        break;
                    case 400:
                        message = 'Bad request';
                        break;
                }
                if (message) {
                    this.$store.commit('setAlert', {
                        display: true,
                        text: message,
                        color: 'error',
                    });
                }
                reject(error)
            })
    });
};

Vue.prototype.showError = function (text, duration = 5) {
    this.$store.commit('setAlert', {
        display: true,
        text: text,
        delay: duration,
        color: 'error'
    });
}

Vue.prototype.showOk = function (text, duration = 5) {
    this.$store.commit('setAlert', {
        display: true,
        text: text,
        delay: duration,
        color: 'success'
    });
}

// Configure router
const router = new VueRouter({
    routes,
    linkActiveClass: 'active',
    mode: 'history'
});

console.log(process.env.BACK_SERVER);

function runApp () {
    window.app = new Vue({
        vuetify,
        i18n,
        store,
        el: '#app',
        render: h => h(App),
        router
    });
    window.app.refreshManager();
    setInterval(() => {
        // check every 10sec
        window.app.refreshManager();
    }, 180000);
}
