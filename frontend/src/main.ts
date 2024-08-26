import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'
import { createPinia } from 'pinia'
import Toast from 'vue-toastification';
import 'vue-toastification/dist/index.css';

// imports for external vue components
import { createVuetify } from 'vuetify';
import '@mdi/font/css/materialdesignicons.css'
import { VSlider, VDataTable, VTextField, VCard, VDialog, VBtn, VSpacer, VIcon } from 'vuetify/components';

const vuetify = createVuetify({
    components: {
        VSlider,
        VDataTable,
        VTextField,
        VCard,
        VDialog,
        VBtn,
        VSpacer,
        VIcon
    }
});

const app = createApp(App);
const pinia = createPinia();

app.use(router)
app.use(pinia)
app.use(Toast, {position: 'bottom-right'});
app.use(vuetify)

app.mount('#app')
