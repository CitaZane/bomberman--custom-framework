import { store } from "./app";
import { setupGame, gameFrame } from "./app";
export let ws;
export function defineWebSocket(name) {
  ws = new WebSocket(`ws://localhost:8080/ws?username=${name}`);

  ws.onopen = () => {
    console.log("Connection initiated");
  };

  ws.onclose = () => {
    console.log("Connection closed");
  };

  ws.onmessage = (e) => {
    const data = JSON.parse(e.data);
    // console.log("DATA", data)
    switch (data["type"]) {
      case "START_GAME":
        if (data.body === "me") {
          store.commit("changePlayerName", data.creator);
        }
        store.commit("updatePlayers", []);

        data.gameState.players.forEach((player) => {
          store.dispatch("registerPlayer", player);
        });
        store.commit('updateMap', data.gameState.map)
        // console.log("Game started with: ", data.gameState.map)
        // window.location.href = window.location.origin + "/#/game";
        setupGame();
        break;

      // queue cases
      case "NEW_USER":
      case "USER_LEFT":
        store.commit("updateUserQueueCount", data.body);
        break;

      case "JOIN_QUEUE":
        // store.commit("updateUserQueueCount", data.body);
        // window.location.href = window.location.origin + "/#/queue";

        // create new game and update players
        store.commit("updatePlayers", []);

        data.gameState.players.forEach((player) => {
          store.dispatch("registerPlayer", player);
        });
        store.commit('updateMap', data.gameState.map)
        setupGame();
        // ---------------
        break;

      case "TEXT_MESSAGE":
        store.dispatch("addNewMessage", data);
        break;
      // game  stuff
      case "PLAYER_MOVE":
        store.commit("updatePlayers", data.gameState.players);
    }
  };
}
