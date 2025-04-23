<script setup>
	import {marked} from 'marked';

	const props = defineProps({
		message: Object,
	});

	const html = marked.parse(props.message.content);
</script>

<template>
	<div class="w-full flex"
		:class="{'justify-end': message.role === 'user', 'justify-start': message.role !== 'user'}">
		<div v-if="message.role === 'user'" class="px-4 py-2 rounded-full bg-base-200">
			{{ message.content }}
		</div>
		<div v-else v-html="html" class="px-4 prose">
		</div>
	</div>
</template>
