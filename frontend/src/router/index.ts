import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import logoutView from '../views/logoutView.vue'

const router = createRouter({
    history: createWebHistory("/"),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView
        },
        {
            path: '/logout',
            name: 'logout',
            component: logoutView
        }
    ]
})

export default router
