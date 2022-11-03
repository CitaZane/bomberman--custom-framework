import { store } from "./app";
import { setupGame, gameFrame } from "./app";
export let ws;
// export let players = [];
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
          store.dispatch('registerPlayer', {playerName,i})
        })
        store.dispatch('registerCurrentPlayer', store.state.currentPlayerName)

        window.location.href = window.location.origin + "/#/game";
        setupGame();
        break

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
        break;
        // game  stuff
      case "PLAYER_MOVE":
        let index= Number(data.creator) //index of player sending the movement
        if(data.body == 'LEFT'){
          store.dispatch('movePlayerLeft', {index, gameFrame})
        }else if (data.body == 'RIGHT'){
          store.dispatch('movePlayerRight', {index, gameFrame})
        }else if (data.body == 'UP'){
          store.dispatch('movePlayerUp', {index, gameFrame})
        }else if (data.body == 'DOWN'){
          store.dispatch('movePlayerDown', {index, gameFrame})
        }
    }
  };
}
