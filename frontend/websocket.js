import { Player, store } from "./app";
import { setupGame } from "./app";
export let ws;
export let players = [];
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

        data.gameState.players.forEach((playerName, i) => {
          const player = new Player(playerName);
          player.y = i * 50;
          players.push(new Player(playerName))
        })

        window.location.href = window.location.origin + "/#/game";

        setupGame();
        break

      // console.log("Players", players)
      // queue cases  
      case "NEW_USER":
      case "USER_LEFT":
        store.commit("updateUserQueueCount", data.gameState.players.length);
        break;

      case "JOIN_QUEUE":
        store.commit("updateUserQueueCount", data.gameState.players.length);
        window.location.href = window.location.origin + "/#/queue";
        break;

      case "TEXT_MESSAGE":
        store.dispatch("addNewMessage", data);
    }
  };
}
