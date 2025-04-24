<script setup>
	import ChatArea from '@/components/ChatArea.vue';
	import Message from '@/components/Message.vue';
	import {CHAT, SEND} from '@/queries';
	import {useQuery, useSubscription} from '@vue/apollo-composable';
	import {computed, nextTick, onMounted, ref, watch} from 'vue';
	import {useRoute} from 'vue-router'

	const route = useRoute()

	const chatId = computed(() => route.params.id)


	const chat = ref(null)
	const current = ref(null)
	const anchor = ref(null);
	const query = ref("")

	const {result: qresult, loading: qloading, error: qerror} = useQuery(CHAT, () => ({
		id: chatId.value,
	}))

	watch(qresult, (newResult) => {
		if (newResult?.chat) {
			chat.value = JSON.parse(JSON.stringify(newResult.chat)) // deep clone
		}
	})

	const generating = ref(false)
	const enabled = ref(false)
	const {onResult, loading, onError} = useSubscription(
		SEND,
		() => ({
			id: chatId.value,
			model: 'gpt-4o-mini',
			content: query.value,
		}),
		() => ({
			enabled: enabled.value
		}))

	onResult(async (result, _) => {
		if (current.value) {
			if (result.data.send.done) {
				chat.value.messages.push(current.value);
				current.value = null;
				query.value = ""
				generating.value = false;
			}
			else {
				current.value.content += result.data.send.content;
				await nextTick();
			}
		}
		anchor.value.scrollIntoView({behavior: 'smooth'});

	})

	onError(({error}) => {
		console.error("Subscription error:", error);
	});

	watch(() => loading, val => {
		if (val) {
			anchor.value.scrollIntoView({behavior: 'smooth'});
		}
	});

	function send(input) {
		generating.value = true;
		query.value = input.content
		chat.value.messages.push({role: "user", content: input.content})
		current.value = {role: "assistant", content: ""}
		anchor.value.scrollIntoView({behavior: 'smooth'});
		enabled.value = true
	}


	watch(() => chat.value?.messages, () => {
		nextTick(() => {
			anchor.value?.scrollIntoView({behavior: 'instant'});
		});
	}, {immediate: true});

	onMounted(() => {
		enabled.value = false
	})

</script>

<template>
	<div class="grow relative flex flex-col min-h-0 min-w-0">
		<div v-if="qloading" class="flex justify-center items-center">
			<span class="loading loading-spinner loading-xl"></span>
		</div>
		<div v-else-if="qerror" class="w-full h-full flex justify-center items-center">

		</div>
		<div v-else class="overflow-y-scroll grow min-h-0">
			<div class="w-full max-w-4xl min-w-0 mx-auto px-4 flex flex-col gap-4 pb-24">
				<Message v-for="(message, id) in chat.messages" :message="message" :key="id" />
				<span v-if="loading" class="loading loading-spinner loading-md"></span>
				<Message v-if="current != null" :message="current" :key="'current'" />
				<div ref="anchor"></div>
			</div>
		</div>

		<ChatArea v-if="!qloading" @send="send" :generating="generating" />
	</div>
</template>
