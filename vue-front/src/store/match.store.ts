import { defineStore } from "pinia";
import { ref } from "vue";
import { Match, fetchMatchesByLeagueAndDate } from "../api/api";

export const useMatchStore = defineStore("match", () => {
	const matchs = ref<Match[]>([]);

	async function fetchMatchs(league: string) {
		const res = await fetchMatchesByLeagueAndDate(league, new Date());
		matchs.value = res;
	}

	return { matchs, fetchMatchs };
});
