<script setup lang="ts">
import { League, useLeagueStore } from "../store/leagues.store";
import { useMatchStore } from "../store/match.store";

const leagueStore = useLeagueStore();
const matchStore = useMatchStore();

async function selectLeague(id: string) {
	leagueStore.selectLeague(id);
	await matchStore.fetchMatchs(id);
}
defineProps<{ league: League }>();
</script>

<template>
	<div
		class="league"
		:class="{ 'league--active': league.id === leagueStore.selectedLeague }"
		@click="selectLeague(league.id)"
	>
		<img :src="league.logo" :alt="league.name" class="league__logo" />
	</div>
</template>

<style scoped lang="scss">
.league {
	background-color: #191919;
	border-radius: 15px;
	padding: 15px;
	width: fit-content;
	cursor: pointer;
	transition: 0.2 ease-in-out;

	&:hover {
		background-color: #202020;
		transform: scale(1.1);
	}

	&--active {
		background-color: green;
	}

	&__logo {
		width: 48px;
		height: 48px;
	}
}
</style>
