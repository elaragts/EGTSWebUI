import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import LogoutView from '../views/LogoutView.vue'
import RegisterView from '../views/RegisterView.vue'
import GuideView from '../views/GuideView.vue'

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
            component: LogoutView
        },
        {
            path: '/register',
            name: 'register',
            component: RegisterView
        },
        {
            path: '/guide',
            name: 'guide',
            component: GuideView
        }
    ]
})

export default router
