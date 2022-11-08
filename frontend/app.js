import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { ws } from "./websocket";
import { advThrottle } from "./throttle";

import { defineWebSocket } from "./websocket";
const store = createStore(storeObj);
const router = createRouter(routes);
let gameStarted = false;

// add keyup and KeyDown listeners
export function setupGame() {
  document.addEventListener("keyup", (e) => {
    store.dispatch("registerKeyUp", e.code);
  });
  document.addEventListener("keydown", (e) => {
    store.dispatch("registerKeyDown", e.code);
  });
  if (!gameStarted) {
    gameStarted = true;
    animate();
  }
}

let throttleDropBomb = advThrottle(dropBomb, 3000, { leading: true });

function dropBomb(currentPlayerName) {
  ws.send(
    JSON.stringify({
      type: "PLAYER_DROPPED_BOMB",
      creator: currentPlayerName,
    })
  );
}
// name provided by backend
defineWebSocket("user");

// main game loop
let gameFrame = 0;
function animate() {
  let movement = store.state.movement;
  let currentPlayerName = store.state.currentPlayerName;
  let inputs = store.state.inputs;
  // sends all 4 movements
  if (movement.move) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: movement.move,
      })
    );
  } else if (movement.stop) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: movement.stop,
      })
    );
    //send movement stop only once, so clear the variable after sending
    store.dispatch("clearStopMovement");
  }

  if (inputs?.Space) {
    // console.log("here");
    throttleDropBomb(currentPlayerName);
    // console.log("dropping bomb!");
  }

  gameFrame++;
  requestAnimationFrame(animate);
}

export { store, router, gameFrame };
