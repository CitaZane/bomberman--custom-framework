/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ChatRoom } from "../components/ChatRoom";
import { ws } from "../websocket";
import { LobbyPlayers } from "../components/LobbyPlayers";
import { state } from "../store/index";

function startGame() {
	ws.send(
		JSON.stringify({
			type: "START_GAME",
		})
	);
}

export function QueueView() {
	if (!ws) {
		window.location.href = window.location.origin + "/";
		return;
	}

	return {
		template: (
			<div id="lobby-layout">
				<button onClick={startGame}>Start Game</button>
				<h2 id="queueMessage"></h2>
				<LobbyPlayers />
				<ChatRoom />
			</div>
		),
	};
}
