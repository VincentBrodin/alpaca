<script setup>
	import {ref} from 'vue';


	const emit = defineEmits(["send"]);
	const content = ref("")
	function submit() {
		emit("send", {content: content.value});
		content.value = "";
	}
</script>


<template>
	<form @submit.prevent="submit" class="absolute bottom-8 left-0 right-0 px-4 flex justify-center items-center">
		<input type="text" v-model="content" placeholder="Chat here" class="max-w-4xl w-full px-auto input input-bordered" />
	</form>
</template>
