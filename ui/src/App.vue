<script setup>
	import Navbar from './components/Navbar.vue';
	import {useQuery} from '@vue/apollo-composable';
	import {CHATS} from './queries';
	import {store} from './store.js'
	import {computed, watch} from 'vue'
	import {useLocalStorage} from '@vueuse/core';
	import {RouterView, useRoute} from 'vue-router';
	import CreateChat from './components/CreateChat.vue';
	import ChatButton from './components/ChatButton.vue';

	const drawerOpen = useLocalStorage('drawer-open', true);
	const route = useRoute()
	const chatId = computed(() => route.params.id)


	const {loading, error, result, refetch} = useQuery(CHATS, {
		fetchPolicy: 'no-cache' // Optional: to always get fresh data
	});

	watch(result, (newResult) => {
		console.log("WATCHER GOING OFF");
		if (newResult && newResult.chats) {
			const sortedChats = [...newResult.chats].sort(
				(a, b) => Date.parse(b.last_changed) - Date.parse(a.last_changed)
			);
			store.chats = sortedChats;
		} else {
			console.log("No chats found in the response or response is invalid:", newResult);
		}
	});

	function update() {
		refetch();
	}

</script>

<template>
	<div class="flex w-full h-full">
		<!-- Sidebar -->
		<div class="transition-all duration-300 ease-in-out bg-base-200 overflow-hidden"
			:class="{ 'min-w-80 max-w-80': drawerOpen, 'min-w-0 w-0 max-w-0': !drawerOpen }">
			<div v-if="loading" class="flex justify-center items-center">
				<span class="loading loading-spinner loading-md"></span>
			</div>
			<p v-else-if="error">ERROR {{error}}</p>
			<ul v-else class="menu text-base-content min-h-full w-full flex items-end">
				<li>
					<a class="btn btn-circle" @click="drawerOpen = false">
						<i class="bi bi-arrow-bar-left text-sm"></i>
					</a>
				</li>
				<!-- Add -->
				<CreateChat @created="update" />
				<!-- Chats -->
				<ChatButton v-for="chat in store.chats" :key="chat.id" :chat="chat" :selected="chat.id == chatId"
					@removed="update" />
			</ul>
		</div>

		<!-- Main Content -->
		<div class="flex p-4 transition-all duration-300 ease-in-out grow" :class="{ 'ml-0': drawerOpen }">
			<Navbar />
			<RouterView></RouterView>
		</div>
	</div>
</template>
