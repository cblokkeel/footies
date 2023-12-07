<script setup lang="ts">
import { onMounted } from "vue";
import { useMatchStore } from "../store/match.store";
import Match from "../components/Match.vue";
import LeagueSelector from "../components/LeagueSelector.vue";
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
	<CoinsComponent></CoinsComponent>
	<div class="container">
		<LeagueSelector></LeagueSelector>
		<div class="matchs" v-if="matchStore.matchs.length > 0">
			<Match
				v-for="match in matchStore.matchs"
				:key="match.id"
				:match="match"
			></Match>
		</div>
		<div v-else>Aucun match prévu aujourd'hui 😢</div>
	</div>
</template>

<style scoped lang="scss">
.container {
	display: flex;
	flex-direction: column;
}

.matchs {
	display: flex;
	flex-direction: column;
	gap: 30px;
}
</style>
