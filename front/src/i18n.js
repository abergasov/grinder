import Vue from 'vue'
import VueI18n from 'vue-i18n'

import en from 'vuetify/lib/locale/en'
import ru from 'vuetify/lib/locale/ru'

Vue.use(VueI18n)

const messages = {
    en: {
        ...require('@/locales/en.json'),
        $vuetify: en,
    },
    ru: {
        ...require('@/locales/ru.json'),
        $vuetify: ru,
    },
}
// const messages__ = {
//     en: {
//         ...require('@/locales/en.json'),
//         $vuetify: en,
//     },
//     ru: {
//         ...require('@/locales/ru.json'),
//         $vuetify: ru,
//     },
// }
// const messages = {
//     en: require('@/locales/en.json'),
//     ru: require('@/locales/ru.json'),
// }

// console.log(messages__);

export default new VueI18n({
    locale: process.env.VUE_APP_I18N_LOCALE || 'en',
    fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || 'en',
    messages,
})