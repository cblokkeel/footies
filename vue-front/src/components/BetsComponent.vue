<script setup lang="ts">
import ConfettiExplosion from "vue-confetti-explosion";
import { computed, onMounted, ref, watch } from "vue";
import { Match } from "../api/api";
import { useBetStore } from "../store/bet.store";
import { useMatchStore } from "../store/match.store";

const props = defineProps<{ match: Match }>();
const betStore = useBetStore();
const matchStore = useMatchStore();

const isMatchFinished = computed(() => matchStore.isFinished(props.match.id));

const homeBetEl = ref<HTMLInputElement>();
const awayBetEl = ref<HTMLInputElement>();

onMounted(() => {
	betStore.hasWon = false;
	const bet = betStore.getBetByMatchId(props.match.id);
	if (bet && homeBetEl && homeBetEl.value && awayBetEl && awayBetEl.value) {
		if (bet.on === "home") {
			homeBetEl.value.value = `${bet.bet}`;
		} else {
			awayBetEl.value.value = `${bet.bet}`;
		}
		homeBetEl.value.disabled = true;
		awayBetEl.value.disabled = true;
	}
});

watch(isMatchFinished, () => {
	let winner: "home" | "away" | "nil" = "nil";
	let multiplicater: number = 1;
	if (props.match.homeTeam.score > props.match.awayTeam.score) {
		winner = "home";
		multiplicater = props.match.homeTeam.winProbability;
	} else if (props.match.awayTeam.score > props.match.homeTeam.score) {
		winner = "away";
		multiplicater = props.match.awayTeam.winProbability;
	}
	betStore.matchOver(props.match.id, winner, multiplicater);
});

function handleBet(on: "home" | "away") {
	if (homeBetEl && homeBetEl.value && awayBetEl && awayBetEl.value) {
		let bet = on === "home" ? homeBetEl.value.value : awayBetEl.value.value;
		const hasBet = betStore.betOnMatch(props.match.id, parseInt(bet), on);
		if (!hasBet) {
			homeBetEl.value.value = "";
			awayBetEl.value.value = "";
			return;
		}
		homeBetEl.value.disabled = true;
		awayBetEl.value.disabled = true;
	}
}
</script>

<template>
	<div class="bets__container">
		<div class="confetti" v-if="betStore.hasWon">
			<ConfettiExplosion />
		</div>
		<div class="bets__element">
			<p>{{ match.homeTeam.winProbability }}</p>
			<input
				@keydown.enter="handleBet('home')"
				class="bet"
				placeholder="0.0"
				type="number"
				ref="homeBetEl"
				:disabled="match.status === 'finished'"
			/>
		</div>
		<div class="bets__element">
			<p>{{ match.awayTeam.winProbability }}</p>
			<input
				@keydown.enter="handleBet('away')"
				class="bet"
				placeholder="0.0"
				type="number"
				ref="awayBetEl"
				:disabled="match.status === 'finished'"
			/>
		</div>
	</div>
</template>

<style scoped lang="scss">
.bets {
	&__container {
		display: flex;
		align-items: center;
		justify-content: space-around;
		margin-top: 15px;
	}

	&__element {
		display: flex;
		align-items: center;
		gap: 15px;
	}
}

.bet {
	height: 32px;
	text-align: center;
}

.confetti {
	position: absolute;
	width: 50%;
	height: 50%;
	bottom: 0;
	right: 0;
	margin: auto;
}
</style>
