<script setup lang="ts">
import { onMounted } from "vue";
import { useMatchStore } from "../store/match.store";
import Match from "../components/MatchComponent.vue";
import LeagueSelector from "../components/LeagueSelectorComponent.vue";
import { useLeagueStore } from "../store/leagues.store";
import CoinsComponent from "../components/CoinsComponent.vue";
const matchStore = useMatchStore();
const leagueStore = useLeagueStore();

onMounted(async () => {
	await matchStore.fetchMatchs(leagueStore.selectedLeague);
	await matchStore.monitorMatchs(
		matchStore.matchs
			.filter((match) => match.status === "ongoing")
			.map((match) => match.id),
	);
});
</script>

<template>
	<div class="container">
		<div class="sidebar">
			<CoinsComponent></CoinsComponent>
			<LeagueSelector />
		</div>
		<div class="matchs">
			<h3 class="title">Matchs</h3>
			<div v-if="matchStore.matchs.length > 0">
				<Match
					v-for="match in matchStore.matchs"
					:key="match.id"
					:match="match"
				></Match>
			</div>
			<div v-else>Aucun match prévu aujourd'hui 😢</div>
		</div>
	</div>
</template>

<style scoped lang="scss">
.container {
	display: flex;
	background-color: #eee;
	height: 100%;
	padding: 16px;
	box-sizing: border-box;
}

.sidebar {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	background-color: #ddd;
	box-sizing: border-box;
	border-radius: 1rem;
	padding: 0.5rem;
	border: 1px solid white;
	height: 100%;
}

.matchs {
	background-color: white;
	margin-left: 2rem;
	width: 100%;
	border-radius: 1rem;
	border: 1px solid #bbb;
	padding: 4rem;
}

.title {
	font-weight: bold;
	font-size: 2rem;
	margin: 0;
}
</style>
