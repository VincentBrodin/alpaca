<script setup>
	import {computed, ref, watch} from 'vue';

	const props = defineProps({
		generating: Boolean,
	});

	const emit = defineEmits(["send"]);
	const content = ref("")

	const canSend = computed(() => content.value !== "")

	function submit() {
		if (content.value === "") return
		emit("send", {content: content.value});
		content.value = "";
	}

	watch(() => props.generating, val => console.log('generating now:', val));
</script>


<template>
	<form @submit.prevent="submit" class="absolute bottom-4 left-0 right-0 px-4 flex justify-center items-center">
		<div class="max-w-4xl grow flex justify-center items-center gap-4">
			<textarea rows="1" class="grow px-auto textarea textarea-md resize-none h-auto" v-model="content"
				:disabled="props.generating"></textarea>
			<button type="submit" class="btn btn-circle btn-lg min-h-full" :disabled="!canSend">
				<i class="bi bi-send"></i>
			</button>
		</div>
	</form>
</template>
