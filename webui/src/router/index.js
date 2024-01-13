import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import ChangeUsernameView from '../views/ChangeUsernameView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/user/:uid/stream', component: HomeView},
		{path: '/user/:uid/profile', component: ProfileView},
		{path: '/user/:uid/myusername', component: ChangeUsernameView},
	]
})

export default router
