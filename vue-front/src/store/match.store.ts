import { defineStore } from "pinia";
import { computed, ref } from "vue";
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
					switch (updateType) {
						case "chrono":
							m.elapsed = parseInt(updateValue);
							break;
						case "homegoal":
							m.homeTeam.score = parseInt(updateValue);
							break;
						case "awaygoal":
							m.awayTeam.score = parseInt(updateValue);
							break;
						case "status":
							m.status = updateValue;
							break;
						default:
							console.log(
								`unexpected update type: ${updateType}`,
							);
					}
				});
		};
	}

	function isFinished(id: string) {
		const match = matchs.value.filter((match) => match.id === id).pop();
		if (match && match.status === "finished") {
			return true;
		}
		return false;
	}

	return { matchs, fetchMatchs, monitorMatchs, isFinished };
});
