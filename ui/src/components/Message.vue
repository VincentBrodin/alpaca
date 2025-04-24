<script setup>
	import {marked} from 'marked';
	import {computed} from 'vue';

	const props = defineProps({
		message: Object,
	});

	const html = computed(() => marked.parse(props.message.content));
</script>

<template>
	<div class="w-full flex"
		:class="{'justify-end': message.role === 'user', 'justify-start': message.role !== 'user'}">
		<div v-if="message.role === 'user'" class="px-4 py-2 rounded-xl bg-base-200" style="max-width: 50%;">
			{{ message.content }}
		</div>
		<div v-else v-html="html" class="px-4 py-2 prose max-w-full min-w-0 w-full">
		</div>
	</div>
</template>
