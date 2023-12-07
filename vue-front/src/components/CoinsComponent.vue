<script setup lang="ts">
import { useBetStore } from "../store/bet.store";
import coinsSrc from "../assets/coins.svg";

const betStore = useBetStore();
</script>

<template>
	<p class="coins">
		{{ betStore.coins }}
		<img class="icon-coin" :src="coinsSrc" alt="coin-icon" />
	</p>
</template>

<style scoped lang="scss">
$yellow: #fff200;

.coins {
	display: flex;
	align-items: center;
	gap: 10px;
	color: $yellow;
	position: absolute;
	top: 15px;
	right: 15px;
}

.icon-coin {
	width: 16px;
	height: 16px;
	color: $yellow;
}
</style>
