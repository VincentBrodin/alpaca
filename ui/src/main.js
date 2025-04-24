import './assets/main.css'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { createApp, h, provide } from 'vue';
import router from './router'
import apolloClient from './apollo';
import App from './App.vue'
import { DefaultApolloClient } from '@vue/apollo-composable';


createApp({
	setup() {
		provide(DefaultApolloClient, apolloClient);
	},
	render: () => h(App),
}).use(router).mount('#app');
