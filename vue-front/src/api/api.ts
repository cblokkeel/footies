import axios from "axios";

export interface Match {
	id: string;
	elapsed: number;
	status: string;
	homeTeam: Team;
	awayTeam: Team;
	stadium: Stadium;
}

export interface Team {
	name: string;
	logo: string;
	score: number;
}

export interface Stadium {
	name: string;
	city: string;
}

// const API_BASE_URL = process.env.API_BASE_URL;
const API_BASE_URL = "http://localhost:3000";

export async function fetchMatchesByLeagueAndDate(
	league: string,
	date: Date,
): Promise<Match[]> {
	console.log(league);
	const isoDate = date.toISOString().split("T")[0];
	const res = await axios<Match[]>(
		`${API_BASE_URL}/api/matchs/${league}?date=${isoDate}`,
	);
	return res.data;
}
