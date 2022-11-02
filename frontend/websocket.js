import { Player, store } from "./app";
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
        window.location.href = window.location.origin + "/#/game";

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
