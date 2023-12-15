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
	display: flex;
	justify-content: center;
	align-items: center;
	background-color: white;
	border: 1px solid #bbb;
	width: 4rem;
	height: 4rem;
	border-radius: 1rem;
	cursor: pointer;
	transition: 0.2 ease-in-out;

	&:hover {
		transform: scale(1.05);
	}

	&--active {
		border: 1px solid #009a44;
	}

	&__logo {
		width: 3rem;
		height: 3rem;
	}
}
</style>
