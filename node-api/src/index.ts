import express from "express";
import WebSocket from "ws";
import { config } from "./config";
import cors from "cors";
import { matchsRouter } from "./controllers/matchs.controller";
import { createServer } from "http";
import { createClient } from "redis";
import dotenv from "dotenv";

try {
	dotenv.config();
} catch (err) {
	console.log("No env file found");
}

(async () => {
	const app = express();
	const server = createServer(app);
	app.use(express.json());
	app.use(cors()); // TODO: secure cors
	app.use(matchsRouter);

	const redis = createClient({
		url: process.env.REDIS_ADDR,
	});
	await redis.connect();
	const sub = redis.duplicate();
	await sub.connect();

	const wss = new WebSocket.Server({ server });

	wss.on("connection", function (ws) {
		console.log("New connection");
		ws.on("message", function (msg) {
			msg.toString()
				.split(",")
				.forEach((matchId) => {
					sub.subscribe(`match_${matchId}_update`, (msg) => {
						ws.send(`${matchId}_${msg}`);
					});
				});
			redis.publish("monitoring", msg.toString());
		});
	});

	server.listen(config.API_PORT, () => {
		console.log(`API running on port ${config.API_PORT}`);
	});
})();
