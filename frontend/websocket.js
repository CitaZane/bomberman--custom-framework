import { store } from "./app";
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
    console.log("DATA", data)
    switch (data["type"]) {
      case "NEW_USER":
      case "USER_LEFT":
        store.commit("updateUserQueueCount", data.gameState.players.length);
        break;

      case "INIT_ROOM":
        store.commit("updateUserQueueCount", data.gameState.players.length);
        window.location.href = window.location.origin + "/#/game";
        break;

      case "TEXT_MESSAGE":
        console.log("Got text message", data);
        store.dispatch("addNewMessage", data);
    }
  };
}
