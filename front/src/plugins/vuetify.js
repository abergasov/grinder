import Vue from 'vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)

window.chartColors = {
    color_1: "rgb(54, 162, 235)",
    color_2: "rgb(75, 192, 192)",
    color_3: "rgb(201, 203, 207)",
    color_4: "rgb(255, 159, 64)",
    color_5: "rgb(153, 102, 255)",
    color_6: "rgb(255, 99, 132)",
    color_7: "rgb(255, 205, 86)",
    color_8: "#7B1FA2",
    color_9: "#C2185B",
    color_10: "#5C6BC0",
    color_11: "#1A237E",
    color_12: "#00695C",
    color_13: "#2E7D32",
    color_14: "#EF6C00",
    incoming_color: 'rgba(76, 175, 80, 0.1)',
    limit_color: 'rgba(255, 0, 0, 0.9)',
};

const opts = {
    theme: {
        themes: {
            light: {
                background: '#cccccc',
                primary: '#3f51b5',
                secondary: '#b0bec5',
                accent: '#8c9eff',
                error: '#b71c1c',
                warning: '#fb8c00',
                info: '#2196f3',
                success: '#4caf50',
            },
            dark: {
                background: '#555555',
                primary: '#3f51b5',
                secondary: '#b0bec5',
                accent: '#8c9eff',
                error: '#b71c1c',
                warning: '#fb8c00',
                info: '#2196f3',
                success: '#4caf50',
            }
        },
        dark: true,
    },
}

export default new Vuetify(opts)