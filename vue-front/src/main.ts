import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { createPinia } from "pinia";
import { routes } from "./routes";
import {
	createRouter,
	createWebHashHistory,
	createWebHistory,
} from "vue-router";

const pinia = createPinia();

const app = createApp(App);
const router = createRouter({
	history: createWebHistory(),
	routes: routes,
});

app.use(pinia);
app.use(router);
app.mount("#app");
