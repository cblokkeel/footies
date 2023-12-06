import express from "express";
import { config } from "./config";
import cors from "cors";
import { matchsRouter } from "./controllers/matchs.controller";

const app = express();
app.use(express.json());
app.use(cors()); // TODO: secure cors
app.use(matchsRouter);

app.listen(config.API_PORT, () => {
	console.log(`API running on port ${config.API_PORT}`);
});
