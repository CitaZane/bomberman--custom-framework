import createRouter from "../framework/router";
import router from "./router";

createRouter(router);

window.addEventListener("DOMContentLoaded", () => {
    // make sure user can be on the game page only if he has a websocket connection
})