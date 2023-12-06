import { match } from "assert";
import { Router } from "express";
import { getMatchs } from "~/services/matchs.service";

export const matchsRouter = Router();

matchsRouter.get("/api/matchs/:league", async (req, res) => {
	const league = req.params["league"];
	const date = req.query["date"] as string;

	if (!league || !date) {
		return res.status(400).json({
			error: "invalid parameters (league or date)", // maybe be more specific about whats missing
		});
	}

	const matchs = await getMatchs({ league, date });
	return res.json(matchs);
});
