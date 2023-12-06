import { defineStore } from "pinia";
import { ref } from "vue";
import { Match, fetchMatchesByLeagueAndDate } from "../api/api";

export const useMatchStore = defineStore("match", () => {
	const matchs = ref<Match[]>([]);

	async function fetchMatchs(league: string) {
		const res = await fetchMatchesByLeagueAndDate(league, new Date());
		matchs.value = res;
	}

	async function monitorMatchs(ids: string[]) {
		const connexion = new WebSocket("ws://localhost:3000");
		connexion.onopen = function (_) {
			console.log("Connected");
			connexion.send(ids.join(","));
		};

		connexion.onmessage = function (msg) {
			console.log(msg);
		};
	}

	return { matchs, fetchMatchs, monitorMatchs };
});
