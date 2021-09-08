import Vue from "vue"
import { BootstrapVue, IconsPlugin } from "bootstrap-vue"
import Vuelidate from 'vuelidate'
import App from "./App.vue"
import router from "./router"
// import store from "./store"
import "./assets/scss/index.scss"

Vue.use(Vuelidate)

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
new Vue({
    router,
    render: h => h(App),
}).$mount("#app")
