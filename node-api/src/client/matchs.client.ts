import { Match } from "~/types/match";
import axios from "axios";

export async function fetchMatchsByLeagueAndDate(
	league: string,
	date: string,
	season: string,
): Promise<Match[]> {
	try {
		const res = await axios<Match[]>(
			`${process.env.WORKER_URL}/matchs?league=${league}&date=${date}&season=${season}`,
		);
		return res.data;
	} catch (err) {
		return []; // not the best error handling, to refacto
	}
}
