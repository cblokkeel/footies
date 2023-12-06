import { Match } from "~/types/match";
import axios from "axios";

const API_BASE_URL = "http://localhost:8888";

export async function fetchMatchsByLeagueAndDate(
	league: string,
	date: string,
	season: string,
): Promise<Match[]> {
	try {
		const res = await axios<Match[]>(
			`${API_BASE_URL}/matchs?league=${league}&date=${date}&season=${season}`,
		);
		return res.data;
	} catch (err) {
		return []; // not the best error handling, to refacto
	}
}
