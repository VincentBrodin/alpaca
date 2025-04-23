<script setup>
	import Navbar from './components/Navbar.vue';
	import {ApolloClient, createHttpLink, InMemoryCache} from '@apollo/client/core'
	import {provideApolloClient, useQuery} from '@vue/apollo-composable';
	import {CHATS} from './queries';
	import {store} from './store.js'
	import {watch} from 'vue';
	import {useLocalStorage} from '@vueuse/core';
	import {RouterView} from 'vue-router';

	const drawerOpen = useLocalStorage('drawer-open', true);

	const httpLink = createHttpLink({
		uri: 'http://localhost:1345/query',
	})

	const cache = new InMemoryCache()

	const apolloClient = new ApolloClient({
		link: httpLink,
		cache,
	})
	provideApolloClient(apolloClient);
	const {result, loading, error} = useQuery(CHATS);
	watch(result, (newResult) => {
		const sortedChats = [...newResult.chats].sort(
			(a, b) => Date.parse(b.last_changed) - Date.parse(a.last_changed)
		)
		store.chats = sortedChats
	})
</script>


<template>
	<div class="flex w-full h-full">
		<!-- Sidebar -->
		<div class="transition-all duration-300 ease-in-out bg-base-200 overflow-hidden"
			:class="{ 'min-w-80': drawerOpen, 'min-w-0 w-0': !drawerOpen }">
			<div v-if="loading" class="flex justify-center items-center">
				<span class="loading loading-spinner loading-md"></span>
			</div>
			<p v-else-if="error">ERROR</p>
			<ul v-else class="menu text-base-content min-h-full w-full flex items-end">
				<li><a class="btn btn-circle" @click="drawerOpen = false"><i
							class="bi bi-arrow-bar-left text-sm"></i></a>
				</li>
				<li v-for="chat in store.chats" :key="chat.id" class="w-full">
					<router-link :to="`/chat/${chat.id}`" class="w-full">
						<p class="truncate  whitespace-nowrap">{{ chat.title }}</p>
					</router-link>
				</li>
			</ul>
		</div>

		<!-- Main Content -->
		<div class="flex p-4 transition-all duration-300 ease-in-out grow"
			:class="{ 'ml-0': drawerOpen, 'ml-0': drawerOpen }">
			<Navbar />
			<router-view></router-view>
		</div>
	</div>
</template>
