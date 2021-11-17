import Vue from 'vue'
import Vuetify from 'vuetify/lib/framework'

Vue.use(Vuetify)

const light = {
  primary: '#2196f3',
  secondary: '#00bcd4',
  accent: '#3f51b5',
  error: '#f44336',
  warning: '#ff5722',
  info: '#ff9800',
  success: '#4caf50',
  reset: '#684bff'
}

export default new Vuetify({
  theme: {
    dark: false,
    themes: {
      light: light
    }
  }
})
