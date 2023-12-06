import { defineStore } from "pinia";
import { ref } from "vue";
import { Match, fetchMatchesByLeagueAndDate } from "../api/api";

export interface League {
	id: string;
	name: string;
	logo: string;
}

export const useLeagueStore = defineStore("league", () => {
	const allLeagues = ref<League[]>([]);
	const selectedLeague = ref<string>("");

	function loadLeagues() {
		// Should fetch from api or smthing, maybe do later
		allLeagues.value = [
			{
				id: "39",
				name: "Premier league",
				logo: "https://media-4.api-sports.io/football/leagues/39.png",
			},
			{
				id: "61",
				name: "Ligue 1",
				logo: "https://media-4.api-sports.io/football/leagues/61.png",
			},
			{
				id: "78",
				name: "Bundesliga",
				logo: "https://media-4.api-sports.io/football/leagues/78.png",
			},
			{
				id: "135",
				name: "Serie A",
				logo: "https://media-4.api-sports.io/football/leagues/135.png",
			},
			{
				id: "140",
				name: "La Liga",
				logo: "https://media-4.api-sports.io/football/leagues/140.png",
			},
		];
		selectedLeague.value = "39";
	}

	function selectLeague(id: string) {
		selectedLeague.value = id;
	}

	return { allLeagues, selectedLeague, loadLeagues, selectLeague };
});
