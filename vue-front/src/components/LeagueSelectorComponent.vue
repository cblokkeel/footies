<script setup lang="ts">
import { onMounted } from "vue";
import { useLeagueStore } from "../store/leagues.store";
import LeagueComponent from "./LeagueComponent.vue";

const leagueStore = useLeagueStore();
onMounted(() => {
	leagueStore.loadLeagues();
});
</script>

<template>
	<div class="league-selector">
		<div v-for="league in leagueStore.allLeagues">
			<LeagueComponent :league="league" />
		</div>
	</div>
</template>

<style scoped lang="scss">
.league-selector {
	display: flex;
	flex-direction: column;
	justify-content: center;
	gap: 1rem;
	margin-bottom: auto;
}
</style>
