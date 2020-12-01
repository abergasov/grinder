import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        auth: 0,
        auth_expires: 0,
        auth_user: 0,
        dataLoading: false,
        categories_expenses: null,
        categories_incoming: null,
        categories_assets: null,
        expenses: [],
        debts: [],
        alertData: {
            display: false,
            text: '',
            color: '',
            delay: 0,
        },
        el_chart: {},
        per_days_chart: {},
        statistic_raw_config: JSON.parse(localStorage.getItem('statistic_raw_config')) || {},
    },
    mutations: {
        setRawStatisticConfig(state, payload) {
            state.statistic_raw_config = payload;
            localStorage.setItem('statistic_raw_config', JSON.stringify(payload));
        },
        setLoading(state, payload) {
            state.dataLoading = payload;
        },

        setUserId(state, payload) {
            state.auth_user = +payload;
        },
        setDebts(state, payload) {
            state.debts = payload || [];
        },
        setAuthExpires(state, payload) {
            state.auth_expires = +payload;
        },
        setAuth (state, payload) {
            state.auth = payload;
            localStorage.auth = +payload;
        },
        setExpenses (state, payload) {
            state.expenses = payload;
        },
        setElChart (state, payload) {
            state.el_chart = payload;
        },
        setPerDays (state, payload) {
            state.per_days_chart = payload;
        },
        setCategoriesExpenses (state, payload) {
            state.categories_expenses = payload;
        },
        setCategoriesIncoming (state, payload) {
            state.categories_incoming = payload;
        },
        setCategoriesAssets (state, payload) {
            state.categories_assets = payload;
        },
        setAlert (state, payload) {
            console.log('set alert');
            state.alertData = payload;
            setTimeout(() => {
                state.alertData.display = false;
                console.log('clear alert');
            }, (payload.delay || 5) * 1000)
        }
    },
    getters: {
        authGetter: state => {
            return state.auth;
        },
        statisticRawConfig: state => {
            return state.statistic_raw_config;
        }
    }
})

export default store;
