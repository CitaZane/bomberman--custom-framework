import { store } from "./app";
import { setupGame } from "./app";
export let ws;
export function defineWebSocket(name) {
  ws = new WebSocket(`ws://localhost:8080/ws?username=${name}`);

  ws.onopen = () => {
    console.log("Connection initiated");
  };

  ws.onclose = (e) => {
    console.log("Connection closed");
    if (e.code == 1008) {  // 1008 is policy violation
       alert(e.reason)
    }
  };


  ws.onmessage = (e) => {
    const data = JSON.parse(e.data);
    // console.log("DATA", data)
    switch (data["type"]) {
      case "START_GAME":
        store.commit("updatePlayers", data.gameState.players);
        store.commit("updateMap", data.gameState.map);
        window.location.href = window.location.origin + "/#/game";
        setupGame();
        break;

      // queue cases
      case "NEW_USER":
        case "USER_LEFT":
          store.commit("updateUserQueueCount", data.body);
          break;
          
          case "JOIN_QUEUE":
            store.commit("updateUserQueueCount", data.body);
            window.location.href = window.location.origin + "/#/queue";
            break;
        case "JOIN_SPECTATOR":
          store.commit("updatePlayers", data.gameState.players);
          store.commit("updateMap", data.gameState.map);
          console.log("Game is already in action. And you are ", data.body, "in list")
           window.location.href = window.location.origin + "/#/game";
            break;
      // game  stuff
      case "PLAYER_MOVE":
        if (data.body === "PICKED_UP_POWERUP") {
          store.commit("updatePowerUps", data.gameState["power_ups"]);
        }
        store.commit("updatePlayers", data.gameState.players);
        break;

      case "TEXT_MESSAGE":
        store.dispatch("addNewMessage", data);
        break;
      case "MAP_UPDATE":
        store.commit("updateMap", data.gameState.map);
        store.commit("updatePowerUps", data.gameState["power_ups"]);
        break;
      case "PLAYER_DROPPED_BOMB":
      case "BOMB_EXPLODED":
      case "EXPLOSION_COMPLETED":
      case "PLAYER_LEFT":
      case "PLAYER_REBORN":
        store.commit("updatePlayers", data.gameState.players);
        break;
      case "FINISH":
        console.log("Player: ", store.state.currentPlayerName)
        console.log("Players: ", data.gameState.players)
        var isPlayer = data.gameState.players.some(player => player.name == store.state.currentPlayerName)
        if (isPlayer){
          console.log("Is player!")
          window.location.href = window.location.origin + "/";
        }else{
          console.log("Watcher")
          window.location.href = window.location.origin + "/#/queue"
        }
        break;

    }
  };
}

export function SendWsMessage(type, creator, body, delta) {
  ws.send(
    JSON.stringify({
      type: type,
      creator: creator,
      body: body,
      delta:delta
    })
  );
}
