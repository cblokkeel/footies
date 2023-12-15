<script setup lang="ts">
import { Match } from "../api/api";
import BetsComponent from "./BetsComponent.vue";
import Team from "./Team.vue";

defineProps<{ match: Match }>();
</script>

<template>
	<div class="match">
		<div class="match__status">
			<span v-if="match.status === 'ongoing'" class="match__time"
				>{{ match.elapsed }}'</span
			>
			<div v-if="match.status === 'ongoing'" class="status-bar">
				<div class="status-indicator"></div>
			</div>
			<span v-if="match.status === 'finished'">Terminé</span>
			<span v-if="match.status === 'not_started'">Match à venir</span>
			<span v-if="match.status === 'half_time'">Mi temps</span>
		</div>
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
		color: #3ae374;
	}

	&__team-container {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
	}

	&__status {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-align: center;
	}
}

.status-bar {
	position: relative;
	background-color: #3ae374;
	height: 0.15rem;
	width: 1.6rem;
	overflow: hidden;

	.status-indicator {
		position: absolute;
		top: 0;
		left: -100%;
		height: 100%;
		width: 100%;
		background-color: #09fd5f;
		animation: move 1.5s linear infinite;
	}
}

@keyframes move {
	to {
		left: 100%;
	}
}
</style>
