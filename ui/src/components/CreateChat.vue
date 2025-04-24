<script setup>
	import {CREATE_CHAT} from '@/queries';
	import {useMutation} from '@vue/apollo-composable'

	const emits = defineEmits(["created"])

	const {mutate, loading, error, onDone} = useMutation(
		CREATE_CHAT,
		{},
		{
			fetchPolicy: 'no-cache'
		}
	)


	function create() {
		mutate();
	}

	onDone((result) => {
		console.log(result.data)
		emits("created")
	})
</script>

<template>
	<li class="w-full">
		<a @click.prevent="create" class="w-full flex flex-row justify-center items-center" :disabled="loading">
			<i class="bi bi-plus-lg"></i>
		</a>
	</li>
</template>
