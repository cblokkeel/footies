<script setup lang="ts">
import { useBetStore } from "../store/bet.store";

const betStore = useBetStore();
</script>

<template>
	<div class="coins-container">
		<p class="coins">
			{{ betStore.coins }}
		</p>
	</div>
</template>

<style scoped lang="scss">
.coins-container {
	display: flex;
	align-items: center;
	justify-content: center;
	background-color: white;
	border: 1px solid #bbb;
	border-radius: 1rem;
	width: 8rem;
	height: 4rem;
	margin-bottom: auto;
}
.coins {
	display: flex;
	align-items: center;
	font-weight: bold;
	font-size: 1.5rem;
	top: 1rem;
	right: 1rem;
}

.icon-coin {
	width: 1rem;
	height: 1rem;
}
</style>
