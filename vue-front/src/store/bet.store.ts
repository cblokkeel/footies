import { defineStore } from "pinia";
import { ref } from "vue";

export const useBetStore = defineStore(
	"bet",
	() => {
		const bets = ref<Record<string, { bet: number; on: "home" | "away" }>>(
			{},
		);
		const coins = ref<number>(10000);

		function betOnMatch(
			matchId: string,
			bet: number,
			betOn: "home" | "away",
		): boolean {
			if (coins.value < bet) {
				return false;
			}
			bets.value[matchId] = { bet, on: betOn };
			coins.value -= bet;
			return true;
		}

		function matchOver(matchId: string, winner: "home" | "away" | "nil") {
			const bet = bets.value[matchId];
			if (!bet) {
				return;
			}
			if (bet.on === winner) {
				coins.value += bet.bet * 2;
			}
			removeBet(matchId);
		}

		function getBetByMatchId(matchId: string): {
			on: "home" | "away";
			bet: number;
		} | null {
			const bet = bets.value[matchId];
			if (!bet) {
				return null;
			}
			return bet;
		}

		function removeBet(id: string) {
			const filteredBets: Record<
				string,
				{ bet: number; on: "home" | "away" }
			> = {};
			for (let matchId in bets.value) {
				if (matchId !== id) {
					filteredBets[matchId] = bets.value[matchId];
				}
			}
			bets.value = filteredBets;
		}

		return { bets, coins, betOnMatch, getBetByMatchId, matchOver };
	},
	{
		persist: true,
	},
);
