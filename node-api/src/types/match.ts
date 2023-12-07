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
	winProbability: number;
}

export interface Stadium {
	name: string;
	city: string;
}
