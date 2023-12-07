import { RouteRecordRaw } from "vue-router";
import AppVue from "./pages/MatchsPage.vue";
import ConfigVue from "./pages/ConfigPage.vue";

export const routes: RouteRecordRaw[] = [
	{
		path: "/",
		component: AppVue,
	},
	{
		path: "/config",
		component: ConfigVue,
	},
];
