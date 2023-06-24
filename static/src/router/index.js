import { createRouter, createWebHistory } from "vue-router"
import Home from "@/view/home.vue"
const routes = [
    {
        path: '/home',
        name: 'home',
        component: Home
    },
    {
        path: '/about',
        name: 'about',
        component: () => import('../view/about.vue')
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router