import cors from "cors";
import express from "express";
import { config } from "./config";
import { mockedMatchs } from "./mock";

const app = express();
app.use(cors());

app.get("/", (req, res) => {
	const league = req.query["league"];
	if (league) {
		const matchs = findByLeague(league as string);
		if (matchs.length === 0) {
			return res.sendStatus(404);
		}
		return res.json({
			response: matchs,
		});
	}
	const matchId = req.query["id"];
	if (matchId) {
		const matchs = findById(matchId as string);
		if (matchs.length === 0) {
			return res.sendStatus(404);
		}
		return res.json({
			response: matchs,
		});
	}
	return res.sendStatus(400);
});

function findByLeague(id: string) {
	return mockedMatchs.filter((m) => m.league.id === parseInt(id));
}

function findById(id: string) {
	return mockedMatchs.filter((m) => m.fixture.id === id);
}

app.listen(config.API_PORT, () => {
	console.log(`Listening on port ${config.API_PORT}`);
});
