import express from "express";
import { config } from "./config";
import cors from "cors";

const app = express();
app.use(express.json());
app.use(cors()); // TODO: secure cors

app.get("/", (req, res) => {
	res.send("Hello world");
});

app.listen(config.API_PORT, () => {
	console.log(`API running on port ${config.API_PORT}`);
});
