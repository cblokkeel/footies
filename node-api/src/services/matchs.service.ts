import { fetchMatchsByLeagueAndDate } from "~/client/matchs.client";
import { Match } from "~/types/match";

export interface GetMatchsParams {
	league: string;
	date: string;
}

export async function getMatchs({
	league,
	date,
}: GetMatchsParams): Promise<Match[]> {
	const year = parseInt(date.split("-")[0]);
	const season = `${year < 6 ? year - 1 : year}`;
	return fetchMatchsByLeagueAndDate(league, date, season);
}
