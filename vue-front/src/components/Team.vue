<script setup lang="ts">
import { Team } from "../api/api";

defineProps<{ team: Team; away: boolean }>();
</script>

<template>
	<div class="team" :class="{ 'team--away': away }">
		<div class="team__informations">
			<img class="team__logo" :src="team.logo" />
			<span>{{ team.name }}</span>
		</div>
		<span>{{ team.score }}</span>
	</div>
</template>

<style scoped lang="scss">
.team {
	display: flex;
	align-items: center;
	justify-content: space-between;
	width: 100%;

	&__informations {
		display: flex;
		align-items: center;
		gap: 15px;
		margin: 0 50px;
	}

	&__logo {
		height: 48px;
		width: 48px;
	}

	&--away {
		flex-direction: row-reverse;

		.team__informations {
			flex-direction: row-reverse;
		}
	}
}
</style>
