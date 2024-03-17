import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import LogoutView from '../views/LogoutView.vue'
import RegisterView from '../views/RegisterView.vue'
import GuideView from '../views/GuideView.vue'
import DashboardView from '../views/DashboardView.vue'
import LeaderboardView from "@/views/LeaderboardView.vue";
import profileView from "@/views/ProfileView.vue";
import editProfileView from "@/views/EditProfileView.vue";

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
        },
        {
            path: '/dashboard',
            name: 'dashboard',
            component: DashboardView,
            children: [
                {
                    path: 'profile',
                    component: profileView
                },
                {
                    path: 'edit',
                    component: editProfileView
                }
            ]
        },
        {
            path: '/leaderboard',
            name: 'leaderboard',
            component: LeaderboardView
        }
    ]
})

export default router
