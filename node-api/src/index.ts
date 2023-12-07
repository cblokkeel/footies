import express from "express";
import WebSocket from "ws";
import { config } from "./config";
import cors from "cors";
import { matchsRouter } from "./controllers/matchs.controller";
import { createServer } from "http";
import { createClient } from "redis";

(async () => {
	const app = express();
	const server = createServer(app);
	app.use(express.json());
	app.use(cors()); // TODO: secure cors
	app.use(matchsRouter);

	const redis = createClient();
	await redis.connect();

	const wss = new WebSocket.Server({ server });

	wss.on("connection", function (ws) {
		console.log("New connection");
		ws.on("message", function (msg) {
			console.log(msg.toString());
			redis.publish("monitoring", msg.toString());
			msg.toString()
				.split(",")
				.forEach((matchId) => {
					redis.subscribe(`match_${matchId}_update`, (msg) => {
						ws.send(`${matchId}_${msg}`);
					});
				});
		});
	});

	server.listen(config.API_PORT, () => {
		console.log(`API running on port ${config.API_PORT}`);
	});
})();
