/* @jsx jsx */
import store from "../store";
import jsx from "../framework/vDom/jsx";

export function LobbyPlayers() {
  const players = store.state.players;
  return {
    template: (
      <div id="lobby-players">
        <h3>Waiting players to join...</h3>
        <ul id="lobby-players_list">
          {players.map((playerName, i) => (
            <li>
              <div class="player-monster" id={`monster-${i}`}></div>
              <p class="player-name">{playerName}</p>
            </li>
          ))}
        </ul>

        <button class="btn">Leave lobby</button>
      </div>
    ),
  };
}
