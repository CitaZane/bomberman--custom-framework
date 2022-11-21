import createRouter from "./framework/router";
import createStore from "./framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { SendWsMessage } from "./websocket";

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

// main game loop
let gameFrame = 0;
function animate() {
  let movement = store.state.movement;
  let currentPlayerName = store.state.currentPlayerName;

  // sends all 4 movements
  if (movement.move) {
    SendWsMessage("PLAYER_MOVE", currentPlayerName, movement.move);
  } else if (movement.stop) {
    SendWsMessage("PLAYER_MOVE", currentPlayerName, movement.stop);
    //send movement stop only once, so clear the variable after sending
    store.dispatch("clearStopMovement");
  }

  if (movement.bomb) {
    SendWsMessage("PLAYER_DROPPED_BOMB", currentPlayerName);
    store.dispatch("clearBombDrop");
  }

  gameFrame++;
  requestAnimationFrame(animate);
}

export { store, router, gameFrame };
