<script setup lang="ts">
import { Match } from "../api/api";
import BetsComponent from "./BetsComponent.vue";
import Team from "./Team.vue";

defineProps<{ match: Match }>();
</script>

<template>
	<div class="match">
		<span v-if="match.status === 'ongoing'" class="match__time"
			>{{ match.elapsed }}'</span
		>
		<span v-if="match.status === 'finished'">Terminé</span>
		<span v-if="match.status === 'not_started'">Match à venir</span>
		<span v-if="match.status === 'half_time'">Mi temps</span>
		<div class="match__team-container">
			<Team :team="match.homeTeam" :away="false" />
			-
			<Team :team="match.awayTeam" :away="true" />
		</div>
		<BetsComponent :match="match"></BetsComponent>
	</div>
</template>

<style scoped lang="scss">
.match {
	padding: 10px;

	&__time {
		color: green;
	}

	&__team-container {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
}
</style>
