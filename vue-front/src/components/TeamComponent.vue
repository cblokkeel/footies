<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { Team } from "../api/api";

const newGoal = ref<boolean>(false);
const props = defineProps<{ team: Team; away: boolean }>();
const hasScored = computed(() => props.team.score);

watch(hasScored, () => {
	newGoal.value = true;
	setTimeout(() => {
		newGoal.value = false;
	}, 1300);
});
</script>

<template>
	<div class="team" :class="{ 'team--away': away }">
		<div class="team__informations">
			<img class="team__logo" :src="team.logo" />
			<span>{{ team.name }}</span>
		</div>
		<div class="team__score-container">
			<span>{{ team.score }}</span>
			<span v-if="newGoal" class="team__new-goal">+1</span>
		</div>
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
		margin: 0 3rem;
	}

	&__logo {
		height: 3rem;
		width: 3rem;
	}

	&__score-container {
		position: relative;
	}

	&__new-goal {
		color: #009a44;
		position: absolute;
		top: -1rem;
		left: 0.5rem;
		animation: fadeInOut 1.5s ease-in-out;
	}

	&--away {
		flex-direction: row-reverse;

		.team__informations {
			flex-direction: row-reverse;
		}
	}
}

@keyframes fadeInOut {
	0% {
		opacity: 0;
		transform: translateY(1.4rem);
	}

	50% {
		opacity: 1;
		transform: translateY(0);
	}

	100% {
		opacity: 0;
		transform: translateY(-1.4rem);
	}
}
</style>
