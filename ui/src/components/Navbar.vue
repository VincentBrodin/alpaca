<script setup>
	import {useLocalStorage} from '@vueuse/core';
	const drawerOpen = useLocalStorage('drawer-open', true);

</script>

<template>

	<button v-if="!drawerOpen" class="btn btn-circle" @click="drawerOpen = true"><i
			class="bi bi-arrow-bar-right text-sm"></i></button>

</template>
