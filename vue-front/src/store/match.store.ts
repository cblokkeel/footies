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
			const [updatedMatchId, updateType, updateValue] = (
				msg.data as string
			).split("_");
			matchs.value
				.filter((m) => m.id === updatedMatchId)
				.map((m) => {
					if (updateType === "chrono") {
						m.elapsed = parseInt(updateValue);
					}
					if (updateType === "homegoal") {
						m.homeTeam.score = parseInt(updateValue);
					}
					if (updateType === "awaygoal") {
						m.awayTeam.score = parseInt(updateValue);
					}
					if (updateType === "status") {
						m.status = updateValue;
					}
				});
			console.log(updatedMatchId, updateType, updateValue);
		};
	}

	return { matchs, fetchMatchs, monitorMatchs };
});
