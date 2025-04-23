<script setup>
	import ChatArea from '@/components/ChatArea.vue';
	import Message from '@/components/Message.vue';
	import { CHAT, SEND } from '@/queries';
	import { useQuery, useSubscription } from '@vue/apollo-composable';
	import {computed, onUnmounted, ref, watch} from 'vue';
	import {useRoute} from 'vue-router'

	const route = useRoute()

	const chatId = computed(() => route.params.id)

	watch(chatId, (newId) => {
		qrefetch({id: newId})
		console.log(qresult.value)
	})


	const chat = ref(null)
	const currentMsg = ref(null)

	const currentContent = ref("")
	const enabled = ref(false)

	const {result: qresult, loading: qloading, error: qerror, refetch: qrefetch} = useQuery(CHAT, () => ({
		id: chatId.value,
	}))

	watch(qresult, (newResult) => {
		if (newResult?.chat) {
			chat.value = JSON.parse(JSON.stringify(newResult.chat)) // deep clone
		}
	})


	const {result: sresult, stop} = useSubscription(
		SEND,
		() => ({
			id: chatId.value,
			model: 'mistral',
			content: currentContent.value,
		}),
		{enabled: () => enabled.value}
	)
	
	watch(sresult, ({response}) => {
		currentMsg.value.content += response.message.content
		}
	})


	function send(input) {
		currentContent.value = input.content
		chat.value.messages.push({role: "user", content: input.content})
        currentMsg.value = {role: "assistant", content: ""}
        enabled.value = true
	}

		onUnmounted(() => {
		stop()
	})

</script>

<template>
	<div class="grow relative flex flex-col min-h-0">
		<div v-if="qloading" class="flex justify-center items-center">
			<span class="loading loading-spinner loading-xl"></span>
		</div>
		<div v-else-if="qerror" class="w-full h-full flex justify-center items-center">

		</div>
		<div v-else class="overflow-y-scroll grow min-h-0">
			<div class="w-full max-w-4xl mx-auto px-4 flex flex-col gap-4 pb-24">
				<Message v-for="message in chat.messages" :message="message" />
				<Message v-if="currentMsg != null" :message="currentMsg" />
			</div>
		</div>

		<ChatArea v-if="!qloading" @send="send" />
	</div>
</template>
