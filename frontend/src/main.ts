import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'
import { createPinia } from 'pinia'
import Toast from 'vue-toastification';
import 'vue-toastification/dist/index.css';


const app = createApp(App);
const pinia = createPinia();

app.use(router)

app.use(pinia)
app.use(Toast, {position: 'bottom-right'});

app.mount('#app')
