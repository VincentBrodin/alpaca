<script setup>
	import {DELETE_CHAT} from '@/queries';
	import {useMutation} from '@vue/apollo-composable';
	import {ref} from 'vue';


	const props = defineProps({
		chat: Object,
		selected: Boolean
	});


	const emits = defineEmits(["removed"])
	const id = ref(props.chat.id)

	const {mutate, loading, error, onDone} = useMutation(
		DELETE_CHAT,
		() => ({
			variables: {
				id: id.value,
			},
		}),
		{
			fetchPolicy: 'no-cache'
		}
	)

	onDone((result) => {
		console.log(result.data)
		emits("removed")
	})


	function remove() {
		id.value = props.chat.id;
		mutate();
	}
</script>

<template>
	<li class="w-full flex flex-row justify-between items-center gap-1 overflow-hidden px-2 max-h-10">
		<router-link :to="`/chat/${props.chat.id}`" class="grow max-w-[calc(100%-3rem)]"
			:class="{'outline': props.selected, 'outline-primary': props.selected}">
			<p class="truncate whitespace-nowrap overflow-hidden text-ellipsis">{{ props.chat.title }}</p>
		</router-link>
		<button class="btn btn-circle" @click="remove" :disabled="loading">
			<i class="bi bi-trash text-error"></i>
		</button>
	</li>
</template>
