/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ChatRoom } from "../components/ChatRoom";
import { ws } from "../websocket";
import { LobbyPlayers } from "../components/LobbyPlayers";
function startGame() {
  ws.send(
    JSON.stringify({
      type: "START_GAME",
    })
  );
}

export function QueueView() {
  let userQueueCount = store.state.userQueueCount;

  return {
    template: (
      <div id="lobby-layout">
        <button onClick={startGame}>Start Game</button>
        <h2>
          {/* Game starts in<span id="timer">10</span> */}
          Need at least 2 players to start timer
        </h2>
        <LobbyPlayers />
        <ChatRoom />
      </div>
    ),
  };
}
