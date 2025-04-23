import { createRouter, createWebHistory } from 'vue-router';
import Chat from './views/Chat.vue';

const routes = [
	{ path: '/chat/:id', name: 'chat', component: Chat },
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
