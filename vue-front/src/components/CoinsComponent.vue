<script setup lang="ts">
import { useBetStore } from "../store/bet.store";

const betStore = useBetStore();
</script>

<template>
	<p class="coins">
		{{ betStore.coins }}
	</p>
</template>

<style scoped lang="scss">
.coins {
	position: absolute;
	top: 15px;
	right: 15px;
}
</style>
